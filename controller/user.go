package controller

import (
	"encoding/json"
	"log"
	"medikakh/application/utils"
	"medikakh/domain/constants"
	"medikakh/domain/datastore"
	"medikakh/domain/models"
	"medikakh/logic"
	"medikakh/service/authentication"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sinabakh/go-zarinpal-checkout"
)

type UserController interface {
	Register(c *gin.Context)
	RegisterCallback(c *gin.Context)
	ReadUser(c *gin.Context)
	GetUserId(username string) (*string, error)
	Login(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type user struct {
	logic logic.UserLogic
}

func NewUserController(logic logic.UserLogic) UserController {
	u := new(user)
	u.logic = logic
	return u
}

func (u *user) Register(c *gin.Context) {

	var userInfo models.UserRegisterationRequest
	c.BindJSON(&userInfo)
	// in this block vlidation of user information will be checked
	err := utils.CheckUsernameValueValidation(userInfo.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = utils.CheckPasswordValueValidation(userInfo.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	roleCorrectness := utils.CheckForRoleStatmentCorrectness(userInfo.Role)
	if !roleCorrectness {
		log.Println("role statement is incorrect")
		c.JSON(http.StatusBadRequest, "role statement is incorrect")
		return
	}
	//  end of validation check

	// PaymentPrice returns cust of requested role
	price := utils.PaymentPrice(userInfo.Role)

	// in this part we check for user existance in db
	err = u.logic.IsUserExists(constants.SystemRoleObject, userInfo.Username)
	if err == nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "there is an internal error or user may alredy exists")
		return
	}

	// creating a payment gate
	zarinPay, err := zarinpal.NewZarinpal("123456123456123456123456123456123456", true)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "error in payment")
		return
	}
	paymentURL, authority, statusCode, err := zarinPay.NewPaymentRequest(
		price,
		"http://localhost:50501/test/register/callback/"+userInfo.Username,
		"Test", "", "")
	if err != nil {
		if statusCode == -3 {
			log.Println("Amount is not accepted in banking system")
		}
		log.Fatal(err)
	}
	log.Println(authority)  // Save authority in DB
	log.Println(paymentURL) // Send user to paymentURL

	// saving user data in redis db to caching it in callback. this data just remaines for 1 minute
	redisDb := datastore.NewRedisDbConnection()
	data, err := json.Marshal(userInfo)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "error on registration")
		return
	}
	err = redisDb.Set(userInfo.Username, data, time.Minute*1).Err()

	// redirecting client to payment page
	c.Redirect(302, paymentURL)

}

func (u *user) ReadUser(c *gin.Context) {
	username := c.Param("username")
	claimes := utils.GetCurrentUserClaimes(c)
	user, err := u.logic.ReadUser(claimes.UserRole, claimes.Id, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err) // needs to be upgrated
		return
	}

	c.JSON(http.StatusOK, user) // needs to be filtred so private info not be porject to client
}

func (u *user) RegisterCallback(c *gin.Context) {
	// reading user data from redis db
	redisdb := datastore.NewRedisDbConnection()
	uI, err := redisdb.Get(c.Param("username")).Result()
	if err != nil {
		log.Println("error in getting info")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var userInfo models.UserRegisterationRequest
	err = json.Unmarshal([]byte(uI), &userInfo)
	if err != nil {
		log.Println("error in getting info")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// getting result of payment from queey string parameters
	authority := c.Query("Authority")
	status := c.Query("Status")
	if authority == "" || status == "" || status != "OK" {
		log.Println("error in payment")
		c.JSON(http.StatusInternalServerError, "error in payment")
		return
	}

	// checking validity of payment by reciving data
	price := utils.PaymentPrice(userInfo.Role)
	zarinPay, err := zarinpal.NewZarinpal("123456123456123456123456123456123456", true)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "error in payment")
		return
	}

	verified, refID, statusCode, err := zarinPay.PaymentVerification(price, authority)
	if err != nil || !verified {
		if statusCode == 101 {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "this payment alredy done")
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "error in payment")
		return
	}

	// if payment went ok, registration will be started ...
	err = u.logic.Register(userInfo.Username, userInfo.Password, userInfo.Role, userInfo.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error()) // error handling needs to be implemented specially later
		return
	}

	c.JSON(http.StatusOK, "User seccessfully registered ref id : "+refID)

}

func (u *user) GetUserId(username string) (*string, error) {

	name, err := u.logic.GetUserId(constants.SystemRoleObject, username)
	if err != nil {
		return nil, err
	}

	return name, nil
}

func (u user) Login(c *gin.Context) {
	authentication.Login(c)
}

func (u *user) UpdateUser(c *gin.Context) {
	var newUser models.UserUpdate
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error on parsing json request"})
		return
	}

	userClaimes := utils.GetCurrentUserClaimes(c)
	if userClaimes == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on extracting user data from gin context"})
		return
	}
	err = u.logic.UpdateUser(userClaimes.UserRole, userClaimes.Userid.String(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

package controller

import (
	"log"
	"medikakh/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
	ReadUser(c *gin.Context)
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

	//	var userInfo models.UserRegisterationRequest
	//	c.BindJSON(&userInfo)
	//	err := u.logic.Register(c, userInfo.Username, userInfo.Password, userInfo.Role, userInfo.Email)

	err := u.logic.Register(c, "fardin", "54321", "silver", "mmaa@gmail.com")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error()) // error handling needs to be implemented specially later
		return
	}

	c.JSON(http.StatusOK, "User seccessfully registered")

}

func (u *user) ReadUser(c *gin.Context) {
	username := c.Param("username")
	user, err := u.logic.ReadUser(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err) // needs to be upgrated
		return
	}

	c.JSON(http.StatusOK, user) // needs to be filtred so private info not be porject to client
}

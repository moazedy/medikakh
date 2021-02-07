package authentication

import (
	"errors"
	"fmt"
	"medikakh/domain/constants"
	"medikakh/domain/datastore"
	"medikakh/domain/models"
	"medikakh/logic"
	"medikakh/repository"
	"net/http"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var key = []byte(constants.JwtSecretKey)

func Login(c *gin.Context) {
	dbsession, err := datastore.NewCouchbaseSession()
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userLogic := logic.NewUserLogic(repository.NewUserRpo(dbsession))

	cridentials := new(models.Cridentials)
	err = c.BindJSON(&cridentials)
	if err != nil {
		fmt.Println("bind error")
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if cridentials.Username == "" || cridentials.Password == "" {
		fmt.Println("empty fields in cridentials")
		c.AbortWithError(http.StatusBadRequest, errors.New(
			"empty fields are not allowed !"))
		return
	}
	err = userLogic.IsUserExists(constants.SystemRoleObject, cridentials.Username)
	if err != nil {
		fmt.Println(err)
		if err == gocb.ErrNoResult {
			c.AbortWithError(http.StatusNotFound, errors.New("user does not exist"))
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	user, err := userLogic.ReadUser(constants.SystemRoleObject, " ", cridentials.Username)
	if err != nil {
		fmt.Println(err)
		if err == gocb.ErrNoResult {
			c.AbortWithError(http.StatusNotFound, errors.New("user does not exist"))
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	passPointer, err := userLogic.GetPassword(constants.SystemRoleObject, " ", user.Id.String())
	if err != nil {
		fmt.Println(err)
		if err == gocb.ErrNoResult {
			c.AbortWithError(http.StatusNotFound, errors.New("user does not exist"))
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	pass := *passPointer
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(cridentials.Password))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	expTime := time.Now().Add(5 * time.Minute)
	claimes := &models.Claimes{
		Userid:    user.Id,
		UserRole:  user.Role,
		CreatedAt: user.CreatedAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimes)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.SetCookie(
		"mediKakh",
		tokenString,
		int(expTime.Unix()),
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "wellcome " + cridentials.Username + " .you are logged in.",
	})

}

func Authenticlation(c *gin.Context) {
	tokenString, err := c.Cookie("mediKakh")
	if err != nil {
		fmt.Println(err)
		if err == http.ErrNoCookie {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	claimes := new(models.Claimes)
	token, err := jwt.ParseWithClaims(
		tokenString,
		claimes,
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		fmt.Println(err)
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !token.Valid {
		fmt.Println("token is invalid")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}

func Refresh(c *gin.Context) {
	tokenString, err := c.Cookie("mediKakh")
	if err != nil {
		fmt.Println(err)
		if err == http.ErrNoCookie {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	claime := new(models.Claimes)
	token, err := jwt.ParseWithClaims(
		tokenString,
		claime,
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)
	if err != nil {
		fmt.Println(err)
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !token.Valid {
		fmt.Println("token is invlid")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if time.Unix(claime.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		fmt.Println("to soon to refresh request")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	expTime := time.Now().Add(5 * time.Minute)
	claime.ExpiresAt = expTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claime)
	newTokenString, err := newToken.SignedString(key)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.SetCookie(
		"mediKakh",
		newTokenString,
		int(expTime.Unix()),
		"/",
		"localhost",
		true,
		true,
	)

}

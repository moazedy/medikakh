package controller

import (
	"medikakh/domain/datastore"
	"medikakh/logic"
	"medikakh/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	dbSession, err := datastore.NewCouchbaseSession()
	if err != nil {
		panic(err)
	}
	userController := NewUserController(
		logic.NewUserLogic(repository.NewUserRpo(dbSession)),
	)

	engine := gin.Default()
	test := engine.Group("test")
	test.GET("/register", userController.Register)
	test.GET("/callback", CallBack)

	engine.Run(port)
}

func CallBack(c *gin.Context) {
	c.JSON(http.StatusOK, "callback done")
}

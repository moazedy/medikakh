package controller

import (
	"medikakh/domain/datastore"
	"medikakh/logic"
	"medikakh/repository"

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
	test.GET("/register/callback/:username", userController.RegisterCallback)

	engine.Run(port)
}

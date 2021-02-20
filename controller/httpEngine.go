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
	// controllers :
	userController := NewUserController(
		logic.NewUserLogic(repository.NewUserRpo(dbSession)),
	)

	articleController := NewArticleController(
		logic.NewArticleLogic(repository.NewArticleRpo(dbSession)),
	)

	videoController := NewVideoController(
		logic.NewVideoLogic(repository.NewVideoRepo(dbSession)),
	)

	ddController := NewDDcontroller(logic.NewDDLogic(repository.NewDDrepo(dbSession)))

	categoryController := NewCategoryController(logic.NewCategoryLogic(repository.NewCategoryRepo(dbSession)))

	engine := gin.Default()
	test := engine.Group("test")
	test.POST("/register", userController.Register)
	test.GET("/register/callback/:username", userController.RegisterCallback)
	test.POST("/login", userController.Login)

	test.GET("/users/read/:username", userController.ReadUser)
	test.PATCH("/users", userController.UpdateUser)
	test.DELETE("/users/user/:user_id", userController.DeleteUser)

	test.POST("/articles", articleController.Save)
	test.GET("/articles/article/:title", articleController.ReadArticle)
	test.GET("/articles/all", articleController.GetArticlesList)
	test.GET("/articles/category/:category", articleController.GetArticlesByCategory)
	test.DELETE("/articles/article/:title", articleController.DeleteArticle)
	test.PATCH("/articles", articleController.UpdateArticle)

	test.POST("/videos", videoController.Save)
	test.GET("/video/video/:title", videoController.Read)
	test.DELETE("/videos/video/:title", videoController.Delete)
	test.PATCH("/videos", videoController.UpdateVideo)
	test.GET("/videos/all", videoController.GetAllVideosList)
	test.GET("/videos/category/:category", videoController.GetVideosByCategory)

	test.POST("/dd/insert", ddController.InsertData)
	test.GET("/dd/read/:title", ddController.ReadData)
	test.GET("/dd/pattern/:pattern", ddController.ReadDataUsingPattern)

	test.POST("/categories", categoryController.AddCategory)

	engine.Run(port)
}

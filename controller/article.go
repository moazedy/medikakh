package controller

import (
	"errors"
	"medikakh/domain/models"
	"medikakh/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	Save(c *gin.Context)
	ReadArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
}

type article struct {
	logic logic.ArticleLogic
}

func NewArticleController(logic logic.ArticleLogic) ArticleController {
	a := new(article)
	a.logic = logic

	return a
}

func (a *article) Save(c *gin.Context) {
	var newArticle models.Article
	err := c.BindJSON(&newArticle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("failed to get article data"))
		return
	}

	err = a.logic.SaveArticle("", newArticle) // need to be fixed
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "article seccessfuly saved to db")
}

func (a *article) ReadArticle(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, "empty title not accepted")
		return
	}

	newArticle, err := a.logic.ReadArticle("", title) // need to be fixed
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newArticle)

}

func (a *article) DeleteArticle(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, "empty title not accepted")
		return
	}

	err := a.logic.DeleteArticle("", title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "article seccessfuly deleted")

}

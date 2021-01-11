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
	UpdateArticle(c *gin.Context)
	GetArticlesList(c *gin.Context)
	// TODO:
	// GetArticlesByCategory(c *gin.Context)
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

func (a *article) UpdateArticle(c *gin.Context) {
	var art models.ArticleUpdate
	err := c.BindJSON(&art)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in parsing json update request"})
		return
	}

	err = a.logic.UpdateArticle("", art)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while updating article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article updated"})
}

func (a *article) GetArticlesList(c *gin.Context) {
	titles, err := a.logic.GetArticlesList("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, titles)
}

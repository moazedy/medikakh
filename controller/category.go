package controller

import (
	"medikakh/application/utils"
	"medikakh/domain/models"
	"medikakh/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	AddCategory(ctx *gin.Context)
	// TODO ReadCategory(ctx *gin.Context)
	// TODO  DeleteCategory(ctx *gin.Context)
	// TODO  UpdateCategory(ctx *gin.Context)
}

type category struct {
	logic logic.CategoryLogic
}

func NewCategoryController(logic logic.CategoryLogic) CategoryController {
	c := new(category)
	c.logic = logic

	return c
}

func (c *category) AddCategory(ctx *gin.Context) {
	role := utils.ExtractRoleFromToken(ctx)
	if role == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error on extracting role form token"})
		return
	}
	var cat models.CatRequest
	err := ctx.BindJSON(&cat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error in binding json object"})
		return
	}

	newCat := models.Category{
		Name:          cat.Name,
		SubCategories: cat.SubCategories,
	}
	err = c.logic.InsertCategory(*role, newCat)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "category created"})
}

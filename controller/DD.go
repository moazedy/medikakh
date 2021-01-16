package controller

import (
	"medikakh/application/utils"
	"medikakh/domain/models"
	"medikakh/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DDcontroller interface {
	InsertData(c *gin.Context)
	ReadData(c *gin.Context)
	ReadDataUsingPattern(c *gin.Context)
}

type dd struct {
	logic logic.DDLogic
}

func NewDDcontroller(logic logic.DDLogic) DDcontroller {
	d := new(dd)
	d.logic = logic

	return d
}

func (d *dd) InsertData(c *gin.Context) {
	var newDD models.DDmodel
	err := c.BindJSON(&newDD)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error on parsing json request"})
		return
	}

	role := utils.ExtractRoleFromToken(c)
	if role != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to extract role from token"})
		return
	}

	err = d.logic.InsertData(*role, newDD)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on inserting data to dd"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DD added to db"})
}

func (d *dd) ReadData(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "title can not be empty"})
		return
	}

	dd, err := d.logic.ReadDD(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dd)
}

func (d *dd) ReadDataUsingPattern(c *gin.Context) {
	pattern := c.Param("pattern")
	if pattern == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "pattern can not be empty"})
		return
	}

	ddList, err := d.logic.ReadDDbyPattern(pattern)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ddList)
}

package controller

import (
	"medikakh/application/utils"
	"medikakh/domain/models"
	"medikakh/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(c *gin.Context)
	Read(c *gin.Context)
	Delete(c *gin.Context)
	UpdateVideo(c *gin.Context)
	GetVideosByCategory(c *gin.Context)
	GetAllVideosList(c *gin.Context)
}

type video struct {
	logic logic.VideoLogic
}

func NewVideoController(logic logic.VideoLogic) VideoController {
	v := new(video)
	v.logic = logic
	return v
}

func (v *video) Save(c *gin.Context) {
	var newVideo models.Video
	err := c.BindJSON(&newVideo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	role := utils.ExtractRoleFromToken(c)
	if role == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = v.logic.Save(*role, newVideo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "video saved"})
}

func (v *video) Read(c *gin.Context) {
	videoTitle := c.Param("title")

	role := utils.ExtractRoleFromToken(c)
	if role == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on extracting role from token"})
		return
	}

	newVideo, err := v.logic.GetVideo(*role, videoTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newVideo)
}

func (v *video) Delete(c *gin.Context) {
	title := c.Param("title")

	role := utils.ExtractRoleFromToken(c)
	if role == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on extracting role from token"})
		return
	}

	err := v.logic.Delete(*role, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "video deleted"})

}

func (v *video) UpdateVideo(c *gin.Context) {
	role := utils.ExtractRoleFromToken(c)
	if role == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on extracting role from token"})
		return
	}
	var newVid models.VideoUpdate
	err := c.BindJSON(&newVid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error on parsing json request"})
		return
	}

	err = v.logic.Update(*role, newVid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "video updated"})
}

func (v *video) GetVideosByCategory(c *gin.Context) {
	role := utils.ExtractRoleFromToken(c)
	if role == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on extracting role from token"})
		return
	}

	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "category value can not be empty"})
		return
	}

	vids, err := v.logic.GetVideosByCategory(*role, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vids)

}

func (v *video) GetAllVideosList(c *gin.Context) {
	role := utils.ExtractRoleFromToken(c)
	if role == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error on extracting role from token"})
		return
	}

	vids, err := v.logic.GetAllVideosList(*role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vids)

}

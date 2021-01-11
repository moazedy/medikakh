package controller

import (
	"medikakh/domain/models"
	"medikakh/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(c *gin.Context)
	Read(c *gin.Context)
	Delete(c *gin.Context)
	// TODO:
	// UpdateVideo(c *gin.Context)
	// GetVideosByCategory(c *gin.Context)
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

	err = v.logic.Save("", newVideo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "video saved"})
}

func (v *video) Read(c *gin.Context) {
	videoTitle := c.Param("video_title")

	newVideo, err := v.logic.GetVideo("", videoTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newVideo)
}

func (v *video) Delete(c *gin.Context) {
	title := c.Param("title")

	err := v.logic.Delete("", title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "video deleted"})

}

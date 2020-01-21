package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	s := service.CreateVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情
func ShowVideo(c *gin.Context) {
	s := service.ShowVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListVideo 视频列表
func ListVideo(c *gin.Context) {
	s := service.ListVideoService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

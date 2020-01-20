package api

import (
	"github.com/gin-gonic/gin"
	"singo/serializer"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "成功",
	})
}

package routes

import (
	"OmyBlog/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinLogger())

	r.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})
	return r
}

package http

import "github.com/gin-gonic/gin"

type Request interface {
	all(c *gin.Context)
}

package port

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(c *gin.Context)
}

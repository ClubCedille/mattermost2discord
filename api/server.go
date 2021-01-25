package api

import (
	"github.com/gin-gonic/gin"
)

// SetupServer -
func SetupServer() *gin.Engine {
	r := gin.Default()

	r.POST("/", ProcessMattermost)

	return r
}

package api

import "github.com/gin-gonic/gin"

// SetupServer -
func SetupServer() *gin.Engine {
	r := gin.Default()

	// Test handle - to be removed
	r.GET("/", HelloWorld)

	return r
}

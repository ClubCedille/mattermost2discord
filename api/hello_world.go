package api

import "github.com/gin-gonic/gin"

// HelloWorld - to be removed
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world;)",
	})
}

package api

import (
	"fmt"
	"github.com/TylerBrock/colorjson"
	"github.com/gin-gonic/gin"
)

// HelloWorld - to be removed
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world;)",
	})
}

// Simple example function logging payload to console
func LogMattermost(c *gin.Context) {
	/*
		Usually you would type your json to have an easier access to the properties
		and to prevent you from doing tons of nil checks in your code.

		You should not use the generic type map[string]interface{}.
		I used it here to quickly setup this example to allow Jo, Sukh and Alex to continue with the webhook.
	*/
	var payload map[string]interface{}
	c.BindJSON(&payload)

	jsonFormatter := colorjson.NewFormatter()
	jsonFormatter.Indent = 4
	prettyPayload, _ := jsonFormatter.Marshal(payload)

	fmt.Printf("Mattermost payload:\n%s", string(prettyPayload))
}

// You can change the name of this function
func ProcessMattermost() {
	// Here you could implement the processing of the Mattermost payload
}

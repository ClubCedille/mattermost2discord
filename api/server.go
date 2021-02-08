package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DiscordApiCheck struct {
	Status *DiscordStatus `json:"status"`
}

type DiscordStatus struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

func SetupServer() *gin.Engine {
	discordBot := CreateDiscordBot()
	r := gin.Default()
	r.GET("/health/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
	r.POST("/v1/discord-message/", discordBot.SendMessage)
	return r
}


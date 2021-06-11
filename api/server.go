package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiscordAPICheck struct {
	Status *DiscordStatus `json:"status"`
}

type DiscordStatus struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

func SetupServer() *gin.Engine {
	discordBot := CreateDiscordBot()
	r := gin.Default()
	r.Use(HandleBotError)
	r.GET("/healthz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	r.POST("/v1/discord-message", discordBot.SendMessage)

	logger := loggerInstance()
	logger.Info("Server setup done.")
	return r
}

func HandleBotError(context *gin.Context) {
	context.Next()
	lastError := context.Errors.Last()
	if lastError != nil {
		logger := loggerInstance()
		logger.Error(lastError)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": lastError,
		})
	}
}

package api

import "github.com/gin-gonic/gin"

type BotActions interface {
	SendMessage(context *gin.Context)
	GetPayloadFrom(context *gin.Context)
	GetContentFrom(payload Payload)
}

type Content struct {
	User    string
	Message string
}

type Payload struct {
	*DiscordPayload
	*MattermostPayload
}

type DiscordPayload struct{}

type MattermostPayload struct {
	Text     string `json:"text"`
	Username string `json:"user_name"`
	UserID   string `json:"user_id"`
}

package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

/*
The interface to be implemented by
the DiscordBot and the MattermostBot
*/
type Bot interface {
	SendMessage(context *gin.Context)
	GetPayload(context *gin.Context) Payload
	GetContent(payload Payload) Content
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
	Text      string    `json:"text"`
	Username  string    `json:"user_name"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	Timestamp time.Time `json:"timestamp"`
}

package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DiscordBot struct {
	Session *discordgo.Session
}

func CreateDiscordBot() *DiscordBot {
	PanicIfDiscordBotMissesInformation()
	session, _ := discordgo.New(fmt.Sprintf("Bot %s", DiscordToken))
	return &DiscordBot{
		Session: session,
	}
}

func (bot *DiscordBot) SendMessage(context *gin.Context) {
	payload := bot.GetPayload(context)
	content := bot.GetContent(payload)
	logger := loggerInstance()

	discordMessage := fmt.Sprintf("%s said: %s", content.User, content.Message)
	_, err := bot.Session.ChannelMessageSend(DiscordChannel, discordMessage)

	logger.Infof("Sending message to discord",
		zap.String("user", content.User),
		zap.String("message", content.Message),
		zap.String("Channel", DiscordChannel),
	)

	if err != nil {
		logger.Error("DiscordMessageError", err)
		context.Error(err)
		return
	}
}

func (*DiscordBot) GetPayload(context *gin.Context) Payload {
	var payload Payload
	err := context.BindJSON(&payload.MattermostPayload)
	if err != nil {
		logger := loggerInstance()
		logger.Error("GetPayloadError", err)
		context.Error(err)
		return Payload{}
	}
	if payload.MattermostPayload.Token != MattermostToken {
		err = errors.New("status unauthorized")
		context.Error(err)
		return Payload{}
	}

	return payload
}

func (*DiscordBot) GetContent(payload Payload) Content {
	message := strings.TrimSpace(strings.Split(payload.MattermostPayload.Text, TriggerWordMattermost)[1])
	user := payload.MattermostPayload.Username
	timestamp := payload.MattermostPayload.Timestamp
	return Content{
		User:      user,
		Message:   message,
		Timestamp: timestamp,
	}
}

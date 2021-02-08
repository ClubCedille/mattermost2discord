package api

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"strings"
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
	discordMessage := fmt.Sprintf("%s said: %s", content.User, content.Message)
	_, err := bot.Session.ChannelMessageSend(DiscordChannel, discordMessage)

	if err != nil {
		fmt.Printf("DiscordMessageError: %s\n", err)
		return
	}
}

func (*DiscordBot) GetPayload(context *gin.Context) Payload {
	var payload Payload
	err := context.BindJSON(&payload.MattermostPayload)
	if err != nil {
		fmt.Printf("GetPayloadError: %s\n", err)
		return Payload{}
	}
	return payload
}

func (*DiscordBot) GetContent(payload Payload) Content {
	message := strings.TrimSpace(strings.Split(payload.MattermostPayload.Text, TriggerWordMattermost)[1])
	user := payload.MattermostPayload.Username
	return Content{
		User:    user,
		Message: message,
	}
}

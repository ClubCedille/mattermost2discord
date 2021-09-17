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
	message := bot.GetMention(content)

	bot.Session.ChannelTyping(DiscordChannel)
	logger := loggerInstance()

	_, err := bot.Session.ChannelMessageSend(DiscordChannel, message)

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
	message = fmt.Sprintf("%s said: %s", user, message)
	return Content{
		User:    user,
		Message: message,
	}
}

func (bot *DiscordBot) GetMention(content Content) string {

	channel, _ := bot.Session.Channel(DiscordChannel)

	idRoles, _ := bot.Session.GuildRoles(channel.GuildID)

	i := strings.Index(content.Message, "@")
	role := strings.TrimSpace(strings.Split(content.Message[i+1:], " ")[0])

	for _, v := range idRoles {

		if strings.TrimSpace(v.Name) == role {
			content.Message = strings.Replace(content.Message, "@"+role, "<@&"+v.ID+">", -1)
		}
	}

	return content.Message

}

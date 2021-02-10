package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDiscordBot(t *testing.T) {
	DiscordToken = "test"
	DiscordChannel = "test"
	TriggerWordMattermost = "test"
	bot := CreateDiscordBot()
	assert.NotNil(t, bot.Session)
}

func TestDiscordBotGetPayload(t *testing.T) {
	bot := DiscordBot{}
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	testPayload := MattermostPayload{
		Text:     "test",
		Username: "test",
		UserID:   "test",
	}
	jsonData, _ := json.Marshal(testPayload)

	reader := bytes.NewReader(jsonData)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/discord-message", reader)
	realPayload := bot.GetPayload(context)
	assert.Equal(t, &testPayload, realPayload.MattermostPayload)
}

func TestDiscordBotGetPayloadError(t *testing.T) {
	bot := DiscordBot{}
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	bot.GetPayload(context)

	assert.EqualErrorf(t, errors.New("invalid request"), "invalid request", "")
}

func TestDiscordBotGetContent(t *testing.T) {
	bot := DiscordBot{}
	TriggerWordMattermost = "2disc"
	content := bot.GetContent(Payload{
		&DiscordPayload{},
		&MattermostPayload{
			Text:     "2disc test",
			Username: "test",
		},
	})

	assert.Equal(t, content.Message, "test")
	assert.Equal(t, content.User, "test")
}

func TestDiscordBotSendMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	DiscordToken = "test"
	DiscordChannel = "test"
	TriggerWordMattermost = "2disc"
	bot := CreateDiscordBot()

	testPayload := MattermostPayload{
		Text:     "2disc test",
		Username: "test",
		UserID:   "test",
	}
	jsonData, _ := json.Marshal(testPayload)

	reader := bytes.NewReader(jsonData)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/discord-message", reader)

	assert.NotPanics(t, func() {
		bot.SendMessage(context)
	})
}

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SetupTestSuite interface {
	SetupTest()
}

type Suite struct {
	suite.Suite
	DiscordToken          string
	DiscordChannel        string
	TriggerWordMattermost string
	testPayload           MattermostPayload
	bot                   DiscordBot
}

func (suite *Suite) SetupTest() {
	suite.DiscordToken = "test"
	suite.DiscordChannel = "test"
	suite.TriggerWordMattermost = "test"
	suite.testPayload = MattermostPayload{
		Text:     "2disc test",
		Username: "test",
		UserID:   "test",
	}

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestCreateDiscordBot() {
	bot := CreateDiscordBot()
	assert.NotNil(suite.T(), bot.Session)
}

func (suite *Suite) TestDiscordBotGetPayload() {

	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	jsonData, _ := json.Marshal(suite.testPayload)

	reader := bytes.NewReader(jsonData)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/discord-message", reader)

	realPayload := suite.bot.GetPayload(context)
	assert.Equal(suite.T(), &suite.testPayload, realPayload.MattermostPayload)
}

func (suite *Suite) TestDiscordBotGetPayloadError() {

	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	suite.bot.GetPayload(context)

	assert.EqualErrorf(suite.T(), errors.New("invalid request"), "invalid request", "")
}

func (suite *Suite) TestDiscordBotGetContent() {
	content := suite.bot.GetContent(Payload{
		&DiscordPayload{},
		&suite.testPayload,
	})

	assert.Equal(suite.T(), content.Message, "test")
	assert.Equal(suite.T(), content.User, suite.testPayload.Username)
}

func (suite *Suite) TestDiscordBotSendMessage() {
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	bot := CreateDiscordBot()

	jsonData, _ := json.Marshal(suite.testPayload)

	reader := bytes.NewReader(jsonData)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/discord-message", reader)

	assert.NotPanics(suite.T(), func() {
		bot.SendMessage(context)
	})
}

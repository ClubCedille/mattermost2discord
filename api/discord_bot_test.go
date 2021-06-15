package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DiscordTestSuite struct {
	suite.Suite
	DiscordToken          string
	DiscordChannel        string
	TriggerWordMattermost string
	bot                   *DiscordBot
	testPayload           MattermostPayload
}

func (suite *DiscordTestSuite) SetupTest() {
	suite.DiscordToken = "test"
	suite.DiscordChannel = "test"
	suite.TriggerWordMattermost = "2disc"
	suite.bot = &DiscordBot{}
	suite.testPayload = MattermostPayload{
		Text:      "2disc test",
		Username:  "test",
		UserID:    "test",
		Token:     "test",
		Timestamp: time.Now(),
	}

	DiscordToken = suite.DiscordToken
	DiscordChannel = suite.DiscordChannel
	TriggerWordMattermost = suite.TriggerWordMattermost
	MattermostToken = suite.testPayload.Token
}

func TestDiscordBot(t *testing.T) {
	suite.Run(t, new(DiscordTestSuite))
}

func (suite *DiscordTestSuite) TestCreateDiscordBot() {
	bot := CreateDiscordBot()
	assert.NotNil(suite.T(), bot.Session)
}

func (suite *DiscordTestSuite) TestDiscordBotGetPayload() {
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	jsonData, _ := json.Marshal(suite.testPayload)

	reader := bytes.NewReader(jsonData)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/discord-message", reader)
	realPayload := suite.bot.GetPayload(context)
	assert.Equal(suite.T(), &suite.testPayload, realPayload.MattermostPayload)
}

func (suite *DiscordTestSuite) TestDiscordGetPayloadError() {

	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	suite.bot.GetPayload(context)
	assert.EqualErrorf(suite.T(), context.Errors.Last(), "invalid request", "")

	// Providing false token
	suite.testPayload.Token = "foo"

	jsonData, _ := json.Marshal(suite.testPayload)

	reader := bytes.NewReader(jsonData)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/discord-message", reader)
	suite.bot.GetPayload(context)

	assert.EqualErrorf(suite.T(), context.Errors.Last(), "status unauthorized", "")
}

func (suite *DiscordTestSuite) TestDiscordBotGetContent() {
	content := suite.bot.GetContent(Payload{
		&DiscordPayload{},
		&suite.testPayload,
	})

	const desiredResult = "test"
	assert.Equal(suite.T(), desiredResult, content.Message)
	assert.Equal(suite.T(), suite.testPayload.Username, content.User)
}

func (suite *DiscordTestSuite) TestDiscordBotSendMessage() {
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

func (suite *DiscordTestSuite) TestDiscordBotTimestamp() {
	content := suite.bot.GetContent(Payload{
		&DiscordPayload{},
		&suite.testPayload,
	})

	assert.NotEmpty(suite.T(), content.Timestamp)
}

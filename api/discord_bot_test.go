package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/ewohltman/discordgo-mock/mockchannel"
	"github.com/ewohltman/discordgo-mock/mockconstants"
	"github.com/ewohltman/discordgo-mock/mockguild"
	"github.com/ewohltman/discordgo-mock/mockrest"
	"github.com/ewohltman/discordgo-mock/mockrole"
	"github.com/ewohltman/discordgo-mock/mocksession"
	"github.com/ewohltman/discordgo-mock/mockstate"
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
	suite.TriggerWordMattermost = "2disc"
	suite.DiscordChannel = mockconstants.TestChannel
	suite.bot = &DiscordBot{}
	suite.testPayload = MattermostPayload{
		Text:     "2disc test",
		Username: "test",
		UserID:   "test",
		Token:    "test",
	}
	DiscordToken = suite.DiscordToken
	DiscordChannel = suite.DiscordChannel
	TriggerWordMattermost = suite.TriggerWordMattermost
	MattermostToken = suite.testPayload.Token
}

func mockCreateDiscordBot() *DiscordBot {
	PanicIfDiscordBotMissesInformation()
	state, _ := newState()
	session, _ := mocksession.New(
		mocksession.WithState(state),
		mocksession.WithClient(&http.Client{
			Transport: mockrest.NewTransport(state),
		}),
	)

	guildBefore, _ := session.Guild(mockconstants.TestGuild)
	session.GuildRoleCreate(guildBefore.ID)

	return &DiscordBot{
		Session: session,
	}
}

func newState() (*discordgo.State, error) {

	role := mockrole.New(
		mockrole.WithID(mockconstants.TestRole),
		mockrole.WithName(mockconstants.TestRole),
		mockrole.WithPermissions(discordgo.PermissionViewChannel),
	)
	role.Mentionable = true
	channel := mockchannel.New(
		mockchannel.WithID(mockconstants.TestChannel),
		mockchannel.WithGuildID(mockconstants.TestGuild),
		mockchannel.WithName(mockconstants.TestChannel),
		mockchannel.WithType(discordgo.ChannelTypeGuildVoice),
	)

	return mockstate.New(
		mockstate.WithGuilds(
			mockguild.New(
				mockguild.WithID(mockconstants.TestGuild),
				mockguild.WithName(mockconstants.TestGuild),
				mockguild.WithRoles(role),
				mockguild.WithChannels(channel),
			),
		),
	)
}

func TestDiscordBot(t *testing.T) {
	suite.Run(t, new(DiscordTestSuite))
}

func (suite *DiscordTestSuite) TestCreateDiscordBot() {
	bot := mockCreateDiscordBot()
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

	gin.SetMode(gin.TestMode)
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
	suite.testPayload.Text = "2disc test @" + mockconstants.TestRole
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	jsonData, _ := json.Marshal(suite.testPayload)
	reader := bytes.NewReader(jsonData)
	context.Request = httptest.NewRequest(http.MethodPost, "/v1/discord-message", reader)
	bot := mockCreateDiscordBot()
	assert.NotPanics(suite.T(), func() {
		bot.SendMessage(context)
	})
}

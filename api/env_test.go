package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicIfDiscordBotMissesInformation(T *testing.T) {
	DiscordToken = ""
	DiscordChannel = "test"
	TriggerWordMattermost = "test"
	MattermostToken = "test"
	assert.Panicsf(T, func() {
		CreateDiscordBot()
	}, DiscordTokenMissingMessage)
	DiscordToken = "test"
	DiscordChannel = ""
	assert.Panicsf(T, func() {
		CreateDiscordBot()
	}, DiscordChannelMissingMessage)
	DiscordChannel = "test"
	TriggerWordMattermost = ""
	assert.Panicsf(T, func() {
		CreateDiscordBot()
	}, MattermostTriggerwordMissingMessage)
	TriggerWordMattermost = "test"
	MattermostToken = ""
	assert.Panicsf(T, func() {
		CreateDiscordBot()
	}, mattermostTokenMissingMessage)
}

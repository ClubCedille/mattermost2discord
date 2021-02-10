package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicIfDiscordBotMissesInformation(T *testing.T) {
	DiscordToken = ""
	DiscordChannel = "test"
	TriggerWordMattermost = "test"
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
}

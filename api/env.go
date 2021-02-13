package api

import (
	"os"
)

var (
	DiscordToken                        = os.Getenv("DISCORD_TOKEN")
	DiscordChannel                      = os.Getenv("DISCORD_CHANNEL")
	TriggerWordMattermost               = os.Getenv("TRIGGER_WORD_MATTERMOST")
	MattermostToken                     = os.Getenv("MATTERMOST_TOKEN")
	DiscordTokenMissingMessage          = "No discord token provided!"
	DiscordChannelMissingMessage        = "No discord channel provided!"
	MattermostTriggerwordMissingMessage = "No trigger word provided!"
	mattermostTokenMissingMessage       = "No mattermost token provided!"
)

func PanicIfDiscordBotMissesInformation() {
	if len(DiscordToken) == 0 {
		panic(DiscordTokenMissingMessage)
	} else if len(DiscordChannel) == 0 {
		panic(DiscordChannelMissingMessage)
	} else if len(TriggerWordMattermost) == 0 {
		panic(MattermostTriggerwordMissingMessage)
	} else if len(MattermostToken) == 0 {
		panic(mattermostTokenMissingMessage)
	}
}

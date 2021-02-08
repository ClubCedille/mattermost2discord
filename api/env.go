package api

import (
	"os"
)

var (
	DiscordToken            = os.Getenv("DISCORD_TOKEN")
	DiscordChannel          = os.Getenv("DISCORD_CHANNEL")
	DiscordApiCheckEndpoint = os.Getenv("DISCORD_API_CHECK_ENDPOINT")
	TriggerWordMattermost   = os.Getenv("TRIGGER_WORD_MATTERMOST")
)

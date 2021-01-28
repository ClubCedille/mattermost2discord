package api

import (
	"os"
)

var (
	DiscordToken          = os.Getenv("DISCORD_TOKEN")
	DiscordChannel        = os.Getenv("DISCORD_CHANNEL")
	TriggerWordMattermost = os.Getenv("TRIGGER_WORD_MATTERMOST")
)

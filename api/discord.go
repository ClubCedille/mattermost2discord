package api

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// DiscordBot activates the bot in the discord server
func discordBot(mmUser string, mmContent string) {

	goBot, err := discordgo.New("Bot " + DiscordToken)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msgtoDisc := fmt.Sprintf("%s: %s", mmUser, mmContent)
	_, err = goBot.ChannelMessageSend(DiscordChannel, msgtoDisc)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

}

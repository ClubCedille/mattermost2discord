package api

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// DiscordBot activates the bot in the discord server
func discordBot(mmContent string) {

	goBot, err := discordgo.New("Bot " + os.Getenv("dtoken"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = goBot.ChannelMessageSend(os.Getenv("dchannel"), mmContent)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

}

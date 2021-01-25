package api

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

// DiscordBot activates the bot in the discord server
func discordBot(mmUser string, mmContent string) {

	goBot, err := discordgo.New("Bot " + os.Getenv("dtoken"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msgtoDisc := fmt.Sprintf("%s: %s", mmUser, mmContent)
	_, err = goBot.ChannelMessageSend(os.Getenv("dchannel"), msgtoDisc)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

}

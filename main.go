package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var token string = "MTA3MTgwOTYyNDcxOTc3Mzc1Ng.G4nkQ0.TS_LRZRuya9BJZX-WC5ZIOc0_IYgsom1PcW8-c"

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Intents
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Event handlers
	discord.AddHandler(onMessage)

	err = discord.Open()
	defer discord.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Session opened.")
		<-make(chan struct{})
	}
}

func onMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	} else {
		if strings.HasPrefix(message.Content, "&go") {
			var messages []*discordgo.Message
			var err error
			messages, err = discord.ChannelMessages(message.ChannelID, 100, "", "", "")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			discord.ChannelMessageSend(message.ChannelID, messages[99].Content)
		}
	}
}

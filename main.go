package main

import (
	"fmt"

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
	discord.AddHandler(messageHandler)

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

func messageHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Username == discord.State.User.ID {
		return
	} else {
		discord.ChannelMessageSend(message.ChannelID, message.Content)
	}
}

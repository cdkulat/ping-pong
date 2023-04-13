package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "MTA5NjE3MTYyMzkwMTU3NzMwNg.Gq9krr.JYtby-gBnhPSCL947oB6BMusB7SE8SbMmisu6k")
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return // Ignore messages sent by bots
	}

	if strings.HasPrefix(m.Content, "!ping") {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

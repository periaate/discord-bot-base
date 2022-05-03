package cmd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var cmd = discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "pong",
}

// ping pong slash command
func pingpongSlash(bot *discordgo.Session, i *discordgo.InteractionCreate) {
	resp := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}
	err := bot.InteractionRespond(i.Interaction, &resp)
	if err != nil {
		log.Println("error reacting to interaction", err)
	}
}

// ping pong message command
func pingpongMsg(bot *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := bot.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		log.Println("error sending message", err)
	}
}

func init() {
	_, cmdFound := Commands["ping"]
	_, handlerFound := Handlers["ping"]
	_, messageCmd := MessageCreate["ping"]
	if !cmdFound && !handlerFound && !messageCmd {
		Commands["ping"] = &cmd
		Handlers["ping"] = pingpongSlash
		MessageCreate["ping"] = pingpongMsg
	}
}

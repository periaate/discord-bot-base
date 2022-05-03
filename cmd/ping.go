package cmd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const name = "ping"

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
	_, cmdFound := Commands[name]
	_, handlerFound := Handlers[name]
	_, messageCmd := MessageCreate[name]
	if !cmdFound && !handlerFound && !messageCmd {
		Commands[name] = &cmd
		Handlers[name] = pingpongSlash
		MessageCreate[name] = pingpongMsg
	}
}

package cmd

import "github.com/bwmarrin/discordgo"

// These maps make it so that adding functions to the bot is done on a file by file basis
var (
	// Commands is an array of metadata for commands
	Commands = map[string]*discordgo.ApplicationCommand{}
	// Handlers is an array of functions for commands
	Handlers = map[string]func(bot *discordgo.Session, i *discordgo.InteractionCreate){}
	// MessageCreate is an array of functions for MessageCreate commands
	MessageCreate = map[string]func(bot *discordgo.Session, m *discordgo.MessageCreate){}
)

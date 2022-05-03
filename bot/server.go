package bot

import (
	"botbase/cmd"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	// Bot is public
	Bot *discordgo.Session
	c   = new(cfg)
)

func init() {
	var err error
	c.load()
	Bot, err = discordgo.New("Bot " + c.Token)
	if err != nil {
		log.Fatalln(err)
	}
}

// Start discord bot server
func Start() {
	Bot.AddHandler(func(bot *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", bot.State.User.Username, bot.State.User.Discriminator)
	})
	err := Bot.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding cmd.Commands...")
	for k, v := range cmd.Commands {
		_, err := Bot.ApplicationCommandCreate(Bot.State.User.ID, c.Guild, v)
		if err != nil {
			log.Panicf("Cannot create '%v' Command: %v", v.Name, err)
			// Remove command from map
			delete(cmd.Commands, k)
		}
	}

	defer Bot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}

func init() {
	Bot.AddHandler(func(Bot *discordgo.Session, i *discordgo.InteractionCreate) {
		if fn, ok := cmd.Handlers[i.ApplicationCommandData().Name]; ok {
			fn(Bot, i)
		}
	})

	Bot.AddHandler(func(Bot *discordgo.Session, m *discordgo.MessageCreate) {
		if fn, ok := cmd.MessageCreate[m.Content]; ok {
			fn(Bot, m)
		}
	})
}

// Create a config.json file in the base of the project (where main.go is)
// with the following fields (guild may not be necessary)
type cfg struct {
	Token string `json:"token"`
	Guild string `json:"Guild"`
}

// Load config.json file and parse it to struct
func (c *cfg) load() {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, c)
	if err != nil {
		log.Fatalln(err)
	}
}

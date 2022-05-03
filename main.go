package main

import (
	"botbase/bot"
	_ "botbase/cmd" // Using the underscore is so that the init is run
)

func main() {
	bot.Start()
}

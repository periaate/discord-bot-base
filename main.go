package main

import (
	_ "botbase/cmd" // Using the underscore is so that the init is run
	"botbase/server"
)

func main() {
	server.Start()
}

# discord-bot-base
A base for a discord bot  

## Installation
You will need to create a config.json file, config.json file has two fields, `"token"` and `"guild"`.  
Token is your bot token. Guild is your server id, and is not necessary to be set at this moment.  

## Running

to run server: 
`go run main.go`

## Structure

Slash/message commands are all located under cmd.  
  
To create a new command, create a new file, create the functions/structs necessary.  
Add them to cmd/cmd.go maps in an init function.

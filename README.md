# discord-bot-base
A base for a discord bot


create a file called "config.json"  
In this file, add two string fields "token" and "guild".  
These should your bot token, and your testing guild (guild not currently used)  
to run server: 
`go run main.go`


```
main.go -> run server / bot

server.go -> run bot, load handlers for slash, msg commands

cmd/cmd -> create package (folder) wide variables for defining functions

cmd/* -> files define their functionality, and add them accordingly to cmd variables in the init() function

go run main.go
init functions are run in all files that were imported from cmd with an _ in main.go

Server init adds functions to Bot
server.Start is ran, begins listening for events.
If event happens, the command used is matched against the maps in cmd, if found, run the function found in the map
```

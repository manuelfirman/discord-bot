package main

import (
	"fmt"

	"github.com/manuelfirman/discord-bot/config"
	"github.com/manuelfirman/discord-bot/internal/bot"
)

func main() {
	// Read the config file
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Start the bot
	bot.Run()

	// Wait here
	<-make(chan struct{})
	return
}

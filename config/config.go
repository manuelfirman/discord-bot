package config

import (
	"fmt"
	"os"
)

// Global variables for the bot
var (

	// Token: the bot token
	Token string
	// BotPrefix: the bot prefix for commands
	BotPrefix string

	// ErrConfigNotSet is returned when the config file is not set
	ErrConfigNotSet = fmt.Errorf("cannot read from env variables. please check if they are set")
)

// ReadConfig reads the config file and sets the global variables
func ReadConfig() (err error) {
	// Set the env variables
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	// If the env variables are not set, read from the config file
	if Token == "" || BotPrefix == "" {
		err = ErrConfigNotSet
		fmt.Println(err)
		return
	}

	return
}

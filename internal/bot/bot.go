package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/manuelfirman/discord-bot/config"
)

var (
	// BotID: the bot ID
	BotID string
	// goBot: the Discord session
	goBot *discordgo.Session
)

func Run() {
	// Create a new Discord session using the provided bot token.
	goBot, err := discordgo.New("Bot" + config.Token)
	if err != nil {
		println(err.Error())
		return
	}

	// Get the account information
	user, err := goBot.User("@me")
	if err != nil {
		println(err.Error())
		return
	}

	// Store the account ID for later use
	BotID = user.ID

	// Register the messageCreate func as a callback for MessageCreate events.
	goBot.AddHandler(messageHandler)

	// Open a websocket connection to Discord and begin listening.
	err = goBot.Open()
	if err != nil {
		println(err.Error())
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is running! Press CTRL-C to exit.")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	var message string

	// switch on the message content
	switch m.Content {
	case config.BotPrefix:
		message = "I'm here!"
	case "ping":
		message = "Pong!"
	case "pong":
		message = "Ping!"
	case "hello":
		message = "Hi!"
	case "bye":
		message = "Goodbye!"
	case "music":
		message = "https://www.youtube.com/watch?v=E17hDWJKBsY&t=368s"
	case "help":
		message = "I'm here to help you! You can use the following commands: \n" + config.BotPrefix + "ping \n" + config.BotPrefix + "pong \n" + config.BotPrefix + "hello \n" + config.BotPrefix + "bye \n" + config.BotPrefix + "music"
	default:
		message = "I don't understand that command. Try " + config.BotPrefix + "help" + " for a list of commands"
	}

	// Send the message to the channel
	s.ChannelMessageSend(m.ChannelID, message)
}

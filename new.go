package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// The struct for the bot.
type Bot struct {
	Session *discordgo.Session
	BotName string
	// Reply rules.
	replyRules []ReplyRule
}

// Create a new bot.
func New(token string, botName string) *Bot {
	// Return value error is always nil.
	newSession, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	// The new bot is created with no rules.
	newBot := &Bot{
		Session:    newSession,
		BotName:    botName,
		replyRules: []ReplyRule{},
	}
	return newBot
}

// Start the bot.
func (bot *Bot) Open() error {
	err := bot.Session.Open()
	if err != nil {
		return err
	}
	fmt.Println("Bot", bot.BotName, "opened.")
	return nil
}

// Close the bot.
func (bot *Bot) Close() {
	bot.Session.Close()
	fmt.Println("Bot", bot.BotName, "closed.")
}

package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	Session *discordgo.Session
	BotName string
	// Reply rules.
	replyRules []ReplyRule
}

// Create a new bot.
func New(token string, botName string) *DiscordBot {
	// Return value error is always nil.
	newSession, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	newBot := &DiscordBot{
		Session:    newSession,
		BotName:    botName,
		replyRules: []ReplyRule{},
	}
	return newBot
}

// Start a bot.
func (bot *DiscordBot) Open() error {
	err := bot.Session.Open()
	if err != nil {
		return err
	}
	fmt.Println("Bot", bot.BotName, "opened.")
	return nil
}

// Close a bot.
func (bot *DiscordBot) Close() {
	bot.Session.Close()
	fmt.Println("Bot", bot.BotName, "closed.")
}

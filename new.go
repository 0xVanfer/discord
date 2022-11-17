package discord

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Create a new bot.
//
// If the bot is not a reply bot, "replyFrequency" can use 0.
func New(token string, botName string, replyFrequency time.Duration) *DiscordBot {
	// Return value error is always nil.
	newSession, _ := discordgo.New("Bot " + token)
	newBot := &DiscordBot{
		Session:        newSession,
		BotName:        botName,
		lastRead:       map[string]msgInfo{},
		replyRules:     []ReplyRule{},
		replyFrequency: replyFrequency,
	}
	return newBot
}

// Start a bot.
func (bot *DiscordBot) Open() error {
	err := bot.Session.Open()
	if err != nil {
		return err
	}
	fmt.Println("bot", bot.BotName, "opened")
	return nil
}

// Close a bot.
func (bot *DiscordBot) Close() {
	bot.Session.Close()
	fmt.Println("bot closed")
}

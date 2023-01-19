package discord

import (
	"strings"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Simple rule for reply.
type ReplyRule struct {
	// Channel id in string.
	ChannelIDs []string
	// 0: Equalfold. The msg content must be equalfold as the required text.
	//
	// 1: Contain. The msg content must contain the required text.
	//
	// 2: Start with. The msg content must start with the required text.
	//
	// 3. Any.
	RuleType int
	// Only used when "RuleType" is 2.
	//
	// Msg content after the required text may has a length limit.
	// 0 for no limit.( If limited to 0, use 0 for "RuleType" instead.)
	LengthLimit int
	// Text to check if should reply.
	CheckText string

	// Use function to decide what to reply.
	ReplyFunc func(bot *DiscordBot, msg *discordgo.MessageCreate) (replyMsg *discordgo.MessageSend)
	// Use function to decide what to react.
	ReactFunc func(bot *DiscordBot, msg *discordgo.MessageCreate) (reactEmojiIDs []string)

	// Whether should reply to the user replies to.
	ReplyToInitialMessage bool
	// Whether to reply in DM.
	// Meaningless if ReplyToInitialMessage is true.
	ReplyInDM bool

	// Whether the author can be a bot.
	ReplyToBot bool
	// Whether reply to the msg sent by the bot itself.
	// Meaningless if ReplyToBot is false.
	ReplyToSelf bool
}

// Add reply rules to the bot.
func (bot *DiscordBot) AddReplyRules(rules ...ReplyRule) {
	bot.replyRules = append(bot.replyRules, rules...)
}

// Whether the rule should be replied.
func (rule *ReplyRule) shouldReply(msg *discordgo.MessageCreate) bool {
	// Should not reply if channel not match.
	if !utils.ContainInArrayX(msg.ChannelID, rule.ChannelIDs) {
		return false
	}
	if !rule.ReplyToBot && msg.Author.Bot {
		return false
	}
	switch rule.RuleType {
	// Equalfold.
	case 0:
		return strings.EqualFold(rule.CheckText, msg.Content)
	// Contain.
	case 1:
		if rule.CheckText == "" {
			return false
		}
		return strings.Contains(strings.ToLower(msg.Content), strings.ToLower(rule.CheckText))
	// Start with.
	case 2:
		if rule.CheckText == "" {
			return false
		}
		if len(msg.Content) <= len(rule.CheckText) {
			return false
		}
		if !strings.EqualFold(msg.Content[:len(rule.CheckText)], rule.CheckText) {
			return false
		}
		if rule.LengthLimit == 0 {
			return true
		}
		return len(msg.Content[len(rule.CheckText):]) == rule.LengthLimit
	// Any.
	case 3:
		return true
	default:
		return false
	}
}

// Whether the rule should be replied.
func (rule *ReplyRule) shouldReact(msg *discordgo.MessageCreate) bool {
	// Should not reply if channel not match.
	if !utils.ContainInArrayX(msg.ChannelID, rule.ChannelIDs) {
		return false
	}
	switch rule.RuleType {
	// Contain.
	case 1:
		if rule.CheckText == "" {
			return false
		}
		return strings.Contains(strings.ToLower(msg.Content), strings.ToLower(rule.CheckText))
		// Any.
	case 3:
		return true
	default:
		return false
	}

}

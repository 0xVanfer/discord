package discord

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Add a reply rule to the bot.
func (bot *DiscordBot) AddReplyRule(rule ReplyRule) error {
	// Try to read the channel's last msg.
	msgs, err := bot.Session.ChannelMessages(rule.ChannelID, 1, "", "", "")
	if err != nil {
		// Has no access to this channel.
		return err
	}

	// If content is "", auth of the bot is not enough.
	content := msgs[0].Content
	if content == "" {
		return errors.New("bot auth not enough, could not read msg content")
	}

	// Add a new rule.
	bot.ReplyRules = append(bot.ReplyRules, rule)

	// If the channel already exist, skip and not update last read map.
	for channel := range bot.LastRead {
		if channel == rule.ChannelID {
			return nil
		}
	}

	// Last read map add the channel.
	bot.LastRead[rule.ChannelID] = MsgInfo{}
	// Update last read.
	err = bot.UpdateLastReadMsgs()
	if err != nil {
		return err
	}
	return nil
}

// Whether the rule should be replied.
func (rule *ReplyRule) shouldReply(msg *discordgo.Message) bool {
	if msg.ChannelID != rule.ChannelID {
		return false
	}
	switch rule.RuleType {
	// Equalfold.
	case 0:
		return strings.EqualFold(rule.CheckText, msg.Content)
	// Contain.
	case 1:
		return strings.Contains(strings.ToLower(msg.Content), strings.ToLower(rule.CheckText))
	// Start with.
	case 2:
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
	default:
		return false
	}
}

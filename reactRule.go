package discord

import (
	"strings"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Add reaction to a msg.
func (bot *DiscordBot) react(channelID string, msgID string, function any) {
	var emojiIDs []string
	switch v := function.(type) {
	case []string:
		emojiIDs = v
	case func(*DiscordBot, string, string) []string:
		emojiIDs = v(bot, channelID, msgID)
	}
	if emojiIDs == nil {
		return
	}
	for _, emojiID := range emojiIDs {
		bot.Session.MessageReactionAdd(channelID, msgID, emojiID)
	}
}

// Add reply rules to the bot.
func (bot *DiscordBot) AddReactRules(rules ...ReactRule) {
	for _, rule := range rules {
		for _, channelID := range rule.ChannelIDs {
			// If the channel already exist, skip and not update last read map.
			exist := false
			for channel := range bot.lastRead {
				if channel == channelID {
					exist = true
					break
				}
			}
			if !exist {
				// Last read map add this new channel.
				bot.lastRead[channelID] = msgInfo{}
			}
		}
		// Add a new rule.
		bot.reactRules = append(bot.reactRules, rule)
	}
}

// Whether the rule should be replied.
func (rule *ReactRule) shouldReact(msg *discordgo.Message) bool {
	// Should not reply if channel not match.
	if !utils.ContainInArrayX(msg.ChannelID, rule.ChannelIDs) {
		return false
	}
	switch rule.RuleType {
	// Any.
	case 0:
		return true
	// Contain.
	case 1:
		if rule.RequiredText == "" {
			return false
		}
		return strings.Contains(strings.ToLower(msg.Content), strings.ToLower(rule.RequiredText))
	default:
		return false
	}

}

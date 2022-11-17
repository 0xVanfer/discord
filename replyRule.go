package discord

import (
	"strings"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Deprecated: Use AddReplyRules instead.
//
// Add a reply rule to the bot.
func (bot *DiscordBot) AddReplyRule(rule ReplyRule) {
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
	bot.replyRules = append(bot.replyRules, rule)
}

// Add reply rules to the bot.
func (bot *DiscordBot) AddReplyRules(rules ...ReplyRule) {
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
		bot.replyRules = append(bot.replyRules, rule)
	}
}

// Whether the rule should be replied.
func (rule *ReplyRule) shouldReply(msg *discordgo.Message) bool {
	// Should not reply if channel not match.
	if !utils.ContainInArrayX(msg.ChannelID, rule.ChannelIDs) {
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

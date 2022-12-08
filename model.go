package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	Session        *discordgo.Session
	BotName        string
	lastRead       map[string]msgInfo
	replyRules     []ReplyRule
	replyFrequency time.Duration
	reactRules     []ReactRule
}

// Simple msg info for recording last read msg.
type msgInfo struct {
	MsgID  string    // Msg id.
	SendAt time.Time // Last update time.
}

// Simple rule for reply.
type ReplyRule struct {
	// Channel id in string.
	ChannelIDs []string
	// 0: Equalfold. The msg content must be equalfold as the required text.
	// 1: Contain. The msg content must contain the required text.
	// 2: Start with. The msg content must start with the required text.
	RuleType int
	// Only used when "RuleType" is 2. Msg content after the required text may has a length limit.
	// 0 for no limit.( If limited to 0, use 0 for "RuleType" instead.)
	LengthLimit int
	// Text to check if should reply.
	CheckText string
	// Text to reply.
	ReplyText string
	// Use function to decide what to reply.
	// Input: bot, channel, msgId.
	// Output: replyText.
	ReplyFunc func(bot *DiscordBot, channelID, msgID string) (replyText string)
	// Whether to reply in DM.
	ReplyInDM bool
}

type ReactRule struct {
	// Channel id in string.
	ChannelIDs []string
	// 0: Any. React to any msg.
	// 1: Contain. The msg content must contain the required text.
	RuleType int
	// Only used when "RuleType" is 1.
	RequiredText string
	// Emoji to react.
	//
	// Emoji string: github.com/enescakir/emoji suggested.
	ReactEmojiIDs []string
	// Use function to decide what to react.
	// Input: bot, channel, msgId.
	// Output: reply emojis.
	ReactFunc func(bot *DiscordBot, channelID, msgID string) (reactEmojiIDs []string)
}

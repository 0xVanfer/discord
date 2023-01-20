package discord

import (
	"fmt"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Reply to a msg.
func (bot *DiscordBot) reply(msg *discordgo.Message, rule ReplyRule) {
	replyMsg := rule.ReplyFunc(bot, msg)
	// Some error occured or nothing to reply.
	if replyMsg == nil {
		return
	}

	// Whether should reply in DM.
	if rule.ReplyInDM {
		// DM message has no reference msg.
		replyMsg.Reference = nil
		bot.SendDM(msg.Author.ID, replyMsg)
		return
	}

	// Change the target to reply to.
	if rule.ReplyToInitialMessage && msg.ReferencedMessage != nil {
		replyMsg.Reference = msg.ReferencedMessage.Reference()
	}

	// If not replying to anything, should send the msg in ReplyFunc.
	if replyMsg.Reference == nil {
		replyMsg.Reference = msg.Reference()
	}

	// Send the reply.
	bot.Session.ChannelMessageSendComplex(msg.ChannelID, replyMsg)
}

// React to a msg.
func (bot *DiscordBot) react(msg *discordgo.Message, rule ReplyRule) {
	reactEmojiIDs := rule.ReactFunc(bot, msg)
	// Some error occured or nothing to react.
	if reactEmojiIDs == nil {
		return
	}

	// React.
	for _, emojiID := range reactEmojiIDs {
		bot.Session.MessageReactionAdd(msg.ChannelID, msg.ID, emojiID)
	}
}

// Start the reply and never stop.
func (bot *DiscordBot) StartReply() {
	defer utils.Restart(bot.StartReply)
	fmt.Println("Bot", bot.BotName, "start to listen and reply.")
	bot.Session.AddHandler(
		func(se *discordgo.Session, msg *discordgo.MessageCreate) {
			for _, rule := range bot.replyRules {
				if rule.shouldReply(msg.Message) {
					bot.reply(msg.Message, rule)
				}
				if rule.shouldReact(msg.Message) {
					bot.react(msg.Message, rule)
				}
			}
		},
	)
	select {}
}

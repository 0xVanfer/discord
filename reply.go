package discord

import (
	"fmt"
	"time"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Reply to a msg.
func (bot *DiscordBot) reply(channelID string, msg *discordgo.Message, rule ReplyRule) {
	var replyMsg *discordgo.MessageSend

	if rule.ReplyMsg != nil {
		// Use reply msg.
		replyMsg = rule.ReplyMsg
	} else if rule.ReplyMsgFunc != nil {
		// Use reply msg func.
		replyMsg = rule.ReplyMsgFunc(bot, msg)
	} else if rule.ReplyFunc != nil {
		// Deprecated.
		fmt.Println("ReplyRule.ReplyFunc is deprecated. Use ReplyMsgFunc instead.")
		replyMsg = &discordgo.MessageSend{Content: rule.ReplyFunc(bot, channelID, msg.ID)}
	} else if rule.ReplyText != "" {
		// Deprecated.
		fmt.Println("ReplyRule.ReplyText is deprecated. Use ReplyMsg instead.")
		replyMsg = &discordgo.MessageSend{Content: rule.ReplyText}
	} else {
		return
	}

	// Some error occured.
	if replyMsg == nil {
		return
	}

	// Change the target to reply to.
	if rule.ReplyToInitialMessage {
		// Should not reply to another's message in DM.
		// Will cause chaos.
		if rule.ReplyInDM {
			return
		}
		if msg.ReferencedMessage != nil {
			replyMsg.Reference = msg.ReferencedMessage.Reference()
		}
	}

	// Whether should reply in DM.
	if rule.ReplyInDM {
		// DM message has no reference msg.
		replyMsg.Reference = nil
		bot.SendDM(msg.Author.ID, replyMsg)
	} else {
		// If not replying to anything, should send the msg in ReplyMsgFunc.
		if replyMsg.Reference == nil {
			replyMsg.Reference = msg.Reference()
		}
		bot.Session.ChannelMessageSendComplex(channelID, replyMsg)
	}
}

// Start the reply and never stop.
func (bot *DiscordBot) StartReply() {
	defer utils.Restart(bot.StartReply)
	// Update last read msgs map first.
	err := bot.UpdateLastReadMsgs()
	if err != nil {
		panic(err)
	}
	fmt.Println(bot.BotName, "start to listen and reply")
	for {
		for channel := range bot.lastRead {
			msgs, err := bot.readLatestMsgs(channel, 100)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, msg := range msgs {
				for _, rule := range bot.replyRules {
					if rule.shouldReply(msg) {
						bot.reply(channel, msg, rule)
					}
				}
				for _, rule := range bot.reactRules {
					if rule.shouldReact(msg) {
						if rule.ReactFunc == nil {
							bot.react(channel, msg.ID, rule.ReactEmojiIDs)
						} else {
							bot.react(channel, msg.ID, rule.ReactFunc)
						}
					}
				}
			}
		}
		time.Sleep(bot.replyFrequency)
	}
}

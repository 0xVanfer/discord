package discord

import (
	"fmt"
	"time"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Reply to a msg.
func (bot *DiscordBot) reply(channelID string, msg *discordgo.Message, rule ReplyRule) {
	var replyText string
	if rule.ReplyFunc == nil {
		replyText = rule.ReplyText
	} else {
		replyText = rule.ReplyFunc(bot, channelID, msg.ID)
	}
	if replyText == "" {
		return
	}
	var replyMsgID string = msg.ID
	// Change the target to reply to.
	if rule.ReplyToInitialMessage {
		// Should not reply to another's message in DM.
		// Will cause chaos.
		if rule.ReplyInDM {
			return
		}
		if msg.ReferencedMessage != nil {
			replyMsgID = msg.ReferencedMessage.ID
		}
	}
	// Whether should reply in DM.
	if rule.ReplyInDM {
		bot.SendDM(msg.Author.ID, replyText)
	} else {
		bot.Session.ChannelMessageSendReply(channelID, replyText, &discordgo.MessageReference{MessageID: replyMsgID, ChannelID: channelID})
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
			msgs, err := bot.readLastMsgs(channel, 10)
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

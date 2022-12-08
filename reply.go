package discord

import (
	"fmt"
	"time"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Reply to a msg.
//
// Input:
//
//	function: Must be string or func(*DiscordBot, string, string) string.
func (bot *DiscordBot) reply(channelID string, msgID string, userID string, replyInDM bool, function any) {
	var replyText string
	switch v := function.(type) {
	case string:
		replyText = v
	case func(*DiscordBot, string, string) string:
		replyText = v(bot, channelID, msgID)
	}
	if replyText == "" {
		return
	}
	if replyInDM {
		bot.SendDM(userID, replyText)
	} else {
		bot.Session.ChannelMessageSendReply(channelID, replyText, &discordgo.MessageReference{MessageID: msgID, ChannelID: channelID})
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
						if rule.ReplyFunc == nil {
							bot.reply(channel, msg.ID, msg.Author.ID, rule.ReplyInDM, rule.ReplyText)
						} else {
							bot.reply(channel, msg.ID, msg.Author.ID, rule.ReplyInDM, rule.ReplyFunc)
						}
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

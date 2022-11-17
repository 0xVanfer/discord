package discord

import (
	"fmt"
	"time"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Reply to a msg.
func (bot *DiscordBot) reply(channel string, msgId string, function any) {
	var replyText string
	switch v := function.(type) {
	case string:
		replyText = v
	case func(*DiscordBot, string, string) string:
		replyText = v(bot, channel, msgId)
	}
	if replyText == "" {
		return
	}
	// Will not reply to bot msg.
	msg, err := bot.Session.ChannelMessage(channel, msgId)
	if (err == nil) && msg.Author.Bot {
		return
	}
	bot.Session.ChannelMessageSendReply(channel, replyText, &discordgo.MessageReference{MessageID: msgId, ChannelID: channel})
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
							bot.reply(channel, msg.ID, rule.ReplyText)
						} else {
							bot.reply(channel, msg.ID, rule.ReplyFunc)
						}
					}
				}
			}
		}
		time.Sleep(bot.replyFrequency)
	}
}

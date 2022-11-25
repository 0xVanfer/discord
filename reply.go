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
func (bot *DiscordBot) reply(channelID string, msgID string, function any) {
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
	bot.Session.ChannelMessageSendReply(channelID, replyText, &discordgo.MessageReference{MessageID: msgID, ChannelID: channelID})
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
				for _, rule := range bot.reactRules {
					if rule.shouldReact(msg) {
						for _, enojiID := range rule.ReactEmojiIDs {
							bot.react(channel, msg.ID, enojiID)
						}
					}
				}
			}
		}
		time.Sleep(bot.replyFrequency)
	}
}

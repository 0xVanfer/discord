package discord

import (
	"fmt"
	"time"

	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Reply to a msg.
func (bot *DiscordBot) reply(channel string, msgId string, text string) {
	bot.Session.ChannelMessageSendReply(channel, text, &discordgo.MessageReference{MessageID: msgId, ChannelID: channel})
}

// Start the reply and never stop.
func (bot *DiscordBot) StartReply() {
	fmt.Println("start to listen and reply")
	defer utils.Restart(bot.StartReply)
	// Update last read msgs map first.
	err := bot.UpdateLastReadMsgs()
	if err != nil {
		panic(err)
	}
	for {
		for channel := range bot.LastRead {
			msgs, err := bot.readLastMsgs(channel, 10)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, msg := range msgs {
				for _, rule := range bot.ReplyRules {
					if rule.shouldReply(msg) {
						bot.reply(channel, msg.ID, rule.ReplyText)
					}
				}
			}
		}
		time.Sleep(bot.ReplyFrequency)
	}
}

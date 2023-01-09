package discord

import "github.com/bwmarrin/discordgo"

// Update msg info in "bot.lastRead".
func (bot *DiscordBot) UpdateLastReadMsgs() error {
	for channel := range bot.lastRead {
		res, err := bot.Session.ChannelMessages(channel, 1, "", "", "")
		if err != nil {
			return err
		}
		bot.lastRead[channel] = msgInfo{MsgID: res[0].ID, SendAt: res[0].Timestamp}
	}
	return nil
}

// Read the last "amount" of msgs of a channel.
func (bot *DiscordBot) readLastMsgs(channel string, amount int) ([]*discordgo.Message, error) {
	afterId := bot.lastRead[channel].MsgID
	res, err := bot.Session.ChannelMessages(channel, amount, "", afterId, "")
	if len(res) == 0 {
		return res, err
	}
	bot.lastRead[channel] = msgInfo{MsgID: res[0].ID, SendAt: res[0].Timestamp}
	return res, nil
}

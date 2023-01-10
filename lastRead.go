package discord

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

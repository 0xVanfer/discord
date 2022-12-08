package discord

func (bot *DiscordBot) SendDM(userID string, text string) error {
	newChannel, err := bot.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}
	_, err = bot.Session.ChannelMessageSend(newChannel.ID, text)
	return err
}

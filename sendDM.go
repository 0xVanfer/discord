package discord

import "github.com/bwmarrin/discordgo"

func (bot *DiscordBot) SendDM(userID string, msg *discordgo.MessageSend) error {
	newChannel, err := bot.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}
	_, err = bot.Session.ChannelMessageSendComplex(newChannel.ID, msg)
	return err
}

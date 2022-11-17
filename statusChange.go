package discord

import "errors"

// Change the bot's nicknames in the guilds.
func (bot *DiscordBot) ChangeNames(nickname string, guilds ...string) error {
	if nickname == "" {
		return errors.New("nickname must not be empty")
	}
	for _, guild := range guilds {
		err := bot.Session.GuildMemberNickname(guild, "@me", nickname)
		if err != nil {
			return err
		}
	}
	return nil
}

// Change the bot's "playing" status.
func (bot *DiscordBot) ChangePlaying(gameName string) error {
	err := bot.Session.UpdateGameStatus(1, gameName)
	if err != nil {
		return err
	}
	return nil
}

// Change the bot's "listening" status.
func (bot *DiscordBot) ChangeListening(listenTo string) error {
	err := bot.Session.UpdateListeningStatus(listenTo)
	if err != nil {
		return err
	}
	return nil
}

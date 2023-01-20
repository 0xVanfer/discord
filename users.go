package discord

import "github.com/bwmarrin/discordgo"

// Read all members of a guild.
func (bot *Bot) AllMembers(guildID string) (allMembers []*discordgo.Member, err error) {
	lastID := ""
	for {
		newMembers, err := bot.Session.GuildMembers(guildID, lastID, 1000)
		if err != nil {
			return nil, err
		}
		allMembers = append(allMembers, newMembers...)
		if len(newMembers) < 1000 {
			return allMembers, nil
		}
		lastID = newMembers[len(newMembers)-1].User.ID
	}
}

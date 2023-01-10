package discord

import "github.com/0xVanfer/utils"

// Check if the user is the role of the guild.
func (bot *DiscordBot) CheckUserRole(userID string, guildID string, roleID string) (isRole bool) {
	member, err := bot.Session.GuildMember(guildID, userID)
	if err != nil {
		return false
	}
	return utils.ContainInArray(roleID, member.Roles)
}

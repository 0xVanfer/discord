package discord

import (
	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Check if the user is the role of the guild.
func (bot *DiscordBot) CheckUserRole(userID string, guildID string, roleID string) (isRole bool) {
	member, err := bot.Session.GuildMember(guildID, userID)
	if err != nil {
		return false
	}
	return utils.ContainInArray(roleID, member.Roles)
}

// Read all members of a role.
func (bot *DiscordBot) FilterRole(guildID string, roleID string) (roleMembers []*discordgo.Member, err error) {
	allMembers, err := bot.AllMembers(guildID)
	if err != nil {
		return
	}
	for _, member := range allMembers {
		if utils.ContainInArray(roleID, member.Roles) {
			roleMembers = append(roleMembers, member)
		}
	}
	return
}

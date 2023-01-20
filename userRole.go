package discord

import (
	"github.com/0xVanfer/utils"
	"github.com/bwmarrin/discordgo"
)

// Check if the user is the role of the guild.
func (bot *Bot) CheckUserRole(userID string, guildID string, roleID string) (isRole bool) {
	member, err := bot.Session.GuildMember(guildID, userID)
	if err != nil {
		return false
	}
	return utils.ContainInArray(roleID, member.Roles)
}

// Deprecated: Use FilterRoles instead.
// Read all members of a role.
func (bot *Bot) FilterRole(guildID string, roleID string) (roleMembers []*discordgo.Member, err error) {
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

// Read all members of some roles.
func (bot *Bot) FilterRoles(guildID string, roleIDs ...string) (roleMembers []*discordgo.Member, err error) {
	allMembers, err := bot.AllMembers(guildID)
	if err != nil {
		return
	}
	for _, member := range allMembers {
		for _, roleID := range roleIDs {
			if utils.ContainInArray(roleID, member.Roles) {
				roleMembers = append(roleMembers, member)
			}
		}
	}
	roleMembers = utils.RemoveRepetitionInArray(roleMembers)
	return
}

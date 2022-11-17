package discord

import "github.com/bwmarrin/discordgo"

// Return the text to mention a user.
func MentionUserText(user *discordgo.User) string {
	return "<@" + user.ID + ">"
}

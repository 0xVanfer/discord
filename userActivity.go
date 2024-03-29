package discord

import "github.com/bwmarrin/discordgo"

// Struct for user activities.
type UserActivity struct {
	User    *discordgo.User
	Times   int64
	Details []*discordgo.Message
}

// Read all the msgs of selected channels.
//
// startTimestamp and endTimestamp can either be 0 for no start or end limits.
//
// GuildID is only used when channelIDs is nil, read all the channels of the guild.
//
// res: map[userID] = UserActivity
func (bot *Bot) ReadUserActivities(guildID string, channelIDs []string, startTimestamp int64, endTimestamp int64) (res map[string]*UserActivity, err error) {
	// Read the messages.
	allMsgs, err := bot.ReadGuildMsgs(guildID, channelIDs, startTimestamp, endTimestamp)
	if err != nil {
		return res, err
	}

	mapp := make(map[string]*UserActivity)
	for _, msg := range allMsgs {
		// Username will change, but ID will not.
		userID := msg.Author.ID
		if mapp[userID] == nil {
			mapp[userID] = &UserActivity{
				User:    msg.Author,
				Times:   1,
				Details: []*discordgo.Message{msg},
			}
		} else {
			mapp[userID].Times += 1
			mapp[userID].Details = append(mapp[userID].Details, msg)
		}
	}
	return mapp, nil
}

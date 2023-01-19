package discord

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// Read all the msgs of a channel.
//
// startTimestamp and endTimestamp can either be 0 for no start or end limits.
func (bot *DiscordBot) readChannelMsgs(channelID string, startTimestamp int64, endTimestamp int64) (res []*discordgo.Message, err error) {
	// Neither 0, startTimestamp should smaller than endTime.
	if startTimestamp != 0 && endTimestamp != 0 {
		if startTimestamp > endTimestamp {
			return res, errors.New("start time should be smaller than end time")
		}
	}
	// The first time can be empty. Read the latest msgs.
	beforeId := ""
	for {
		// Read 100 more msgs.
		newRes, err := bot.Session.ChannelMessages(channelID, 100, beforeId, "", "")
		if err != nil {
			return res, err
		}
		for _, msg := range newRes {
			msgTime := msg.Timestamp.Unix()
			// Time of the msg must before endTime.
			if endTimestamp != 0 && msgTime > endTimestamp {
				continue
			}
			// Time of the msg must after endTime.
			// If startTime is not 0, and already find a record before
			// startTime, should return.
			if startTimestamp != 0 && msgTime < startTimestamp {
				return res, nil
			}
			res = append(res, msg)
		}
		// If length less than 100, already read all msgs.
		if len(newRes) < 100 {
			return res, err
		}
		// Update the "beforeId".
		beforeId = newRes[99].ID
	}
}

// Read all the msgs of a guild.
// If only one channel needed, can also use this function.
//
// startTimestamp and endTimestamp can either be 0 for no start or end limits.
//
// If channelIDs is nil, read all the channels of the guild.
// If not, use channelIDs as all channels.
func (bot *DiscordBot) ReadGuildMsgs(guildID string, channelIDs []string, startTimestamp int64, endTimestamp int64) (res []*discordgo.Message, err error) {
	// Decide channelIDs.
	if channelIDs == nil {
		channels, err := bot.Session.GuildChannels(guildID)
		if err != nil {
			return res, err
		}
		for _, channel := range channels {
			channelIDs = append(channelIDs, channel.ID)
		}
	}
	// For all the channels.
	for _, channelID := range channelIDs {
		newRes, err := bot.readChannelMsgs(channelID, startTimestamp, endTimestamp)
		if err != nil {
			return res, err
		}
		res = append(res, newRes...)
	}
	return res, nil
}

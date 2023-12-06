package commands

import (
	database "boonsboos.nl/boonsu/Database"
	osu "boonsboos.nl/boonsu/Osu"
	"github.com/bwmarrin/discordgo"
)

func osuLinkCommand(s *discordgo.Session, m *discordgo.Message, c []string) {
	if len(c) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You need to tell me your osu username too")
		return
	}

	user, err := osu.GetUserByUsername(c[1])
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Couldn't find that user")
		return
	}

	err = database.SaveOsuDiscordLink(m.Author.ID, user.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Something went wrong, try again later")
	} else {
		s.ChannelMessageSend(m.ChannelID, "Linked you to "+user.Username)
	}
}

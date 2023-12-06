package commands

import (
	"log"
	"strconv"
	"strings"

	database "boonsboos.nl/boonsu/Database"
	osu "boonsboos.nl/boonsu/Osu"
	"github.com/bwmarrin/discordgo"
)

var profileEmbed discordgo.MessageEmbed = discordgo.MessageEmbed{
	Color:  0xabd5ed,
	Fields: []*discordgo.MessageEmbedField{},
}

func osuProfileCommand(s *discordgo.Session, m *discordgo.Message, c []string) {

	var osuID int = 0

	if len(c) > 1 && len(m.Mentions) > 0 && strings.HasPrefix(c[1], "<@") {
		i, err := database.GetOsuIDFromDiscordID(m.Mentions[0].ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "That user isn't linked! tell them to link with `monke linkosu {yourOsuUsername}`")
			return
		}
		osuID = i
	} else {
		i, err := database.GetOsuIDFromDiscordID(m.Author.ID)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Please link first with `monke linkosu {yourOsuUsername}`")
			return
		}
		osuID = i
	}

	profile, err := osu.GetUserByID(osuID)
	if err != nil {
		log.Println("user for " + strconv.Itoa(osuID) + " not found")
		s.ChannelMessageSend(m.ChannelID, "Could not get your profile")
		return
	}

	profileEmbed.Fields = []*discordgo.MessageEmbedField{
		{
			Name: "Rank",
			Value: "#" + strconv.FormatInt(profile.Statistics.GlobalRank, 10) +
				" (#" + strconv.FormatInt(profile.Statistics.Rank.Country, 10) +
				" " + profile.Country + ")",
			Inline: true,
		},
		{
			Name:   "PP",
			Value:  strconv.FormatFloat(float64(profile.Statistics.PP), 'f', 0, 32),
			Inline: true,
		},
		{
			Name:   "Playcount",
			Value:  strconv.FormatInt(profile.Statistics.PlayCount, 10),
			Inline: true,
		},
	}

	profileEmbed.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: "https://a.ppy.sh/" + strconv.Itoa(osuID),
	}

	profileEmbed.Title = profile.Username

	s.ChannelMessageSendEmbed(m.ChannelID, &profileEmbed)
}

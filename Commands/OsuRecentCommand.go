package commands

import (
	"log"
	"strconv"
	"strings"

	database "boonsboos.nl/boonsu/Database"
	osu "boonsboos.nl/boonsu/Osu"
	"github.com/bwmarrin/discordgo"
)

func osuRecentCommand(s *discordgo.Session, m *discordgo.Message, c []string) {
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

	score, err := osu.GetMostRecentScoreFromUser(osuID)
	if err != nil {
		log.Println("user for " + strconv.Itoa(osuID) + " not found")
		s.ChannelMessageSend(m.ChannelID, "Could not get your score")
		return
	}

	beatmap, err := osu.GetBeatmap(score.Map.MapID)
	if err != nil {
		log.Println("map " + strconv.Itoa(osuID) + " not found")
		s.ChannelMessageSend(m.ChannelID, "Could not find the map")
		return
	}

	embed := discordgo.MessageEmbed{
		Title: score.GetMapInfo() + " +" + score.GetMods(),
		URL:   score.Map.URL,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: osu.BeatmapImageLink(score.Map.MapsetID),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Played by " + score.Player.Username,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Value:  score.GetHitInfo(),
				Inline: true,
			},
			{
				Value:  strconv.FormatFloat(score.PP, 'f', 2, 64) + "pp",
				Inline: true,
			},
			{
				Value: strconv.Itoa(score.MaxCombo) + "/" + strconv.Itoa(beatmap.MaxCombo) + "x",
			},
			{
				Value:  strconv.FormatFloat(score.Accuracy*100, 'f', 2, 64) + "%",
				Inline: true,
			},
			{
				Value:  score.Rank,
				Inline: true,
			},
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

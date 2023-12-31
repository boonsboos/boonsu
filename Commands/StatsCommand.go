package commands

import (
	"log"
	"strconv"

	database "boonsboos.nl/boonsu/Database"
	util "boonsboos.nl/boonsu/Util"
	"github.com/bwmarrin/discordgo"
)

var statsCommandEmbed discordgo.MessageEmbed = discordgo.MessageEmbed{
	Title:  "stats",
	Color:  0xacceb9,
	Fields: []*discordgo.MessageEmbedField{},
}

func statsCommand(s *discordgo.Session, m *discordgo.Message, c []string) {

	statsCommandEmbed.Fields = []*discordgo.MessageEmbedField{
		{
			Name:   "version",
			Value:  util.Version,
			Inline: true,
		},
		{
			Name:   "commands ran",
			Value:  strconv.FormatUint(database.GetCommandsRan(), 10),
			Inline: true,
		},
		{
			Name:   "database size",
			Value:  database.GetDatabaseSize(),
			Inline: true,
		},
		{
			Name:   "tillerino version",
			Value:  "shooting for compatibility with `bc40b3d`",
			Inline: true,
		},
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &statsCommandEmbed)
	if err != nil {
		log.Fatal("failed to make stats")
	}
}

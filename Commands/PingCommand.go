package commands

import "github.com/bwmarrin/discordgo"

func pingCommand(s *discordgo.Session, a *discordgo.Message, b []string) {
	s.ChannelMessageSend(a.ChannelID, "bing")
}

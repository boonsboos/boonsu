package commands

import (
	"strings"

	database "boonsboos.nl/boonsu/Database"
	"github.com/bwmarrin/discordgo"
)

func DispatchCommands(sesh *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot {
		return
	}

	if !strings.HasPrefix(e.Message.Content, "monke") {
		return
	}

	message := strings.TrimSpace(strings.TrimPrefix(e.Message.Content, "monke"))
	if len(message) == 0 {
		return
	}

	command := strings.Split(message, " ")
	executor := all_commands[command[0]]

	go executor(sesh, e.Message, command)
	// idk if making a goroutine here has any performance benefits
	// but i'll do it anyway

	go database.UpdateRanStats() // increases the counter by one
}

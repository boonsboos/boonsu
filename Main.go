package main

import (
	"log"
	"os"
	"os/signal"

	commands "boonsboos.nl/boonsu/Commands"
	osu "boonsboos.nl/boonsu/Osu"
	util "boonsboos.nl/boonsu/Util"
	"github.com/bwmarrin/discordgo"
)

func main() {

	go osu.AutoReAuth()

	session, err := discordgo.New("Bot " + util.Options.DiscordToken)
	if err != nil {
		log.Fatal("Cannot make Discord connection: " + err.Error())
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Ready")
	})

	session.AddHandler(commands.DispatchCommands)

	session.Open()

	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down")
}

package util

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func init() {
	getOptionsFile()
}

var Options BoonsuOptions

type BoonsuOptions struct {
	DiscordToken string `json:"discordToken"`
	OsuToken     string `json:"osuToken"`
	OsuClientID  string `json:"osuClientID"`
	DatabaseURL  string `json:"databaseURL"`
}

func getOptionsFile() {
	a, err := os.Open("options.json")
	if err != nil {
		log.Fatal("Options not found.")
	}

	data, err := io.ReadAll(a)
	if err != nil {
		log.Fatal("Failed to read file.")
	}

	json.Unmarshal(data, &Options)
}

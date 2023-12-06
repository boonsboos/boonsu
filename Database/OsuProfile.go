package database

import (
	"errors"
	"log"
)

func GetOsuIDFromDiscordID(discordID string) (int, error) {
	statement, err := BoonsuDB.db.Prepare("SELECT osu_id FROM osu_discord_linked WHERE discord_id LIKE $1")
	if err != nil {
		log.Println("Failed to prepare statement getting osu ID:", err.Error())
		return 0, errors.New("SQL error")
	}

	var result int = 0

	err = statement.QueryRow(discordID).Scan(&result)
	if err != nil {
		log.Println("Failed to execute query getting osu ID by discord ID: " + err.Error())
		return 0, errors.New("not found?")
	}

	statement.Close()

	return result, nil
}

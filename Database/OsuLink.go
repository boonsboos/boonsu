package database

import (
	"errors"
	"log"
)

func SaveOsuDiscordLink(discordId string, osuId int) error {
	statement, err := BoonsuDB.db.Prepare(
		`INSERT INTO osu_discord_linked (discord_id, osu_id)
		VALUES ( $1 , $2 ) 
		ON CONFLICT (discord_id) DO
		UPDATE SET osu_id = EXCLUDED.osu_id
		WHERE EXCLUDED.discord_id = osu_discord_linked.discord_id`, // profesnal
	)
	if err != nil {
		log.Println("SaveOsuDiscordLink | failed to prepare statement:", err.Error())
		return errors.New("")
	}

	_, err = statement.Exec(discordId, osuId)
	if err != nil {
		log.Println("Failed to execute query inserting discord <-> osu link: " + err.Error())
		return errors.New("")
	}

	statement.Close()
	return nil
}

package database

import "log"

// this package contains common queries related to the bot.
func GetDatabaseSize() string {
	statement, err := BoonsuDB.db.Prepare("SELECT pg_size_pretty( pg_database_size('boonsu') )")
	if err != nil {
		log.Fatal("Failed to prepare statement:", err.Error())
	}

	var result string

	err = statement.QueryRow().Scan(&result)
	if err != nil {
		log.Fatal("Failed to execute query getting size of database: " + err.Error())
	}

	statement.Close()

	return result
}

func GetCommandsRan() uint64 {
	statement, err := BoonsuDB.db.Prepare("SELECT ran FROM ran") // TODO: change table name back to commands_ran
	if err != nil {
		log.Fatal("Failed to prepare statement:", err.Error())
	}

	var result uint64

	err = statement.QueryRow().Scan(&result)
	if err != nil {
		log.Fatal("Failed to execute query getting amount of commands ran: " + err.Error())
	}

	statement.Close()

	return result
}

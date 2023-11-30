package database

import "log"

func UpdateRanStats() {
	statement, err := BoonsuDB.db.Prepare("UPDATE ran SET ran = ran + 1")
	if err != nil {
		log.Fatal("Failed to prepare statement:", err.Error())
	}

	_, err = statement.Exec()
	if err != nil {
		log.Println("Failed to execute query updating the amount of commands ran: " + err.Error())
	}

	statement.Close()
}

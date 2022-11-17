package migration

import (
	"database/sql"
	"log"
)

func UserLoginMigration(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `user_login` (" +
		"`id` VARCHAR(100) NOT NULL," +
		"`id_user` VARCHAR(10) NULL DEFAULT NULL," +
		"`status` VARCHAR(1) NULL DEFAULT NULL," +
		"`token` VARCHAR(255) NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[USER LOGIN - TABLE] berhasil dibuat")
	}

}

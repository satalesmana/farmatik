package migration

import (
	"database/sql"
	"log"
)

func AppConfig(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `app_config` (" +
		"`id` VARCHAR(30) NOT NULL," +
		"`keterangan` VARCHAR(30) NULL DEFAULT NULL," + //Umum, Bidan
		"`configValue` JSON NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[APP CONFIG - TABLE] berhasil dibuat")
	}

}

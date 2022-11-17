package migration

import (
	"database/sql"
	"log"
)

func ProductMigration(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `product` (" +
		"`id` INT NOT NULL AUTO_INCREMENT," +
		"`namaProduct` VARCHAR(50) NULL DEFAULT NULL," +
		"`hargaBeli` INT(10) NULL DEFAULT NULL," +
		"`satuan` VARCHAR(50) NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[PRODUCT - TABLE] berhasil dibuat")
	}

}

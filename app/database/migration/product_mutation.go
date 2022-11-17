package migration

import (
	"database/sql"
	"log"
)

func ProductMutation(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `product_mutation` (" +
		"`id` INT NOT NULL AUTO_INCREMENT," +
		"`productId` VARCHAR(30) NULL DEFAULT NULL," +
		"`trxCode` VARCHAR(30) NULL DEFAULT NULL," +
		"`createdBy` VARCHAR(30) NULL DEFAULT NULL," +
		"`createdDate` DATE NULL DEFAULT NULL," +
		"`type` VARCHAR(1) NULL DEFAULT NULL," + // I, O
		"`value` INT NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[PRODUCT MUTATION - TABLE] berhasil dibuat")
	}

}

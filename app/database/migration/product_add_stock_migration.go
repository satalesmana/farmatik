package migration

import (
	"database/sql"
	"log"
)

func ProductAddStock(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `product_add_stock` (" +
		"`id` VARCHAR(10) NULL DEFAULT NULL," +
		"`usreId` VARCHAR(30) NULL DEFAULT NULL," +
		"`productId` VARCHAR(10) NULL DEFAULT NULL," +
		"`createdBy` VARCHAR(30) NULL DEFAULT NULL," +
		"`createdDate` DATE NULL DEFAULT NULL," +
		"`value` INT NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[PRODUCT ADD STOCK - TABLE] berhasil dibuat")
	}
}

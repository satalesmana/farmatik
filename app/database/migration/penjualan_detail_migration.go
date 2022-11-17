package migration

import (
	"database/sql"
	"log"
)

func PenjualanDetailMigration(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `penjualan_detail` (" +
		"`id` INT NOT NULL AUTO_INCREMENT," +
		"`penjualanId` VARCHAR(10) NOT NULL," +
		"`productId` VARCHAR(10) NOT NULL," + //Umum, Bidan
		"`harga` INT NULL DEFAULT 0," +
		"`qty` INT NULL DEFAULT 0," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[HARGA JUAL PRODUCT - TABLE] berhasil dibuat")
	}

}

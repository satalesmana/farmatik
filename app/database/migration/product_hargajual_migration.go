package migration

import (
	"database/sql"
	"log"
)

func ProductHargaJualMigration(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `product_hargajual` (" +
		"`id` INT NOT NULL AUTO_INCREMENT," +
		"`idProduct` INT(10) NOT NULL," +
		"`kategori` VARCHAR(30) NULL DEFAULT NULL," + //Umum, Bidan
		"`harga` INT(10) NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[HARGA JUAL PRODUCT - TABLE] berhasil dibuat")
	}

}

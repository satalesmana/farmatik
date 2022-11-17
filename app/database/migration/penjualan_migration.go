package migration

import (
	"database/sql"
	"log"
)

func PenjualanMigration(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE `penjualan` (" +
		"`id` VARCHAR(10) NOT NULL," +
		"`createdBy` VARCHAR(70) NULL DEFAULT NULL," +
		"`createdDate` DATE NULL DEFAULT NULL," +
		"`namaPelanggan` VARCHAR(70) NULL DEFAULT NULL," +
		"`tipePelanggan` VARCHAR(50) NULL DEFAULT NULL," +
		"`trxType` VARCHAR(50) NULL DEFAULT NULL," +
		"`docPendukung` VARCHAR(100) NULL DEFAULT NULL," +
		"PRIMARY KEY (`id`))")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("[HARGA JUAL PRODUCT - TABLE] berhasil dibuat")
	}

}

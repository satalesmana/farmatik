package configSeeder

import (
	"database/sql"
	db "farmatik/app/database"
	"log"
	"strconv"
)

type Handler interface {
	AppConfigSeeder()
}

type uscase struct {
	database *sql.DB
}

func NewSeederHandler() Handler {
	return &uscase{
		database: db.GetCoon(),
	}
}

func (uc *uscase) AppConfigSeeder() {
	query := `REPLACE INTO app_config (id, keterangan, configValue)  VALUES 
	('cfg_harga_jual', 'Config Harga Jual', '[{"value": "cfg_hrg_umum", "label": "Umum"}, {"value": "cfg_hrg_bidan", "label": "Tenaga Medis (Bidan)"}]'),
	('cfg_type_trx', 'Config Tipe Transaksi', '[{"value": "cfg_trx_tunai", "label": "Tunai"}, {"value": "cfg_trx_nontunai", "label": "Non Tunai"}]')`

	res, err := uc.database.Exec(query)

	if err != nil {
		log.Println(err)
	}

	intRes, errAft := res.RowsAffected()
	if err != nil {
		log.Println(errAft)
	}

	strIDUser := strconv.Itoa(int(intRes))
	log.Println("[APP CONFIG SEEDER] " + strIDUser + " Rows Affected ")

}

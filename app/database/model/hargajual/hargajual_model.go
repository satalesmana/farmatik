package hargajualModel

import (
	"database/sql"
	db "farmatik/app/database"
)

type Handler interface {
	Insert(data *HargaJual) (int64, error)
	GetByidProduct(idProduct string) ([]HargaJual, error)
}

type HargaJual struct {
	ID        int64  `json:"id,omitempty"`
	IdProduct string `json:"id_product,omitempty"`
	Kategori  string `json:"kategori,omitempty"`
	Harga     int64  `json:"harga,omitempty"`
}

type Db struct {
	database *sql.DB
}

func NewHargaJualHandler() Handler {
	return &Db{database: db.GetCoon()}
}

func (db *Db) Insert(data *HargaJual) (int64, error) {
	query := `INSERT INTO product_hargajual(
		idProduct, kategori, harga
	) VALUES(?, ?, ?)`

	res, err := db.database.Exec(query,
		&data.IdProduct,
		&data.Kategori,
		&data.Harga,
	)

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (db *Db) GetByidProduct(idProduct string) ([]HargaJual, error) {
	query := `SELECT id, idProduct, kategori, harga FROM product_hargajual`
	res, err := db.database.Query(query)
	var output []HargaJual
	if err != nil {
		return output, err
	}
	defer res.Close()

	for res.Next() {
		var data HargaJual
		if err := res.Scan(
			&data.ID,
			&data.IdProduct,
			&data.Kategori,
			&data.Harga,
		); err != nil {
			return output, err
		}

		output = append(output, data)
	}

	return output, nil
}

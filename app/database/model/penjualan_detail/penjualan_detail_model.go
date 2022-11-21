package PenjualanDetailModel

import (
	"database/sql"
	db "farmatik/app/database"
)

type Handler interface {
	Insert(data *PenjualanDetail) (int64, error)
	GetByTrx(PenjualanId string) ([]PenjualanDetail, error)
}

type PenjualanDetail struct {
	ID          int64  `json:"id,omitempty"`
	PenjualanId string `json:"penjualan_id,omitempty"`
	ProductId   string `json:"product_id,omitempty"`
	Harga       int64  `json:"harga,omitempty"`
	Qty         int64  `json:"qty,omitempty"`
}

type uscase struct {
	database *sql.DB
}

func NewPenjualanDetailHandler() Handler {
	return &uscase{database: db.GetCoon()}
}

func (uc *uscase) Insert(data *PenjualanDetail) (int64, error) {
	query := `INSERT INTO penjualan_detail(
		penjualanId, productId, harga, qty
	) VALUES(?, ?, ?, ?)`

	res, err := uc.database.Exec(query,
		&data.PenjualanId,
		&data.ProductId,
		&data.Harga,
		&data.Qty,
	)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (uc *uscase) GetByTrx(PenjualanId string) ([]PenjualanDetail, error) {
	var penjualanDetail []PenjualanDetail

	query := `SELECT id, penjualanId, productId, harga, qty 
		FROM penjualan_detail WHERE penjualanId=?`

	res, err := uc.database.Query(query, PenjualanId)
	if err != nil {
		return penjualanDetail, err
	}
	defer res.Close()

	for res.Next() {
		var item PenjualanDetail
		if err := res.Scan(
			&item.ID,
			&item.PenjualanId,
			&item.ProductId,
			&item.Harga,
			&item.Qty,
		); err != nil {
			return penjualanDetail, err
		}
		penjualanDetail = append(penjualanDetail, item)
	}

	return penjualanDetail, nil
}

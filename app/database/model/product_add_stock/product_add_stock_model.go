package productaddstock

import (
	"database/sql"
	db "farmatik/app/database"
	DBHelper "farmatik/app/database/helper"
)

type ProductAddStock struct {
	ID          string `json:"id,omitempty"`
	UsreId      string `json:"user_id,omitempty"`
	ProductId   string `json:"product_id,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	Value       int64  `json:"value,omitempty"`
}

type uscase struct {
	database *sql.DB
}

type Handler interface {
	Insert(data *ProductAddStock) (*ProductAddStock, error)
}

func NewProductStockHandler() Handler {
	return &uscase{
		database: db.GetCoon(),
	}
}

func (uc *uscase) Insert(data *ProductAddStock) (*ProductAddStock, error) {
	query := `INSERT INTO product_add_stock(
		id, usreId, productId, createdBy, createdDate, value  
	) VALUES(?, ?, ?, ?, ?, ?)`

	//generate trx number
	trxId, errTrxId := DBHelper.GenerateAutoId("id", "STK", "product_add_stock")
	if errTrxId != nil {
		return data, errTrxId
	}
	data.ID = trxId

	_, err := uc.database.Exec(query,
		&data.ID,
		&data.UsreId,
		&data.ProductId,
		&data.CreatedBy,
		&data.CreatedDate,
		&data.Value,
	)

	if err != nil {
		return data, err
	}

	return data, nil
}

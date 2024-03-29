package productmutation

import (
	"database/sql"
	db "farmatik/app/database"
)

type Handler interface {
	Insert(data *ProductMutation) (int64, error)
}

type ProductMutation struct {
	ID          int64  `json:"id,omitempty"`
	ProductId   string `json:"id_product,omitempty"`
	TrxCode     string `json:"trx_code,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	Type        string `json:"type,omitempty"`
	Value       int64  `json:"value,omitempty"`
	Keterangan  string `json:"keterangan,omitempty"`
}

type uscase struct {
	database *sql.DB
}

func NewProductMutationHandler() Handler {
	return &uscase{
		database: db.GetCoon(),
	}
}

func (uc *uscase) Insert(data *ProductMutation) (int64, error) {
	query := `INSERT INTO product_mutation(
		productId, trxCode, createdBy, createdDate,type,value,keterangan  
	) VALUES(?, ?, ?, ?, ?, ?, ?)`

	res, err := uc.database.Exec(query,
		&data.ProductId,
		&data.TrxCode,
		&data.CreatedBy,
		&data.CreatedDate,
		&data.Type,
		&data.Value,
		&data.Keterangan,
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

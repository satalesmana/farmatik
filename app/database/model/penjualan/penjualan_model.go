package PenjualanModel

import (
	"database/sql"
	db "farmatik/app/database"
	DBHelper "farmatik/app/database/helper"
	PenjualanDetailModel "farmatik/app/database/model/penjualan_detail"
)

type Handler interface {
	Insert(data *Penjualan) (string, error)
	GetById(id string) (Penjualan, error)
	GetAll() ([]Penjualan, error)
}

type Penjualan struct {
	ID              string                                  `json:"id,omitempty"`
	CreatedBy       string                                  `json:"created_by,omitempty"`
	CreatedDate     string                                  `json:"created_date,omitempty"`
	NamaPelanggan   string                                  `json:"nama_pelanggan,omitempty"`
	TipePelanggan   string                                  `json:"tipe_pelanggan,omitempty"`
	TrxType         string                                  `json:"tipe_trx,omitempty"`
	DocPendukung    string                                  `json:"doc_pendukung,omitempty"`
	PenjualanDetail []*PenjualanDetailModel.PenjualanDetail `json:"penjualan_detail,omitempty"`
}

type uscase struct {
	database        *sql.DB
	penjualanDetail PenjualanDetailModel.Handler
}

func NewPenjualanHandler() Handler {
	return &uscase{
		database:        db.GetCoon(),
		penjualanDetail: PenjualanDetailModel.NewPenjualanDetailHandler(),
	}
}

func (uc *uscase) Insert(data *Penjualan) (string, error) {
	query := `INSERT INTO penjualan(
		id, createdBy, createdDate, namaPelanggan, tipePelanggan, trxType, docPendukung
	) VALUES(?, ?, ?, ?, ?, ?, ?)`

	//generate trx number
	trxId, errTrxId := DBHelper.GenerateAutoId("id", "TRX", "penjualan")
	if errTrxId != nil {
		return "", errTrxId
	}
	data.ID = trxId

	_, err := uc.database.Exec(query,
		&data.ID,
		&data.CreatedBy,
		&data.CreatedDate,
		&data.NamaPelanggan,
		&data.TipePelanggan,
		&data.TrxType,
		&data.DocPendukung,
	)

	if err != nil {
		return "", err
	}

	return data.ID, nil
}

func (uc *uscase) GetById(id string) (Penjualan, error) {
	var penjualan Penjualan

	query := `SELECT id, createdBy, createdDate, namaPelanggan, tipePelanggan, trxType, docPendukung
		FROM penjualan WHERE id=?`

	res, err := uc.database.Query(query, id)
	if err != nil {
		return penjualan, err
	}
	defer res.Close()

	for res.Next() {
		if err := res.Scan(
			&penjualan.ID,
			&penjualan.CreatedBy,
			&penjualan.CreatedDate,
			&penjualan.NamaPelanggan,
			&penjualan.TipePelanggan,
			&penjualan.TrxType,
			&penjualan.DocPendukung,
		); err != nil {
			return penjualan, err
		}
	}

	return penjualan, nil
}

func (uc *uscase) GetAll() ([]Penjualan, error) {
	var penjualan []Penjualan

	query := `SELECT id, createdBy, createdDate, namaPelanggan, tipePelanggan, trxType, docPendukung
		FROM penjualan`

	res, err := uc.database.Query(query)
	if err != nil {
		return penjualan, err
	}
	defer res.Close()

	for res.Next() {
		var item Penjualan
		if err := res.Scan(
			&item.ID,
			&item.CreatedBy,
			&item.CreatedDate,
			&item.NamaPelanggan,
			&item.TipePelanggan,
			&item.TrxType,
			&item.DocPendukung,
		); err != nil {
			return penjualan, err
		}
		penjualan = append(penjualan, item)
	}

	return penjualan, nil
}

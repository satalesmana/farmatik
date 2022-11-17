package ProductModel

import (
	"database/sql"
	db "farmatik/app/database"
	hargajualModel "farmatik/app/database/model/hargajual"
)

type Handler interface {
	Insert(data *Product) (int64, error)
	GetByid(id string) (Product, error)
	GetAll() ([]Product, error)
	Delete(id string) (string, error)
	Update(data *Product) (string, error)
}

type Product struct {
	ID          int64                       `json:"id"`
	NamaProduct string                      `json:"nama_product,omitempty"`
	HargaBeli   int64                       `json:"harga_beli,omitempty"`
	Satuan      string                      `json:"satuan,omitempty"`
	HargaJual   []*hargajualModel.HargaJual `json:"harga_jual,omitempty"`
}

type uscase struct {
	database       *sql.DB
	hargajualModel hargajualModel.Handler
}

func NewProductModelHandler() Handler {
	return &uscase{
		database:       db.GetCoon(),
		hargajualModel: hargajualModel.NewHargaJualHandler(),
	}
}

// Insert implements Handler
func (uc *uscase) Insert(data *Product) (int64, error) {
	query := `INSERT INTO product(
		namaProduct, hargaBeli, satuan  
	) VALUES(?, ?, ?)`

	res, err := uc.database.Exec(query,
		&data.NamaProduct,
		&data.HargaBeli,
		&data.Satuan,
	)

	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	//add value to harga jual
	for i := 0; i < len(data.HargaJual); i++ {
		var (
			hargaJual hargajualModel.HargaJual
		)

		hargaJual.IdProduct = lastID
		hargaJual.Harga = data.HargaJual[i].Harga
		hargaJual.Kategori = data.HargaJual[i].Kategori

		_, err := uc.hargajualModel.Insert(&hargaJual)

		if err != nil {
			return 0, err
		}
	}

	return lastID, nil
}

func (uc *uscase) GetByid(id string) (Product, error) {
	var output Product
	var hargaProduct []*hargajualModel.HargaJual

	query := `SELECT id, namaProduct, hargaBeli, satuan FROM product WHERE id=?`
	res, err := uc.database.Query(query, id)
	if err != nil {
		return output, err
	}
	defer res.Close()

	for res.Next() {
		if err := res.Scan(
			&output.ID,
			&output.NamaProduct,
			&output.HargaBeli,
			&output.Satuan,
		); err != nil {
			return output, err
		}
	}

	//selecting data harga jual
	queryHarga := `SELECT id,kategori, harga FROM product_hargajual where idProduct=?`
	resHarga, err := uc.database.Query(queryHarga, output.ID)
	if err != nil {
		return output, err
	}
	defer resHarga.Close()

	for resHarga.Next() {
		var hargaJualData hargajualModel.HargaJual
		if err := resHarga.Scan(
			&hargaJualData.ID,
			&hargaJualData.Kategori,
			&hargaJualData.Harga,
		); err != nil {
			return output, err
		}
		hargaProduct = append(hargaProduct, &hargaJualData)
	}

	output.HargaJual = hargaProduct

	return output, nil
}

func (uc *uscase) GetAll() ([]Product, error) {
	query := `SELECT id, namaProduct, hargaBeli, satuan FROM product`
	res, err := uc.database.Query(query)
	var output []Product
	if err != nil {
		return output, err
	}
	defer res.Close()

	for res.Next() {
		var data Product
		if err := res.Scan(
			&data.ID,
			&data.NamaProduct,
			&data.HargaBeli,
			&data.Satuan,
		); err != nil {
			return output, err
		}
		output = append(output, data)
	}

	return output, nil
}

func (uc *uscase) Delete(id string) (string, error) {
	query := `DELETE FROM product WHERE id=?`
	_, err := uc.database.Exec(query, id)
	if err != nil {
		return "", err
	}

	//delete data harga jual
	queryHargaJual := `DELETE FROM product_hargajual WHERE idProduct=?`
	_, errHargaJual := uc.database.Exec(queryHargaJual, id)
	if errHargaJual != nil {
		return "", err
	}

	return "Data berhasil dihapus", nil
}

func (uc *uscase) Update(data *Product) (string, error) {
	query := `UPDATE product SET namaProduct=?, hargaBeli=?, satuan=? WHERE  id=?`
	_, err := uc.database.Exec(query,
		&data.NamaProduct,
		&data.HargaBeli,
		&data.Satuan,
		&data.ID)
	if err != nil {
		return "", err
	}

	return "Data berhasil perbaharui", nil
}

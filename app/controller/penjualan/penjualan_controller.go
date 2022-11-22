package PenjualanController

import (
	PenjualanModel "farmatik/app/database/model/penjualan"
	PenjualanDetailModel "farmatik/app/database/model/penjualan_detail"
	ProductmutationModel "farmatik/app/database/model/product_mutation"
	"farmatik/system/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddNew(c *gin.Context)
	GetAll(c *gin.Context)
	FindById(c *gin.Context)
}

type uscase struct {
	PenjualanModel       PenjualanModel.Handler
	PenjualanDetailModel PenjualanDetailModel.Handler
	ProductmutationModel ProductmutationModel.Handler
}

func NewPenjualanControllerHandler() Handler {
	return &uscase{
		PenjualanModel:       PenjualanModel.NewPenjualanHandler(),
		PenjualanDetailModel: PenjualanDetailModel.NewPenjualanDetailHandler(),
		ProductmutationModel: ProductmutationModel.NewProductMutationHandler(),
	}
}

func (m *uscase) AddNew(c *gin.Context) {
	var (
		data PenjualanModel.Penjualan
	)
	curDate := time.Now()
	myDate := curDate.Format("2006-01-02") //fmt.Println(t.Format("2006-01-02 15:04:05"))
	curId := c.GetString("user_id")

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}

	lastID, err := m.PenjualanModel.Insert(&data)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	//add to detail data
	dataDetail := data.PenjualanDetail
	for _, v := range dataDetail {
		dataDetail := &PenjualanDetailModel.PenjualanDetail{
			ID:          v.ID,
			PenjualanId: lastID,
			ProductId:   v.ProductId,
			Harga:       v.Harga,
			Qty:         v.Qty,
		}

		_, errDetail := m.PenjualanDetailModel.Insert(dataDetail)
		if errDetail != nil {
			c.JSON(response.Format(http.StatusInternalServerError, errDetail))
			return
		}

		//add data mutation
		mutationData := &ProductmutationModel.ProductMutation{
			ID:          0,
			ProductId:   v.ProductId,
			TrxCode:     lastID,
			CreatedBy:   curId,
			CreatedDate: myDate,
			Type:        "O",
			Value:       v.Qty,
		}
		_, errMutation := m.ProductmutationModel.Insert(mutationData)
		if errMutation != nil {
			c.JSON(response.Format(http.StatusInternalServerError, errMutation))
			return
		}
	}

	data.ID = lastID

	c.JSON(response.Format(http.StatusOK, nil, data))
}

func (m *uscase) GetAll(c *gin.Context) {
	data, err := m.PenjualanModel.GetAll()
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

func (m *uscase) FindById(c *gin.Context) {
	id := c.Param("id")

	data, err := m.PenjualanModel.GetById(id)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	//Get data detail penjualan
	resDetail, errDetail := m.PenjualanDetailModel.GetByTrx(id)
	if errDetail != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	//looping data array
	for _, v := range resDetail {
		dataDetail := &PenjualanDetailModel.PenjualanDetail{
			ID:          v.ID,
			PenjualanId: v.PenjualanId,
			ProductId:   v.ProductId,
			Harga:       v.Harga,
			Qty:         v.Qty,
		}
		data.PenjualanDetail = append(data.PenjualanDetail, dataDetail)
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

package ProductController

import (
	ProductModel "farmatik/app/database/model/product"
	"farmatik/system/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAll(c *gin.Context)
	AddNew(c *gin.Context)
	FindById(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type uscase struct {
	ProductModel ProductModel.Handler
}

func NewProductControllerHandler() Handler {
	return &uscase{
		ProductModel: ProductModel.NewProductModelHandler(),
	}
}

func (m *uscase) AddNew(c *gin.Context) {
	var (
		data ProductModel.Product
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}

	lastID, err := m.ProductModel.Insert(&data)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}
	data.ID = lastID.ID

	c.JSON(response.Format(http.StatusOK, nil, data))

}

func (m *uscase) GetAll(c *gin.Context) {
	data, err := m.ProductModel.GetAll()
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

func (m *uscase) FindById(c *gin.Context) {
	id := c.Param("id")

	data, err := m.ProductModel.GetByid(id)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

func (m *uscase) Edit(c *gin.Context) {
	var (
		data ProductModel.Product
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}
	strID := c.Param("id")
	//intID, _ := strconv.ParseInt(strID, 10, 64)
	data.ID = strID

	res, err := m.ProductModel.Update(&data)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, res))
}

func (m *uscase) Delete(c *gin.Context) {
	id := c.Param("id")

	data, err := m.ProductModel.Delete(id)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

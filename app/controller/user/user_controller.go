package usersController

import (
	"errors"
	userModel "farmatik/app/database/model/user"
	"farmatik/system/hash"
	"farmatik/system/response"
	"farmatik/system/validation"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAll(c *gin.Context)
	AddNew(c *gin.Context)
	FindBy(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}

type uscase struct {
	userModel userModel.Handler
}

func NewUserHandler() Handler {
	return &uscase{
		userModel: userModel.NewUserHandler(),
	}
}

/**
 * Function handle Get All data.
 * @GetAll
 */

func (m *uscase) GetAll(c *gin.Context) {
	data, err := m.userModel.GetAll()
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

/**
 * Function handle Add
 * @AddNew
 */
func (m *uscase) AddNew(c *gin.Context) {
	var (
		data userModel.User
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}

	isValidEmail := validation.IsEmailValid(data.Email)
	if !isValidEmail {
		err := errors.New("invalid email format")
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	data.Password = hash.HashPassword(data.Password)

	lastID, err := m.userModel.Insert(&data)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}
	data.ID = lastID

	c.JSON(response.Format(http.StatusOK, nil, data))

}

/**
 * Function find data by id
 * @FindBy
 */
func (m *uscase) FindBy(c *gin.Context) {
	id := c.Param("id")

	data, err := m.userModel.GetByid(id)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))

}

/**
 * Function find edit data
 * @Edit
 */
func (m *uscase) Edit(c *gin.Context) {
	var (
		data userModel.User
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}
	strID := c.Param("id")
	intID, _ := strconv.ParseInt(strID, 10, 64)
	data.ID = intID

	isValidEmail := validation.IsEmailValid(data.Email)
	if !isValidEmail {
		err := errors.New("invalid email format")
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	res, err := m.userModel.Update(&data)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, res))
}

/**
 * Function Delete data
 * @Delete
 */
func (m *uscase) Delete(c *gin.Context) {
	id := c.Param("id")

	data, err := m.userModel.Delete(id)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))

}

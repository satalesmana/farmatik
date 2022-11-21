package appconfigController

import (
	appconfig "farmatik/app/database/model/app_config"
	"farmatik/system/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	FindById(c *gin.Context)
}

type uscase struct {
	AppConfigModel appconfig.Handler
}

func NewAppConfigControllerHandler() Handler {
	return &uscase{
		AppConfigModel: appconfig.NewConfigHandler(),
	}
}

func (m *uscase) FindById(c *gin.Context) {
	var (
		request appconfig.AppConfigRequest
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}

	data, err := m.AppConfigModel.GetById(request)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(response.Format(http.StatusOK, nil, data))
}

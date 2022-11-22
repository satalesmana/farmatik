package midleware

import (
	"errors"
	userLoginModel "farmatik/app/database/model/user_login"
	"farmatik/system/pkg/jwt"
	resp "farmatik/system/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	UseMiddleware() gin.HandlerFunc
}

type uscase struct {
	userLoginModel userLoginModel.Handler
}

func NewAuthHandler() Handler {
	return &uscase{
		userLoginModel: userLoginModel.NewUserLoginHandler(),
	}
}

// Middleware ...
func (m *uscase) UseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authorization = strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)
		)

		if authorization == "" {
			c.JSON(resp.Format(http.StatusBadRequest, errors.New("please provide authorization")))
			c.Abort()
			return
		}

		if err := jwt.TokenValid(authorization); err != nil {
			c.JSON(resp.Format(http.StatusUnauthorized, errors.New("invalid token")))
			c.Abort()
			return
		}

		metadata, err := jwt.ExtractTokenMetadata(authorization)
		if err != nil {
			c.JSON(resp.Format(http.StatusInternalServerError, errors.New("unable to extract token metadata")))
			c.Abort()
			return
		}

		res, errLogin := m.userLoginModel.GetByid(metadata.ID)
		if errLogin != nil {
			c.JSON(resp.Format(http.StatusInternalServerError, errors.New("session not found")))
			c.Abort()
			return
		}

		if res.Status != "A" {
			c.JSON(resp.Format(http.StatusInternalServerError, errors.New("session expired")))
			c.Abort()
			return
		}

		// token valid and forward to original request
		c.Set("id", metadata.ID)
		c.Set("user_id", res.IDUser)
		c.Set("email", metadata.Email)
		c.Set("name", metadata.Name)

		c.Next()
		// original request goes here
	}
}

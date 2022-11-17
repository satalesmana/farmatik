package authController

import (
	"errors"
	userModel "farmatik/app/database/model/user"
	userLoginModel "farmatik/app/database/model/user_login"
	"farmatik/system/hash"
	"farmatik/system/pkg/jwt"
	"farmatik/system/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type uscase struct {
	userModel      userModel.Handler
	userLoginModel userLoginModel.Handler
}

type Response struct {
	Token string `json:"token,omitempty"`
}

type Request struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Handler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

func NewAuthHandler() Handler {
	return &uscase{
		userModel:      userModel.NewUserHandler(),
		userLoginModel: userLoginModel.NewUserLoginHandler(),
	}
}

func (m *uscase) Login(c *gin.Context) {
	var (
		request Request
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(response.Format(http.StatusBadRequest, err))
		return
	}

	data, err := m.userModel.GetByEmail(request.Email)
	if err != nil {
		err := errors.New("username tidak terdaftar")
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	isMatchPwd := hash.DoPasswordsMatch(data.Password, request.Password)
	if !isMatchPwd {
		err := errors.New("invalid username and password")
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	newUUID := uuid.New()
	//genertare token
	token, err := jwt.CreateToken(newUUID.String(), data.Name, request.Email)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	var (
		dataLogin userLoginModel.UserLogin
	)
	dataLogin.ID = newUUID.String()
	strIDUser := strconv.Itoa(int(data.ID))
	dataLogin.IDUser = strIDUser
	dataLogin.Status = "A"
	dataLogin.Token = token

	_, errUserMod := m.userLoginModel.Insert(&dataLogin)
	if errUserMod != nil {
		err := errUserMod
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	output := &Response{Token: "Bearer " + token}
	c.JSON(response.Format(http.StatusOK, nil, output))
}

func (m *uscase) Logout(c *gin.Context) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)

	metaToken, err := jwt.ExtractTokenMetadata(token)
	if err != nil {
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	_, errUpdate := m.userLoginModel.Update(metaToken.ID, "N")
	if errUpdate != nil {
		err := errUpdate
		c.JSON(response.Format(http.StatusInternalServerError, err))
		return
	}

	output := &Response{}
	c.JSON(response.Format(http.StatusOK, nil, output))
}

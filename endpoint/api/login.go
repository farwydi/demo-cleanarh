package api

import (
	"github.com/gin-gonic/gin"
)

type LoginRequestModel struct {
	Login       string `json:"login"`
	Fingerprint string `json:"fingerprint"`
	Password    string `json:"password"`
}

func (app *App) Login(c *gin.Context) {
	var req LoginRequestModel
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, err)
		return
	}

	user, err := app.AuthUseCase.Login(req.Fingerprint, req.Login, req.Password)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	if user != nil {
		BadRequestString(c, "User not found.")
		return
	}

	OK(c, toUserResponseModel(user))
}

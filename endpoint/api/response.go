package api

import (
	"github.com/farwydi/demo-cleanarh/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponseModel struct {
	Code    int
	Message string
}

type UserResponseModel struct {
	ID   string
	Name string
}

func toUserResponseModel(user *domain.User) UserResponseModel {
	return UserResponseModel{
		ID:   user.GetIDString(),
		Name: user.Name,
	}
}

func BadRequestString(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, ErrorResponseModel{
		Code:    http.StatusBadRequest,
		Message: err,
	})
}

func BadRequest(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponseModel{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
}

func InternalServerError(c *gin.Context, err error) {
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, ErrorResponseModel{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
	}
}

func OK(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

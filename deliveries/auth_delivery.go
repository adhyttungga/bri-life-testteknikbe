package deliveries

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adhyttungga/bri-life-testteknikbe/models/dto"
	"github.com/adhyttungga/bri-life-testteknikbe/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthDelivery interface {
	Login(c *gin.Context)
}

type AuthDeliveryImpl struct {
	AU usecases.AuthUsecase
}

func NewAuthDelivery(au usecases.AuthUsecase) AuthDelivery {
	return &AuthDeliveryImpl{
		AU: au,
	}
}

func (ad *AuthDeliveryImpl) Login(c *gin.Context) {
	var req dto.RequestAuth
	resError := dto.ResponseError{
		ResponseCode: "01",
		ResponseDesc: http.StatusText(http.StatusInternalServerError),
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resError.ResponseDesc = fmt.Sprintf("ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, resError)
		return
	}

	res, err := ad.AU.Login(c.Request.Context(), req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) || errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, usecases.ErrInvalidCredential) {
			resError.ResponseDesc = http.StatusText(http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, resError)
		} else {
			resError.ResponseDesc = http.StatusText(http.StatusInternalServerError)
			c.JSON(http.StatusInternalServerError, resError)
		}

		return
	}

	c.JSON(http.StatusOK, res)
}

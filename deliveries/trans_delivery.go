package deliveries

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adhyttungga/bri-life-testteknikbe/models/dto"
	"github.com/adhyttungga/bri-life-testteknikbe/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransDelivery interface {
	CreateTrans(c *gin.Context)
	UpdateTrans(c *gin.Context)
	DeleteTrans(c *gin.Context)
}

type TransDeliveryImpl struct {
	TU usecases.TransUsecase
}

func NewTransDelivery(tu usecases.TransUsecase) TransDelivery {
	return &TransDeliveryImpl{
		TU: tu,
	}
}

func (td *TransDeliveryImpl) CreateTrans(c *gin.Context) {
	var req dto.RequestCreateTrans
	resErr := dto.ResponseError{
		ResponseCode: "01",
		ResponseDesc: http.StatusText(http.StatusInternalServerError),
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resErr.ResponseDesc = fmt.Sprintf("ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res, err := td.TU.CreateTrans(c.Request.Context(), req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) || errors.Is(err, usecases.ErrValidateAge) || errors.Is(err, usecases.ErrValidatePremium) || errors.Is(err, usecases.ErrValidateAgentId) {
			resErr.ResponseDesc = http.StatusText(http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, resErr)
		} else {
			resErr.ResponseDesc = http.StatusText(http.StatusInternalServerError)
			c.JSON(http.StatusInternalServerError, resErr)
		}

		return
	}

	c.JSON(http.StatusOK, res)
}

func (td *TransDeliveryImpl) UpdateTrans(c *gin.Context) {
	var req dto.RequestTrans
	resErr := dto.ResponseError{
		ResponseCode: "01",
		ResponseDesc: http.StatusText(http.StatusInternalServerError),
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resErr.ResponseDesc = fmt.Sprintf("ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res, err := td.TU.UpdateTrans(c.Request.Context(), req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) || errors.Is(err, usecases.ErrValidateAge) || errors.Is(err, usecases.ErrValidatePremium) || errors.Is(err, usecases.ErrValidateAgentId) {
			resErr.ResponseDesc = http.StatusText(http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, resErr)
		} else {
			resErr.ResponseDesc = http.StatusText(http.StatusInternalServerError)
			c.JSON(http.StatusInternalServerError, resErr)
		}

		return
	}

	c.JSON(http.StatusOK, res)
}

func (td *TransDeliveryImpl) DeleteTrans(c *gin.Context) {
	var req dto.RequestTrans
	resErr := dto.ResponseError{
		ResponseCode: "01",
		ResponseDesc: http.StatusText(http.StatusInternalServerError),
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resErr.ResponseDesc = fmt.Sprintf("ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res, err := td.TU.DeleteTrans(c.Request.Context(), req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) || errors.Is(err, usecases.ErrValidateAge) || errors.Is(err, usecases.ErrValidatePremium) || errors.Is(err, usecases.ErrValidateAgentId) {
			resErr.ResponseDesc = http.StatusText(http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, resErr)
		} else {
			resErr.ResponseDesc = http.StatusText(http.StatusInternalServerError)
			c.JSON(http.StatusInternalServerError, resErr)
		}

		return
	}

	c.JSON(http.StatusOK, res)
}

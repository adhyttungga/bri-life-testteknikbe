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

type AgentDelivery interface {
	CreateAgent(c *gin.Context)
	UpdateAgent(c *gin.Context)
	DeleteAgent(c *gin.Context)
}

type AgentDeliveryImpl struct {
	AU usecases.AgentUsecase
}

func NewAgentDelivery(au usecases.AgentUsecase) AgentDelivery {
	return &AgentDeliveryImpl{AU: au}
}

func (ad *AgentDeliveryImpl) CreateAgent(c *gin.Context) {
	var req dto.RequestAgent
	resErr := dto.ResponseError{
		ResponseCode: "01",
		ResponseDesc: http.StatusText(http.StatusInternalServerError),
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resErr.ResponseDesc = fmt.Sprintf("ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res, err := ad.AU.CreateAgent(c.Request.Context(), req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
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

func (ad *AgentDeliveryImpl) UpdateAgent(c *gin.Context) {
	var req dto.RequestAgent
	resErr := dto.ResponseError{
		ResponseCode: "01",
		ResponseDesc: http.StatusText(http.StatusInternalServerError),
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resErr.ResponseDesc = fmt.Sprintf("ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res, err := ad.AU.UpdateAgent(c.Request.Context(), req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) || errors.Is(err, gorm.ErrRecordNotFound) {
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

func (ad *AgentDeliveryImpl) DeleteAgent(c *gin.Context) {
	var req dto.RequestDeleteAgent
	resErr := dto.ResponseError{
		ResponseCode: "01",
		ResponseDesc: http.StatusText(http.StatusInternalServerError),
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resErr.ResponseDesc = fmt.Sprintf("ShouldBindJSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, resErr)
		return
	}

	res, err := ad.AU.DeleteAgent(c.Request.Context(), req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) || errors.Is(err, gorm.ErrRecordNotFound) {
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

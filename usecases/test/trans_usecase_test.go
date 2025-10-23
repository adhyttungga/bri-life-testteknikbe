package test

import (
	"context"
	"testing"

	"github.com/adhyttungga/bri-life-testteknikbe/models/dto"
	"github.com/adhyttungga/bri-life-testteknikbe/models/entity"
	"github.com/adhyttungga/bri-life-testteknikbe/usecases"
	repoMocks "github.com/adhyttungga/bri-life-testteknikbe/usecases/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrans_Success(t *testing.T) {
	mockTransRepo := new(repoMocks.TransRepositoryMock)
	mockProdRepo := new(repoMocks.ProductRepositoryMock)

	uc := usecases.NewTransUsecase(mockTransRepo, mockProdRepo)

	trans := entity.Transaction{}
	mockTransRepo.On("CreateTrans", context.Background(), &trans).Return(nil)
	mockProdRepo.On("GetProdctsWithParamById", context.Background(), "").Return(&map[string]any{}, nil)

	res, err := uc.CreateTrans(context.Background(), dto.RequestCreateTrans{})

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, res.ResponseCode, "00")
}

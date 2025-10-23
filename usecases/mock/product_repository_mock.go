package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (tr *ProductRepositoryMock) GetProdctsWithParamById(ctx context.Context, productId string) (*map[string]any, error) {
	args := tr.Called(ctx, productId)
	return args.Get(0).(*map[string]any), args.Error(1)
}

package mocks

import (
	"context"

	"github.com/adhyttungga/bri-life-testteknikbe/models/entity"
	"github.com/stretchr/testify/mock"
)

type TransRepositoryMock struct {
	mock.Mock
}

func (tr *TransRepositoryMock) CreateTrans(ctx context.Context, trans *entity.Transaction) error {
	args := tr.Called(ctx, trans)
	return args.Error(0)
}

func (tr *TransRepositoryMock) UpdateTrans(ctx context.Context, trans *entity.Transaction) error {
	args := tr.Called(ctx, trans)
	return args.Error(0)
}

func (tr *TransRepositoryMock) DeleteTrans(ctx context.Context, transId int64) error {
	args := tr.Called(ctx, transId)
	return args.Error(0)
}

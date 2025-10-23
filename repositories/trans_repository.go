package repositories

import (
	"context"
	"fmt"

	"github.com/adhyttungga/bri-life-testteknikbe/models/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransRepository interface {
	CreateTrans(ctx context.Context, trans *entity.Transaction) error
	UpdateTrans(ctx context.Context, trans *entity.Transaction) error
	DeleteTrans(ctx context.Context, transId int64) error
}

type TransRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransRepository(db *gorm.DB) TransRepository {
	return &TransRepositoryImpl{DB: db}
}

func (tr *TransRepositoryImpl) CreateTrans(ctx context.Context, trans *entity.Transaction) error {
	if err := tr.DB.
		WithContext(ctx).
		Create(trans).
		Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (tr *TransRepositoryImpl) UpdateTrans(ctx context.Context, trans *entity.Transaction) error {
	if err := tr.DB.
		WithContext(ctx).
		Model(&entity.Transaction{}).
		Where("trans_id = ?", trans.TransId).
		Clauses(clause.Returning{}).
		Updates(trans).
		Error; err != nil {
		return fmt.Errorf("DB.Updates: %w", err)
	}

	return nil
}

func (tr *TransRepositoryImpl) DeleteTrans(ctx context.Context, transId int64) error {
	if err := tr.DB.
		WithContext(ctx).
		Where("trans_id = ?", transId).
		Delete(&entity.Transaction{}).
		Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}

	return nil
}

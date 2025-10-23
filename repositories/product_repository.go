package repositories

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProdctsWithParamById(ctx context.Context, productId string) (*map[string]any, error)
}

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (pr *ProductRepositoryImpl) GetProdctsWithParamById(ctx context.Context, productId string) (*map[string]any, error) {
	var product map[string]any
	query := fmt.Sprintf(`
SELECT 
	product_name,
	premium,
	min_pp.parameter_value AS min_value,
    max_pp.parameter_value AS max_value
FROM product AS p
INNER JOIN (SELECT * FROM product_parameter WHERE product_id = '%s' AND parameter_name = 'MINUSIAMASUK') AS min_pp ON min_pp.product_id = p.product_id
INNER JOIN (SELECT * FROM product_parameter WHERE product_id = '%s' AND parameter_name = 'MAXUSIAMASUK') AS max_pp ON max_pp.product_id = p.product_id
WHERE p.product_id = '%s';	
	`, productId, productId, productId)

	if err := pr.DB.Raw(query).Scan(&product).Error; err != nil {
		return nil, fmt.Errorf("DB.Raw: %w", err)
	}

	return &product, nil
}

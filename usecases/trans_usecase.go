package usecases

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/adhyttungga/bri-life-testteknikbe/models/dto"
	"github.com/adhyttungga/bri-life-testteknikbe/models/entity"
	"github.com/adhyttungga/bri-life-testteknikbe/repositories"
	"github.com/adhyttungga/bri-life-testteknikbe/utils"
	"github.com/shopspring/decimal"
)

var (
	errMapProduct      = errors.New("error mapping product")
	ErrValidateAge     = errors.New("error validating age")
	ErrValidatePremium = errors.New("error validating premium")
	ErrValidateAgentId = errors.New("error validating agent id")
)

type TransUsecase interface {
	CreateTrans(ctx context.Context, req dto.RequestCreateTrans) (*dto.ResponseSuccess, error)
	UpdateTrans(ctx context.Context, req dto.RequestTrans) (*dto.ResponseSuccess, error)
	DeleteTrans(ctx context.Context, req dto.RequestTrans) (*dto.ResponseSuccess, error)
}

type TransUsecaseImpl struct {
	TR repositories.TransRepository
	PR repositories.ProductRepository
}

func NewTransUsecase(
	tr repositories.TransRepository,
	pr repositories.ProductRepository,
) TransUsecase {
	return &TransUsecaseImpl{
		TR: tr,
		PR: pr,
	}
}

func (tu *TransUsecaseImpl) CreateTrans(ctx context.Context, req dto.RequestCreateTrans) (*dto.ResponseSuccess, error) {
	// Validate request
	if err := utils.ValidateStruct(ctx, req); err != nil {
		return nil, fmt.Errorf("ValidateStruct: %w", err)
	}

	// Get product by product id
	product, err := tu.PR.GetProdctsWithParamById(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	premium, ok := (*product)["premium"].(string)
	if !ok {
		return nil, fmt.Errorf("error mapping premium: %w", errMapProduct)
	}

	premiumDec, err := decimal.NewFromString(premium)
	if err != nil {
		return nil, fmt.Errorf("decimal.NewFromString: %w", err)
	}

	minAge, ok := (*product)["min_value"].(string)
	if !ok {
		return nil, fmt.Errorf("error mapping minAge: %w", errMapProduct)
	}

	minAgeInt, err := strconv.Atoi(minAge)
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi: %w", err)
	}

	maxAge, ok := (*product)["max_value"].(string)
	if !ok {
		return nil, fmt.Errorf("error mapping maxAge: %w", errMapProduct)
	}

	maxAgeInt, err := strconv.Atoi(maxAge)
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi: %w", err)
	}

	// Validate min and max age
	if req.Usia < minAgeInt || req.Usia > maxAgeInt {
		return nil, ErrValidateAge
	}

	// Validate premium
	reqPremiumDec, err := decimal.NewFromString(req.Premium)
	if err != nil {
		return nil, fmt.Errorf("decimal.NewFromString: %w", err)
	}

	if !reqPremiumDec.Mod(premiumDec).IsZero() {
		return nil, ErrValidatePremium
	}

	// Create transaction
	trans := entity.Transaction{
		AgentId:   req.AgentId,
		ProductId: req.ProductId,
		Nama:      req.Nama,
		Usia:      req.Usia,
		Premium:   premiumDec,
	}
	if err := tu.TR.CreateTrans(ctx, &trans); err != nil {
		return nil, err
	}

	return &dto.ResponseSuccess{
		ResponseCode: "00",
		ResponseDesc: "Transaction created successfully",
		Data: dto.TransactionData{
			TransId: fmt.Sprintf("%d", trans.TransId),
		},
	}, nil
}

func (tu *TransUsecaseImpl) UpdateTrans(ctx context.Context, req dto.RequestTrans) (*dto.ResponseSuccess, error) {
	// Validate request
	if err := utils.ValidateStruct(ctx, req); err != nil {
		return nil, fmt.Errorf("ValidateStruct: %w", err)
	}

	// Get product by product id
	product, err := tu.PR.GetProdctsWithParamById(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	premium, ok := (*product)["premium"].(string)
	if !ok {
		return nil, fmt.Errorf("error mapping premium: %w", errMapProduct)
	}

	premiumDec, err := decimal.NewFromString(premium)
	if err != nil {
		return nil, fmt.Errorf("decimal.NewFromString: %w", err)
	}

	minAge, ok := (*product)["min_value"].(string)
	if !ok {
		return nil, fmt.Errorf("error mapping minAge: %w", errMapProduct)
	}

	minAgeInt, err := strconv.Atoi(minAge)
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi: %w", err)
	}

	maxAge, ok := (*product)["max_value"].(string)
	if !ok {
		return nil, fmt.Errorf("error mapping maxAge: %w", errMapProduct)
	}

	maxAgeInt, err := strconv.Atoi(maxAge)
	if err != nil {
		return nil, fmt.Errorf("strconv.Atoi: %w", err)
	}

	// Validate min and max age
	if req.Usia < minAgeInt || req.Usia > maxAgeInt {
		return nil, ErrValidateAge
	}

	// Validate premium
	reqPremiumDec, err := decimal.NewFromString(req.Premium)
	if err != nil {
		return nil, fmt.Errorf("decimal.NewFromString: %w", err)
	}

	if !reqPremiumDec.Mod(premiumDec).IsZero() {
		return nil, ErrValidatePremium
	}

	// Update transaction
	trans := entity.Transaction{
		TransId:   req.TransId,
		AgentId:   req.AgentId,
		ProductId: req.ProductId,
		Nama:      req.Nama,
		Usia:      req.Usia,
		Premium:   premiumDec,
	}
	if err := tu.TR.UpdateTrans(ctx, &trans); err != nil {
		return nil, err
	}

	return &dto.ResponseSuccess{
		ResponseCode: "00",
		ResponseDesc: "Transaction updated successfully",
		Data: dto.TransactionData{
			TransId: fmt.Sprintf("%d", req.TransId),
		},
	}, nil
}

func (tu *TransUsecaseImpl) DeleteTrans(ctx context.Context, req dto.RequestTrans) (*dto.ResponseSuccess, error) {
	// Validate Request
	if err := utils.ValidateStruct(ctx, req); err != nil {
		return nil, fmt.Errorf("ValidateStruct: %w", err)
	}

	// Get agent id from access token
	agentId := ctx.Value("agent_id").(string)

	// Validate agentId
	if req.AgentId != agentId {
		return nil, ErrValidateAgentId
	}

	// Delete transaction
	if err := tu.TR.DeleteTrans(ctx, req.TransId); err != nil {
		return nil, err
	}

	return &dto.ResponseSuccess{
		ResponseCode: "00",
		ResponseDesc: "Delete transaction successfully",
		Data: dto.TransactionData{
			TransId: fmt.Sprintf("%d", req.TransId),
		},
	}, nil
}

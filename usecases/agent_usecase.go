package usecases

import (
	"context"
	"fmt"

	"github.com/adhyttungga/bri-life-testteknikbe/models/dto"
	"github.com/adhyttungga/bri-life-testteknikbe/models/entity"
	"github.com/adhyttungga/bri-life-testteknikbe/repositories"
	"github.com/adhyttungga/bri-life-testteknikbe/utils"
)

type AgentUsecase interface {
	CreateAgent(ctx context.Context, req dto.RequestAgent) (*dto.ResponseSuccess, error)
	UpdateAgent(ctx context.Context, req dto.RequestAgent) (*dto.ResponseSuccess, error)
	DeleteAgent(ctx context.Context, req dto.RequestDeleteAgent) (*dto.ResponseSuccess, error)
}

type AgentUsecaseImpl struct {
	AR repositories.AgentRepository
}

func NewAgentUsecase(ar repositories.AgentRepository) AgentUsecase {
	return &AgentUsecaseImpl{
		AR: ar,
	}
}

func (as *AgentUsecaseImpl) CreateAgent(ctx context.Context, req dto.RequestAgent) (*dto.ResponseSuccess, error) {
	// Validate request
	if err := utils.ValidateStruct(ctx, req); err != nil {
		return nil, fmt.Errorf("ValidateStruct: %w", err)
	}

	// Create Agent
	if err := as.AR.CreateAgent(ctx, &entity.Agent{
		AgentId:   req.AgentId,
		AgentName: req.AgentName,
		Password:  req.Password,
		Active:    req.Active,
	}); err != nil {
		return nil, err
	}

	return &dto.ResponseSuccess{
		ResponseCode: "00",
		ResponseDesc: "New agent created successfully",
		Data: dto.AgentData{
			AgentId:   req.AgentId,
			AgentName: req.AgentName,
		},
	}, nil
}

func (as *AgentUsecaseImpl) UpdateAgent(ctx context.Context, req dto.RequestAgent) (*dto.ResponseSuccess, error) {
	// Validate request
	if err := utils.ValidateStruct(ctx, req); err != nil {
		return nil, fmt.Errorf("ValidateStruct: %w", err)
	}

	// Update Agent
	if err := as.AR.UpdateAgent(ctx, &entity.Agent{
		AgentId:   req.AgentId,
		AgentName: req.AgentName,
		Password:  req.Password,
		Active:    req.Active,
	}); err != nil {
		return nil, err
	}

	return &dto.ResponseSuccess{
		ResponseCode: "00",
		ResponseDesc: "Agent updated successfully",
		Data: dto.AgentData{
			AgentId:   req.AgentId,
			AgentName: req.AgentName,
		},
	}, nil
}

func (as *AgentUsecaseImpl) DeleteAgent(ctx context.Context, req dto.RequestDeleteAgent) (*dto.ResponseSuccess, error) {
	// Validate request
	if err := utils.ValidateStruct(ctx, req); err != nil {
		return nil, fmt.Errorf("ValidateStruct: %w", err)
	}

	// Delete Agent
	if err := as.AR.DeleteAgent(ctx, req.AgentId); err != nil {
		return nil, err
	}

	return &dto.ResponseSuccess{
		ResponseCode: "00",
		ResponseDesc: "Agent deleted successfully",
		Data: dto.AgentData{
			AgentId:   req.AgentId,
			AgentName: "",
		},
	}, nil
}

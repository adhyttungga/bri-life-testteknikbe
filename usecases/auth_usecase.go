package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/adhyttungga/bri-life-testteknikbe/models/dto"
	"github.com/adhyttungga/bri-life-testteknikbe/repositories"
	"github.com/adhyttungga/bri-life-testteknikbe/utils"
)

var ErrInvalidCredential = errors.New("authentication failed, invalid credential")

type AuthUsecase interface {
	Login(ctx context.Context, req dto.RequestAuth) (*dto.ResponseSuccess, error)
}

type AuthUsecaseImpl struct {
	AR repositories.AgentRepository
}

func NewAuthUsecase(ar repositories.AgentRepository) AuthUsecase {
	return &AuthUsecaseImpl{
		AR: ar,
	}
}

func (au *AuthUsecaseImpl) Login(ctx context.Context, req dto.RequestAuth) (*dto.ResponseSuccess, error) {
	// Validate request
	if err := utils.ValidateStruct(ctx, req); err != nil {
		return nil, fmt.Errorf("ValidateStruct: %w", err)
	}

	// Get Agent by id
	agent, err := au.AR.GetAgentById(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	// Validate password
	if agent.Password != req.Password {
		return nil, ErrInvalidCredential
	}

	// Generate token
	accessToken, err := utils.GenerateToken(req.AgentId)
	if err != nil {
		return nil, fmt.Errorf("GenerateToken: %w", err)
	}

	return &dto.ResponseSuccess{
		ResponseCode: "00",
		ResponseDesc: "Login successfully",
		Data: dto.AuthData{
			AgentId:     req.AgentId,
			AgentName:   agent.AgentName,
			AccessToken: accessToken,
		},
	}, nil
}

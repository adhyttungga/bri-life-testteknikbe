package repositories

import (
	"context"
	"fmt"

	"github.com/adhyttungga/bri-life-testteknikbe/models/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AgentRepository interface {
	GetAgentById(ctx context.Context, agentId string) (*entity.Agent, error)
	CreateAgent(ctx context.Context, agent *entity.Agent) error
	UpdateAgent(ctx context.Context, agent *entity.Agent) error
	DeleteAgent(ctx context.Context, agentId string) error
}

type AgentRepositoryImpl struct {
	DB *gorm.DB
}

func NewAgentRepository(db *gorm.DB) AgentRepository {
	return &AgentRepositoryImpl{
		DB: db,
	}
}

func (ar *AgentRepositoryImpl) GetAgentById(ctx context.Context, agentId string) (*entity.Agent, error) {
	var agent entity.Agent
	if err := ar.DB.
		WithContext(ctx).
		Where("agent_id = ?", agentId).
		Find(&agent).
		Error; err != nil {
		return nil, fmt.Errorf("DB.Find: %w", err)
	}

	return &agent, nil
}

func (ar *AgentRepositoryImpl) CreateAgent(ctx context.Context, agent *entity.Agent) error {
	if err := ar.DB.WithContext(ctx).Create(agent).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (ar *AgentRepositoryImpl) UpdateAgent(ctx context.Context, agent *entity.Agent) error {
	if err := ar.DB.
		WithContext(ctx).
		Model(&entity.Agent{}).
		Where("agent_id = ?", agent.AgentId).
		Clauses(clause.Returning{}).
		Updates(agent).
		Error; err != nil {
		return fmt.Errorf("DB.Updates: %w", err)
	}

	return nil
}

func (ar *AgentRepositoryImpl) DeleteAgent(ctx context.Context, agentId string) error {
	if err := ar.DB.
		WithContext(ctx).
		Where("agent_id = ?", agentId).
		Delete(&entity.Agent{}).
		Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}

	return nil
}

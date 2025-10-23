package entity

import "time"

type Agent struct {
	AgentId   string    `gorm:"type:varchar(8);not null;unique;<-:create" json:"agent_id"`
	AgentName string    `gorm:"type:varchar(100);not null" json:"agent_name"`
	Password  string    `gorm:"type:varchar(25);not null" json:"password"`
	Active    bool      `gorm:"type:tinyint(1);not null" json:"active"`
	CreateAt  time.Time `gorm:"type:datetime" json:"create_at"`
}

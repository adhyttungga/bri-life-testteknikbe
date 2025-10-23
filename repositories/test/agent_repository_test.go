package repotest

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adhyttungga/bri-life-testteknikbe/models/entity"
	"github.com/adhyttungga/bri-life-testteknikbe/repositories"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var agnt = entity.Agent{
	AgentId:   "BFA01",
	AgentName: "Agent Satu",
	Password:  "P@ssagent1",
	Active:    true,
}

const (
	getRegexQuery    = "^SELECT (.+)"
	createRegexQuery = "^INSERT (.+)"
	updateRegexQuery = "^UPDATE (.+)"
	deleteRegexQuery = "^DELETE (.+)"
)

type AgentSuite struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	mockDB *sql.DB
	db     *gorm.DB
	repo   repositories.AgentRepository
}

func (as *AgentSuite) SetupTest() {
	var err error

	as.mockDB, as.mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	as.db, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      as.mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening GORM connection", err)
	}
	as.repo = repositories.NewAgentRepository(as.db)
}

func (as *AgentSuite) TearDownTest() {
	as.mockDB.Close()
	as.Require().NoError(as.mock.ExpectationsWereMet())
}

func (as *AgentSuite) TestCreateAgent_Success() {
	as.mock.ExpectBegin()
	as.mock.ExpectExec(createRegexQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	as.mock.ExpectCommit()

	err := as.repo.CreateAgent(context.Background(), &agnt)

	as.Nil(err)
}

func (as *AgentSuite) TestUpdateAgent_Success() {
	as.mock.ExpectBegin()
	as.mock.ExpectExec(updateRegexQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	as.mock.ExpectCommit()

	err := as.repo.UpdateAgent(context.Background(), &agnt)

	as.Nil(err)
}

func TestAgentRepositorySuite(t *testing.T) {
	suite.Run(t, new(AgentSuite))
}

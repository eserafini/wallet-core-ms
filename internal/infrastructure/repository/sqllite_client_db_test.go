package repository

import (
	"database/sql"
	"testing"

	"github.com/eserafini/wallet-core-ms/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id TEXT PRIMARY KEY, name TEXT, email TEXT, created_at TEXT, updated_at TEXT)")
	s.clientDB = NewSqClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGetClient() {
	client, err := entity.NewClient("John Doe", "john.doe@example.com")
	s.Nil(err)

	err = s.clientDB.Save(client)
	s.Nil(err)

	clientFound, err := s.clientDB.FindByID(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientFound.ID)
	s.Equal(client.Name, clientFound.Name)
	s.Equal(client.Email, clientFound.Email)
}

func (s *ClientDBTestSuite) TestSaveClient() {
	client, err := entity.NewClient("John Doe", "john.doe@example.com")
	s.Nil(err)

	err = s.clientDB.Save(client)
	s.Nil(err)
}

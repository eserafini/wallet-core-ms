package repository

import (
	"database/sql"
	"testing"

	"github.com/eserafini/wallet-core-ms/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), created_at DATETIME, updated_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255) PRIMARY KEY, client_id VARCHAR(255), balance DECIMAL(10, 2), created_at DATETIME, updated_at DATETIME)")
	s.accountDB = NewAccountDB(db)
	s.client, err = entity.NewClient("John Doe", "john.doe@example.com")
	s.Nil(err)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account, err := entity.NewAccount(s.client)
	s.Nil(err)

	err = s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {

	s.db.Exec("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt, s.client.UpdatedAt)

	account, err := entity.NewAccount(s.client)
	s.Nil(err)

	err = s.accountDB.Save(account)
	s.Nil(err)

	accountDB, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(s.client.Name, accountDB.Client.Name)
	s.Equal(s.client.Email, accountDB.Client.Email)
}

package repository

import (
	"database/sql"
	"testing"

	"github.com/eserafini/wallet-core-ms/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255), email VARCHAR(255), created_at DATETIME, updated_at DATETIME)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255) PRIMARY KEY, client_id VARCHAR(255), balance DECIMAL(10, 2), created_at DATETIME, updated_at DATETIME)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255) PRIMARY KEY, account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount DECIMAL(10, 2), created_at DATETIME)")
	client, err := entity.NewClient("John Doe", "john.doe@example.com")
	s.Nil(err)
	client2, err := entity.NewClient("John Doe 2", "john.doe2@example.com")
	s.Nil(err)

	s.client = client
	s.client2 = client2

	accountFrom, err := entity.NewAccount(client)
	s.Nil(err)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo, err := entity.NewAccount(client2)
	s.Nil(err)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreateTransaction() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Save(transaction)
	s.Nil(err)
}

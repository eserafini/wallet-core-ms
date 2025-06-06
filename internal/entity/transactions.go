package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	if err := transaction.validate(); err != nil {
		return nil, err
	}

	if err := transaction.Commit(); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) Commit() error {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)

	return nil
}

func (t *Transaction) validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient balance")
	}

	return nil
}

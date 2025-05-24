package createtransaction

import (
	"github.com/eserafini/wallet-core-ms/internal/entity"
	"github.com/eserafini/wallet-core-ms/internal/gateway"
)

type CreateTransactionUseCase struct {
	AccountGateway     gateway.AccountGateway
	TransactionGateway gateway.TransactionGateway
}

func NewCreateTransactionUseCase(accountGateway gateway.AccountGateway, transactionGateway gateway.TransactionGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		AccountGateway:     accountGateway,
		TransactionGateway: transactionGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountFromID)
	if err != nil {
		return CreateTransactionOutputDTO{}, err
	}

	accountTo, err := uc.AccountGateway.FindByID(input.AccountToID)
	if err != nil {
		return CreateTransactionOutputDTO{}, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return CreateTransactionOutputDTO{}, err
	}

	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return CreateTransactionOutputDTO{}, err
	}

	return CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil
}

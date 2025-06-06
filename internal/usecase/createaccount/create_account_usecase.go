package createaccount

import (
	"github.com/eserafini/wallet-core-ms/internal/entity"
	"github.com/eserafini/wallet-core-ms/internal/gateway"
)

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.FindByID(input.ClientID)
	if err != nil {
		return CreateAccountOutputDTO{}, err
	}

	account, err := entity.NewAccount(client)
	if err != nil {
		return CreateAccountOutputDTO{}, err
	}

	if err := uc.AccountGateway.Save(account); err != nil {
		return CreateAccountOutputDTO{}, err
	}

	return CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}

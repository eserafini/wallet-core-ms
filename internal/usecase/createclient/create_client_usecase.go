package createclient

import (
	"github.com/eserafini/wallet-core-ms/internal/entity"
	"github.com/eserafini/wallet-core-ms/internal/gateway"
)

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDTO) (CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return CreateClientOutputDTO{}, err
	}

	if err := uc.ClientGateway.Save(client); err != nil {
		return CreateClientOutputDTO{}, err
	}

	return CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}

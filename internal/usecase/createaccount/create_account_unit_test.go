package createaccount

import (
	"testing"

	"github.com/eserafini/wallet-core-ms/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) FindByID(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)

	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)

	return args.Error(0)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, err := entity.NewClient("John Doe", "john.doe@example.com")
	assert.Nil(t, err)
	clientGateway := &ClientGatewayMock{}
	clientGateway.On("FindByID", client.ID).Return(client, nil)

	accountGateway := &AccountGatewayMock{}
	accountGateway.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGateway, clientGateway)

	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)

	clientGateway.AssertExpectations(t)
	accountGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "FindByID", 1)
	accountGateway.AssertNumberOfCalls(t, "Save", 1)
}

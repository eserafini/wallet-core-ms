package createtransaction

import (
	"testing"

	"github.com/eserafini/wallet-core-ms/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
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

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, err := entity.NewClient("John Doe", "john.doe@example.com")
	assert.Nil(t, err)
	account1, err := entity.NewAccount(client1)
	assert.Nil(t, err)
	account1.Credit(1000)

	client2, err := entity.NewClient("Jane Doe", "jane.doe@example.com")
	assert.Nil(t, err)
	account2, err := entity.NewAccount(client2)
	assert.Nil(t, err)
	account2.Credit(1000)

	mockAccountGateway := &AccountGatewayMock{}
	mockAccountGateway.On("FindByID", account1.ID).Return(account1, nil)
	mockAccountGateway.On("FindByID", account2.ID).Return(account2, nil)

	mockTransactionGateway := &TransactionGatewayMock{}
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountFromID: account1.ID,
		AccountToID:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockAccountGateway, mockTransactionGateway)

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)
}

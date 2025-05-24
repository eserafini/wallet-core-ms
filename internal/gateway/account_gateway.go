package gateway

import "github.com/eserafini/wallet-core-ms/internal/entity"

type AccountGateway interface {
	FindByID(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}

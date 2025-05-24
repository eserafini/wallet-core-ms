package gateway

import "github.com/eserafini/wallet-core-ms/internal/entity"

type ClientGateway interface {
	FindByID(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

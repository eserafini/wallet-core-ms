package gateway

import "github.com/eserafini/wallet-core-ms/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

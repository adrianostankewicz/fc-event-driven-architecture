package gateway

import "github.com/adrianostankewicz/fc-event-driven-architecture/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

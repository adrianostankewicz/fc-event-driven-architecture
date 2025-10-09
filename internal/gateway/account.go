package gateway

import "github.com/adrianostankewicz/fc-event-driven-architecture/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}

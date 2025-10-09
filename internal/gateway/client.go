package gateway

import "github.com/adrianostankewicz/fc-event-driven-architecture/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

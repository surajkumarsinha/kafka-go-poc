package routing

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	"github.com/surajkumarsinha/kafka-go-poc/routing/producer"
	generalEndpoint "github.com/surajkumarsinha/kafka-go-poc/types/interfaces/endpoints"
)

type IndexEndpoint struct {
	producerEndpoint *producer.ProducerEndpoint
}

func (indexEndPoint *IndexEndpoint) Routes(r *chi.Mux) {
	generalEndpoints := infrastructure.Resolve[generalEndpoint.IGeneralEndpoints]()
	indexEndPoint.producerEndpoint.Routes(r)

	r.Post("/health", generalEndpoints.Health)
}

var (
	indexEndPoint     *IndexEndpoint
	indexEndPointOnce sync.Once
)

func InitRoutes() *IndexEndpoint {
	if indexEndPoint == nil {
		indexEndPointOnce.Do(func() {
			indexEndPoint = &IndexEndpoint{
				producerEndpoint: &producer.ProducerEndpoint{},
			}
		})
	}
	return indexEndPoint
}

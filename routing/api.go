package routing

import (
	"github.com/go-chi/chi/v5"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	producerEndpoint "github.com/surajkumarsinha/kafka-go-poc/types/interfaces/endpoints/producer"
)

func Routes(r *chi.Mux) {
	producerEndpoint := infrastructure.Resolve[producerEndpoint.IProducerEndpoint]()

	r.Post("/producer/produce", producerEndpoint.Produce)
}

package interfaces

import "github.com/go-chi/chi/v5"

type IChiRouter interface {
	InitRouter(func(*chi.Mux)) *chi.Mux
}

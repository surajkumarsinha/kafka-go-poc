package infrastructure

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure/interfaces"
)

type Router struct{}

// Init Router takes routerFunc which is an anonymous function which takes in *chi.Mux as the argument
func (router *Router) InitRouter(routerFunc func(*chi.Mux)) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(customMiddlewares.ResponseFormatter)
	routerFunc(r)
	return r
}

var (
	m          *Router
	routerOnce sync.Once
)

func ChiRouter() interfaces.IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &Router{}
		})
	}
	return m
}

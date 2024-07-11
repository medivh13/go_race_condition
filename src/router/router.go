package router

import (
	handler "go_race_condition/src/handlers"

	"github.com/go-chi/chi"
)

func NewRouter(handler *handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/stock", handler.GetStockHandler)
	r.Post("/update-stock", handler.UpdateStock)
	r.Post("/update-stock-wg", handler.UpdateStockWithWaitGroupHandler)
	r.Post("/update-stock-mutex", handler.UpdateStockWithMutexHandler)
	return r
}

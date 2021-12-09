package routers

import (
	middlewares "etby/services/gateway/internal/middlwares"

	"github.com/go-chi/chi"
)

func Register(path string) chi.Router {
	router := chi.NewRouter()
	middlewares.AddToNoAuth(path)
	return router
}

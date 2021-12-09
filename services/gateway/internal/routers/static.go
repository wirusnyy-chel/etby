package routers

import "github.com/go-chi/chi"

func Static(path string) chi.Router {
	router := chi.NewRouter()

	return router
}

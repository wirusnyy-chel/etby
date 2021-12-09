package internal

import (
	middlewares "etby/services/gateway/internal/middlwares"
	"etby/services/gateway/internal/routers"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func Mux() *chi.Mux {
	mux := chi.NewMux()

	//middleware initialization
	//mux.Use(middleware.Logger)
	mux.Use(initCors().Handler)
	mux.Use(middlewares.Authorization())
	mux.Use(middlewares.Redirect())

	//routers init
	mux.Mount("/login", routers.Login("/login"))
	mux.Mount("/register", routers.Register("/register"))
	mux.Mount("/static", routers.Static("/static"))
	mux.Mount("/query", routers.Query("/query"))
	mux.Mount("/playground", routers.Playground("/playground"))
	return mux
}
func initCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:3000"},
		AllowCredentials: true,
		Debug:            false})
}

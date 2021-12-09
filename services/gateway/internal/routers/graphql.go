package routers

import (
	"etby/services/gateway/graph"
	"etby/services/gateway/graph/generated"
	middlewares "etby/services/gateway/internal/middlwares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func Query(path string) chi.Router {
	router := chi.NewRouter()

	router.Handle("/", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))
	return router
}
func Playground(path string) chi.Router {
	router := chi.NewRouter()
	middlewares.AddToNoAuth(path)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	return router
}

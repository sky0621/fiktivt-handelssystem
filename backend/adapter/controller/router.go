package controller

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func SetRoute(w driver.Web) {
	w.SetRoute("/", playgroundHandler())
	w.SetRoute("/graphql", grapqlHandler())
}

func playgroundHandler() http.HandlerFunc {
	h := handler.Playground("fiktivt-handelssystem-graphql-playground", "/graphql")
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func grapqlHandler() http.HandlerFunc {
	h := handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}}))
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", playgroundHandler())
	r.Handle("/graphql", grapqlHandler())
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func playgroundHandler() http.HandlerFunc {
	h := handler.Playground("fiktivt-handelssystem-graphql-playground", "/graphql")
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func grapqlHandler() http.HandlerFunc {
	h := handler.GraphQL(driver.NewExecutableSchema(driver.Config{Resolvers: &controller.Resolver{}}))
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

package driver

import (
	"log"
	"net/http"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"

	"github.com/99designs/gqlgen/handler"

	"github.com/sky0621/fiktivt-handelssystem/config"

	"github.com/gorilla/mux"
)

func NewWeb(cfg config.Config, adapter controller.GraphQLAdapter) Web {
	r := mux.NewRouter()
	// TODO: basic middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})
	r.Handle("/", playgroundHandler())
	r.Handle("/graphql", grapqlHandler(adapter))

	return &web{cfg: cfg, router: r}
}

type Web interface {
	Start() error
}

type web struct {
	cfg    config.Config
	router *mux.Router
}

func (w *web) Start() error {
	if err := http.ListenAndServe(w.cfg.WebConfig.ListenPort, w.router); err != nil {
		log.Println(err) // TODO: カスタムロガー使う？
		return err
	}
	return nil
}

func playgroundHandler() http.HandlerFunc {
	h := handler.Playground("fiktivt-handelssystem-graphql-playground", "/graphql")
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func grapqlHandler(adapter controller.GraphQLAdapter) http.HandlerFunc {
	h := handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{adapter: adapter}}))
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

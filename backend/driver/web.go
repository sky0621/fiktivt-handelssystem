package driver

import (
	"log"
	"net/http"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"

	"github.com/99designs/gqlgen/handler"

	"github.com/sky0621/fiktivt-handelssystem/config"

	"github.com/gorilla/mux"
)

func NewWeb(cfg config.Config, resolver controller.ResolverRoot) Web {
	r := mux.NewRouter()
	// TODO: basic middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})
	r.Handle("/", playgroundHandler())
	r.Handle("/graphql", grapqlHandler(resolver))

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
	lp := w.cfg.WebConfig.ListenPort
	log.Println(lp)
	if err := http.ListenAndServe(lp, w.router); err != nil {
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

func grapqlHandler(resolver controller.ResolverRoot) http.HandlerFunc {
	h := handler.GraphQL(controller.NewExecutableSchema(controller.Config{Resolvers: resolver}))
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

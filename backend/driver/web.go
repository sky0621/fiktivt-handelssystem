package driver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"

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
	h := handler.GraphQL(
		controller.NewExecutableSchema(controller.Config{Resolvers: resolver}),
		handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
			fmt.Println("*************************************************")
			fmt.Println("called RequestMiddleware")
			fmt.Println(ctx)
			rctx := ctx.Value("request_context")
			fmt.Println(rctx)
			fmt.Println("*************************************************")
			return next(ctx)
		}),
		handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			fmt.Println("=================================================")
			fmt.Println("called ResolverMiddleware")
			fmt.Println(ctx)
			rctx := ctx.Value("request_context")
			fmt.Println(rctx)
			fmt.Println("=================================================")
			return next(ctx)
		}),
		handler.ErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
			if appErr, ok := e.(*gqlerror.Error); ok {
				return appErr
			}
			return &gqlerror.Error{
				Message:    e.Error(),
				Extensions: nil,
			}
		}),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
			e, ok := err.(error)
			if ok {
				fmt.Println("graphql: recover panic")
			} else {
				fmt.Println("graphql: recover panic")
			}
			return &gqlerror.Error{
				Message:    e.Error(),
				Extensions: nil,
			}
		}),
		handler.CacheSize(0),
		handler.ComplexityLimit(100),
		// MEMO: for APQ
		//handler.EnablePersistedQueryCache(),
		handler.IntrospectionEnabled(true),
		handler.Tracer(graphql.NopTracer{}),
		handler.UploadMaxMemory(1024),
		handler.UploadMaxSize(2048),
	)
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

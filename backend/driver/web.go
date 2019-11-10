package driver

import (
	"context"
	"net/http"

	"github.com/sky0621/fiktivt-handelssystem/system"

	"github.com/rs/cors"

	"github.com/go-chi/chi"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"

	"github.com/99designs/gqlgen/handler"

	"github.com/sky0621/fiktivt-handelssystem/config"
)

func NewWeb(cfg config.Config, resolver controller.ResolverRoot, logger system.AppLogger) Web {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Handle("/", playgroundHandler())
	r.Handle("/graphql", graphqlHandler(resolver, logger))

	return &web{cfg: cfg, router: r, logger: logger}
}

type Web interface {
	Start() error
}

type web struct {
	cfg    config.Config
	router chi.Router
	logger system.AppLogger
}

func (w *web) Start() error {
	lgr := w.logger.NewLogger("Start")

	lp := w.cfg.WebConfig.ListenPort
	lgr.Info().Str("ListenPort", lp).Send()
	if err := http.ListenAndServe(lp, w.router); err != nil {
		lgr.Err(err)
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

func graphqlHandler(resolver controller.ResolverRoot, logger system.AppLogger) http.HandlerFunc {
	lgr := logger.NewLogger("graphqlHandler")

	h := handler.GraphQL(
		controller.NewExecutableSchema(controller.Config{Resolvers: resolver}),
		handler.RequestMiddleware(func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
			lgr.Info().Msg("called RequestMiddleware")
			return next(ctx)
		}),
		handler.ResolverMiddleware(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			//lgr.Info().Msg("called ResolverMiddleware")
			return next(ctx)
		}),
		handler.ErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
			lgr.Info().Err(e)
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
				lgr.Info().Err(e)
				lgr.Info().Msg("graphql: recover panic")
			} else {
				lgr.Info().Msg("graphql: recover panic")
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

package driver

import (
	"log"
	"net/http"

	"github.com/sky0621/fiktivt-handelssystem/config"

	"github.com/gorilla/mux"
)

func NewWeb(cfg config.Config) Web {
	r := mux.NewRouter()
	// TODO: basic middleware
	return &web{cfg: cfg, router: r}
}

type Web interface {
	SetRoute(path string, f http.HandlerFunc)
	Start() error
}

type web struct {
	cfg    config.Config
	router *mux.Router
}

func (w *web) SetRoute(path string, f http.HandlerFunc) {
	w.router.Handle(path, f)
}

func (w *web) Start() error {
	if err := http.ListenAndServe(w.cfg.WebConfig.ListenPort, w.router); err != nil {
		log.Println(err) // TODO: カスタムロガー使う？
		return err
	}
	return nil
}

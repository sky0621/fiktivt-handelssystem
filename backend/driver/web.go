package driver

func NewWeb() Web {
	return &web{}
}

type Web interface {
	Start() error
}

type web struct {
}

func (w *web) Start() error {
	return nil
}

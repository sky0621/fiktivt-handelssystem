package repository

func NewItemCommand() ItemCommand {
	return &itemCommand{}
}

type ItemCommand interface {
}

type itemCommand struct {
}

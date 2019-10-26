package command

type ItemHolder interface {
	CreateItemHolder(input CommandItemHolderModel) string
}

type CommandItemHolderModel struct {
	ID       string
	Name     string
	Nickname string
}

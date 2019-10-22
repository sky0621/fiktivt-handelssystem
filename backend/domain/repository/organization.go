package repository

func NewOrganization(persistence Persistence) Organization {
	return &organization{persistence: persistence}
}

type Organization interface {
}

type organization struct {
	persistence Persistence
}

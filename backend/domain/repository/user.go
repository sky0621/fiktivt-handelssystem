package repository

func NewUser(persistence Persistence) User {
	return &user{persistence: persistence}
}

type User interface {
}

type user struct {
	persistence Persistence
}

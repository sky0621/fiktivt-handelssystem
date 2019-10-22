package repository

func NewOrder(persistence Persistence) Order {
	return &order{persistence: persistence}
}

type Order interface {
}

type order struct {
	persistence Persistence
}

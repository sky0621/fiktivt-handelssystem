package repository

func NewOrderDetail(persistence Persistence) OrderDetail {
	return &orderDetail{persistence: persistence}
}

type OrderDetail interface {
}

type orderDetail struct {
	persistence Persistence
}

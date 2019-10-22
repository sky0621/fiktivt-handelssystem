package domain

import "github.com/sky0621/fiktivt-handelssystem/domain/repository"

func NewOrder(orderDetailRepository repository.OrderDetail, instructionRepository repository.Instruction) Order {
	return &order{
		orderDetailRepository: orderDetailRepository,
		instructionRepository: instructionRepository,
	}
}

type Order interface {
	GetOrders() ([]*OrderModel, error)
}

type order struct {
	orderDetailRepository repository.OrderDetail
	instructionRepository repository.Instruction
}

func (o *order) GetOrders() ([]*OrderModel, error) {
	// FIXME:
	p1 := 12500
	t1 := "たーげっと１"
	p2 := 9700
	t2 := "たーげっと２"
	return []*OrderModel{
		&OrderModel{
			ID:   "odr-id0001",
			Name: "注文１",
			User: &UserModel{
				ID:   "usr-id0001",
				Name: "ユーザA",
				Organizations: []*OrganizationModel{
					{
						ID:                  "org-id0001",
						Name:                "組織１",
						UpperOrganizationID: nil,
						Users:               nil,
					},
				},
			},
			OrderDetails: []*OrderDetail{
				&OrderDetail{
					ID:    "odt-0001",
					Order: nil,
					Instruction: &InstructionModel{
						ID:     "ist-0001",
						Price:  &p1,
						Target: &t1,
					},
				},
				{
					ID:    "odt-0002",
					Order: nil,
					Instruction: &InstructionModel{
						ID:     "ist-0002",
						Price:  &p2,
						Target: &t2,
					},
				},
			},
		},
	}, nil
}

type OrderModel struct {
	ID           string
	Name         string
	User         *UserModel
	OrderDetails []*OrderDetail
}

type OrderDetail struct {
	ID          string
	Order       *Order
	Instruction *InstructionModel
}

type InstructionModel struct {
	ID     string
	Price  *int
	Target *string
}

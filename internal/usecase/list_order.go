package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type OrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() (OrdersOutputDTO, error) {
	orders, err := l.OrderRepository.GetAll()
	if err != nil {
		return OrdersOutputDTO{}, err
	}

	var ordersOutput []OrderOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return OrdersOutputDTO{Orders: ordersOutput}, nil
}

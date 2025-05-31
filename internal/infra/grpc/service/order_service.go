package service

import (
	"context"

	"github.com/enniosimoes/CleanArch/internal/infra/grpc/pb"
	"github.com/enniosimoes/CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrdersUseCase usecase.ListOrderUseCase, // <==== Add ListOrdersUseCase to the constructor
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase, // <==== Initialize ListOrdersUseCase
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.ListOrdersResponse, error) {
	orders, err := s.ListOrdersUseCase.Execute() // <==== The problem is that the ListOrdersUseCase.Execute() returns a struct, not a slice of orders
	if err != nil {
		println("Error listing orders:", err)
		return nil, err
	}

	// fmt.Printf("ListOrders: %v\n", orders)

	// debudging s.ListOrdersUseCase
	// fmt.Printf("ListOrdersUseCase: %v\n", s.ListOrdersUseCase)

	output := make([]*pb.CreateOrderResponse, len(orders.Orders))

	for i, order := range orders.Orders {
		output[i] = &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
	}
	return &pb.ListOrdersResponse{Orders: output}, nil
	// return &pb.ListOrdersResponse{}, nil
}

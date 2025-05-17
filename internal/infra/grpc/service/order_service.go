package service

import (
	"context"
	"log"

	"github.com/wellalencarweb/challenge-clean-architecture/internal/infra/grpc/pb"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	GetAllOrdersUseCase usecase.GetAllOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getAllOrdersUsecase usecase.GetAllOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		GetAllOrdersUseCase: getAllOrdersUsecase,
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

// Função que implementa a listagem de categorias
func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {

	log.Printf("GetAll Orders gRPC")
	orders, err := s.GetAllOrdersUseCase.GetAllOrders()
	if err != nil {
		log.Printf("Erro ao coletar as Orders do useCase (gRPC): %s", err)
		return nil, err
	}

	// Slice que será utilizado para receber as ordens e retornar os valores
	var OrderList []*pb.Order

	//loop para adicionar todas as ordens no slice.
	for _, orderDTO := range orders.Orders {
		order := &pb.Order{
			Id:         orderDTO.ID,
			Price:      float32(orderDTO.Price),
			Tax:        float32(orderDTO.Tax),
			FinalPrice: float32(orderDTO.FinalPrice),
		}
		OrderList = append(OrderList, order)
	}
	//return &pb.CategoryList{Categories: categoriesResponse}, nil
	return &pb.OrderList{Orders: OrderList}, nil

}

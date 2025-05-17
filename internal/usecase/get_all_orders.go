package usecase

import (
	"github.com/wellalencarweb/challenge-clean-architecture/internal/entity"
)

// Struct que será utilizada para retornar todas as ordens.
type OrdersDTO struct {
	Orders []OrderOutputDTO
}

// Struct que representa o caso de uso.
// Será retornada pela função NewGetAllOrdersUseCase quando for chamada a partir do main.
type GetAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

// Função anexada na struct GetAllOrdersUseCase
// Utilizada para coletar todas as ordens utilizando a interface.
func (g *GetAllOrdersUseCase) GetAllOrders() (OrdersDTO, error) {

	AllOrders := OrdersDTO{
		Orders: []OrderOutputDTO{},
	}

	// Recebendo as ordens a partir do repository
	listOrders, err := g.OrderRepository.GetAllOrders()
	if err != nil {
		return OrdersDTO{}, err
	}

	// Loop para ler todas as ordens coletadas e transformar no DTO
	for _, order := range listOrders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		AllOrders.Orders = append(AllOrders.Orders, dto)

	}

	return AllOrders, nil
}

// "Construtor" do UseCase. Será executado no main.go
// Recebe como parâmetros uma interface de OrderRepository e retorna o ponteiro do GetAllOrdersUseCase
func NewGetAllOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

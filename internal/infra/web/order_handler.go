package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wellalencarweb/challenge-clean-architecture/internal/entity"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/usecase"
	"github.com/wellalencarweb/challenge-clean-architecture/pkg/events"
)

// Essa struct contém as interfaces que serão necessárias para interagir com use case e events.
// Será chamada a partir do main.go para criar as rotas no chi
type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

// Função anexada no struct acima. Ele excuta as funções no webserver e interage com o Usecase.
// Essa função utiliza o caso de uso createOrderUseCase para gerar uma nova ordem.
func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Create Order handler")
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Printf("Erro ao criar a ordem.: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Função anexada no struct acima. Ele excuta as funções no webserver e interage com o Usecase.
// Essa função utiliza o caso de uso GetAllOrdersUseCase para retornar todas as ordens.
func (h *WebOrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {

	log.Printf("GetAll Orders handler")
	// Criando o novo caso de uso
	getOrders := usecase.NewGetAllOrdersUseCase(h.OrderRepository)

	// Buscando todas as ordens existentes
	output, err := getOrders.GetAllOrders()
	if err != nil {
		log.Printf("Erro ao coletar os dados do useCase: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Realizando o encode para json
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		log.Printf("Erro ao realizar o encode para json: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Função "construtora"
func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

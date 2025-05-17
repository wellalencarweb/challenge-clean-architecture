package database

import (
	"database/sql"

	"github.com/wellalencarweb/challenge-clean-architecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// Função criada para retornar a consulta de todas as ordens presentes no banco.
// A função foi anexada no OrderRepository e retorna um array de orders
func (r *OrderRepository) GetAllOrders() ([]entity.Order, error) {

	// Neste caso não precisa utilizar o prepare, pois não há a inserção de dados. Sem risco de sql injection.
	rows, err := r.Db.Query("select id, price, tax, final_price from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// criando a lista de structs que vai recever todos as ordens
	orders := make([]entity.Order, 0)
	for rows.Next() {
		var order entity.Order
		err = rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/entity"

	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, err = db.Exec(`
		CREATE TABLE orders (
			id varchar(255) NOT NULL,
			price float NOT NULL,
			tax float NOT NULL,
			final_price float NOT NULL,
			PRIMARY KEY (id)
		)
	`)
	suite.NoError(err)
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	// Limpa a tabela após cada teste
	_, err := suite.Db.Exec("DELETE FROM orders")
	suite.NoError(err)
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("order1", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestListAllOrders() {
	order1, err := entity.NewOrder("order1", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order1.CalculateFinalPrice())

	order2, err := entity.NewOrder("order2", 11.0, 3.0)
	suite.NoError(err)
	suite.NoError(order2.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)
	suite.NoError(repo.Save(order1))
	suite.NoError(repo.Save(order2))

	orderResults, err := repo.GetAllOrders()
	suite.NoError(err)
	suite.Len(orderResults, 2, "deve retornar exatamente 2 pedidos")

	suite.Equal(order1.ID, orderResults[0].ID)
	suite.Equal(order1.Price, orderResults[0].Price)
	suite.Equal(order1.Tax, orderResults[0].Tax)
	suite.Equal(order1.FinalPrice, orderResults[0].FinalPrice)

	suite.Equal(order2.ID, orderResults[1].ID)
	suite.Equal(order2.Price, orderResults[1].Price)
	suite.Equal(order2.Tax, orderResults[1].Tax)
	suite.Equal(order2.FinalPrice, orderResults[1].FinalPrice)
}

package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/entity"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

// Teste para validar a listagem de todos os pedidos.
func (suite *OrderRepositoryTestSuite) TestListAllOrders() {
	// Criação da primeira ordem
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)
	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)

	// Criação da Segunda ordem
	order2, err := entity.NewOrder("456", 11.0, 3.0)
	suite.NoError(err)
	suite.NoError(order2.CalculateFinalPrice())
	repo2 := NewOrderRepository(suite.Db)
	err = repo2.Save(order2)
	suite.NoError(err)

	var orderResult2 entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order2.ID).
		Scan(&orderResult2.ID, &orderResult2.Price, &orderResult2.Tax, &orderResult2.FinalPrice)
	suite.NoError(err)
	suite.Equal(order2.ID, orderResult2.ID)
	suite.Equal(order2.Price, orderResult2.Price)
	suite.Equal(order2.Tax, orderResult2.Tax)
	suite.Equal(order2.FinalPrice, orderResult2.FinalPrice)

	repo3 := NewOrderRepository(suite.Db)
	orderResults, err := repo3.GetAllOrders()
	suite.NoError(err)

	// Testando o primeiro resultado da consulta
	suite.Equal(orderResults[0].ID, orderResult.ID)
	suite.Equal(orderResults[0].Price, orderResult.Price)
	suite.Equal(orderResults[0].Tax, orderResult.Tax)
	suite.Equal(orderResults[0].FinalPrice, orderResult.FinalPrice)

	// Testando o segundo resultado da consulta
	suite.Equal(orderResults[1].ID, orderResult2.ID)
	suite.Equal(orderResults[1].Price, orderResult2.Price)
	suite.Equal(orderResults[1].Tax, orderResult2.Tax)
	suite.Equal(orderResults[1].FinalPrice, orderResult2.FinalPrice)
}

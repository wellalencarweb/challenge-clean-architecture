package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"github.com/wellalencarweb/challenge-clean-architecture/configs"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/event/handler"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/infra/graph"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/infra/grpc/pb"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/infra/grpc/service"
	"github.com/wellalencarweb/challenge-clean-architecture/internal/infra/web/webserver"
	"github.com/wellalencarweb/challenge-clean-architecture/pkg/events"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel(configs.RabbitUser, configs.RabbitPassword, configs.RabbitHost, configs.RabbitPort)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	getAllOrdersUseCase := NewGetAllOrdersUseCase(db)

	// Criação do webserver
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	//rotas que serão criadas
	webserver.AddHandler("POST", "/order", webOrderHandler.Create)
	webserver.AddHandler("GET", "/order", webOrderHandler.GetAllOrders)

	//start do webserver
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// Criação do gRPC
	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *getAllOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// Criação do graphql
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase:  *createOrderUseCase,
		GetAllOrdersUseCase: *getAllOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

// Função que cria a conexão com o RabbitMQ.
func getRabbitMQChannel(RabbitUser string, RabbitPassword string, RabbitHost string, RabbitPort string) *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", RabbitUser, RabbitPassword, RabbitHost, RabbitPort))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

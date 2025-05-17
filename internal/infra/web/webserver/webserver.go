package webserver

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// Struct que representa o WebServer
type WebServer struct {
	Router            chi.Router
	HandlersByMethods map[string]DadosHandler // Alteração para adicionar o verbo HTTP
	WebServerPort     string
}

// Struct auxiliar para armazenar o método que será utilizado para criar a rota.
type DadosHandler struct {
	Handler http.HandlerFunc
	Path    string
}

// Armazena cada rota enviada para ser adicionada na inicialização do webserver
func (s *WebServer) AddHandler(metodo string, path string, handler http.HandlerFunc) {

	dadosHandler := DadosHandler{
		Handler: handler,
		Path:    path,
	}

	s.HandlersByMethods[metodo] = dadosHandler

}

// Função anexada na Struct do WebServer
// Percorre todas as Rotas e adicina ao router, registra o middleware de log e inicia o server
// No momento, suporta apenas os métodos GET e POST.
func (s *WebServer) Start() {

	s.Router.Use(middleware.Logger)
	for metodo, dadosHandler := range s.HandlersByMethods {
		switch metodo {
		case "GET":
			s.Router.Get(dadosHandler.Path, dadosHandler.Handler)
		case "POST":
			s.Router.Post(dadosHandler.Path, dadosHandler.Handler)
		default:
			log.Printf("Erro ao adicionar handler. Método não suportado.")
		}
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}

// Função de "Construtor" que será chamada a partir do arquivo main
func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:            chi.NewRouter(),
		HandlersByMethods: make(map[string]DadosHandler), // Alteração para adicionar o verbo HTTP
		WebServerPort:     serverPort,
	}
}

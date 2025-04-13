package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelfnaves/go-gateway-api/internal/service"
	"github.com/rafaelfnaves/go-gateway-api/internal/web/handlers"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountservice *service.AccountService
	port           string
}

func NewServer(accountservice *service.AccountService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountservice: accountservice,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(*s.accountservice)
	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}

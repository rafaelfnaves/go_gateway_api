package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafaelfnaves/go-gateway-api/internal/service"
	"github.com/rafaelfnaves/go-gateway-api/internal/web/handlers"
	"github.com/rafaelfnaves/go-gateway-api/internal/web/middleware"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	invoiceHandler := handlers.NewInvoiceHandler(s.invoiceService)
	AuthMiddleware := middleware.NewAuthMiddleware(s.accountService)

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(AuthMiddleware.Authenticate)
		r.Post("/invoice", invoiceHandler.Create)
		r.Get("/invoice", invoiceHandler.ListByAccount)
		r.Get("/invoice/{id}", invoiceHandler.GetByID)
	})
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}

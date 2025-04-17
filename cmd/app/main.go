package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rafaelfnaves/go-gateway-api/internal/repository"
	"github.com/rafaelfnaves/go-gateway-api/internal/service"
	"github.com/rafaelfnaves/go-gateway-api/internal/web/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	defer db.Close()

	// Initialize repositories and services
	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)
	invoiceRepository := repository.NewInvoiceRepository(db)
	invoiceService := service.NewInvoiceService(invoiceRepository, *accountService)

	// Start the server
	port := os.Getenv("HTTP_PORT")
	srv := server.NewServer(accountService, invoiceService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server", err)
	}
}

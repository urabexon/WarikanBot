package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urabexon/WarikanBot/internal/infrastructure/handler"
	"github.com/urabexon/WarikanBot/internal/infrastructure/repository"
	"github.com/urabexon/WarikanBot/internal/usecase"
)

func main() {
	eventRepository, err := repository.NewEventRepository("database.db")
	if err != nil {
		log.Fatalf("Failed to create event repository: %v", err)
	}

	payerRepository, err := repository.NewPayerRepository("database.db")
	if err != nil {
		log.Fatalf("Failed to create payer repository: %v", err)
	}

	paymentRepository, err := repository.NewPaymentRepository("database.db")
	if err != nil {
		log.Fatalf("Failed to create payment repository: %v", err)
	}

	paymentUsecase := usecase.NewPaymentUsecase(eventRepository, payerRepository, paymentRepository)

}

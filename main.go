package main

import (
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

	paymentUsecase := usecase.NewPayment(eventRepository, payerRepository, paymentRepository)
	slackCommandHandler := handler.NewSlackCommandHandler(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_SIGNING_SECRET"), paymentUsecase)
	slackEventHandler := handler.NewSlackEventHandler(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_SIGNING_SECRET"), paymentUsecase)

	log.Println("Starting server on 0.0.0.0:5272")
}

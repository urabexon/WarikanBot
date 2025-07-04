package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/urabexon/WarikanBot/internal/infrastructure/handler"
	"github.com/urabexon/WarikanBot/internal/infrastructure/repository"
	"github.com/urabexon/WarikanBot/internal/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env ファイルが見つかりませんでした。")
	}
}

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

	mux := http.NewServeMux()
	mux.Handle("/slack/command", slackCommandHandler)
	mux.Handle("/slack/event", slackEventHandler)
	log.Println("Starting server on 0.0.0.0:5272")
	if err := http.ListenAndServe("0.0.0.0:5272", mux); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

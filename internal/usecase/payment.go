package usecase

import (
	"fmt"

	"github.com/urabexon/WarikanBot/internal/domain/entity"
	"github.com/urabexon/WarikanBot/internal/domain/repository"
	"github.com/urabexon/WarikanBot/internal/domain/valueobject"
)

type PaymentUsecase struct {
	events   repository.EventRepository
	payers   repository.PayerRepository
	payments repository.PaymentRepository
}

func NewPayment(events repository.EventRepository, payers repository.PayerRepository, payments repository.PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{
		events,
		payers,
		payments,
	}
}

type Settlement struct {
	Total           valueobject.Yen
	AmountsAdvanced map[valueobject.PayerID]valueobject.Yen
	Payers          []*entity.Payer
	Instructions    []*SettlementInstruction
}

type SettlementInstruction struct {
	From   valueobject.PayerID
	To     valueobject.PayerID
	Amount valueobject.Yen
}

func (u *PaymentUsecase) Create(eventID valueobject.EventID, payerID valueobject.PayerID, amount valueobject.Yen) (*entity.Payment, error) {
	if eventID.IsUnknown() {
		return nil, valueobject.NewErrorNotFound("eventID is unknown", nil)
	}
	event := &entity.Event{
		ID: eventID,
	}
	if err := u.events.CreateIfNotExists(event); err != nil {
		return nil, fmt.Errorf("failed to create event: %w", err)
	}

	if payerID.IsUnknown() {
		return nil, valueobject.NewErrorNotFound("payerID is unknown", nil)
	}
	payer := &entity.Payer{
		ID:      payerID,
		EventID: eventID,
	}
	if err := u.payers.CreateIfNotExists(payer); err != nil {
		return nil, fmt.Errorf("failed to create payer: %w", err)
	}

	payment := &entity.Payment{
		ID:      valueobject.NewPaymentID(),
		EventID: eventID,
		PayerID: payerID,
		Amount:  amount,
	}
	if err := u.payments.Create(payment); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return payment, nil
}

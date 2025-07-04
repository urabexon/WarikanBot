package entity

import (
	"github.com/urabexon/WarikanBot/internal/domain/valueobject"
)

type Event struct {
	ID valueobject.EventID
}

type Payer struct {
	ID      valueobject.PayerID
	EventID valueobject.EventID
	Weight  valueobject.Percent
}

type Payment struct {
	ID      valueobject.PaymentID
	EventID valueobject.EventID
	PayerID valueobject.PayerID
	Amount  valueobject.Yen
}

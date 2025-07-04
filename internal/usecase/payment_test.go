package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/urabexon/WarikanBot/internal/domain/entity"
	"github.com/urabexon/WarikanBot/internal/domain/valueobject"
)

type MockEventRepository struct{}

func (m *MockEventRepository) CreateIfNotExists(event *entity.Event) error {
	return nil
}

type MockPayerRepository struct {
	Payers []*entity.Payer
}

func (m *MockPayerRepository) Create(payer *entity.Payer) error {
	return nil
}

func (m *MockPayerRepository) CreateIfNotExists(payer *entity.Payer) error {
	return nil
}
func TestSettle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		eventID            valueobject.EventID
		payers             []*entity.Payer
		payments           []*entity.Payment
		expectedSettlement *Settlement
	}{
		{
			name: "no payments",
			eventID: valueobject.NewEventID(),
			payers: []*entity.Payer{
				{ID: valueobject.NewPayerID(), EventID: valueobject.NewEventID(), Weight: valueobject.NewPercent(100)},
			},
		},
	}
}

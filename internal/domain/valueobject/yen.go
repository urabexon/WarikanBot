package valueobject

import (
	"errors"
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type Yen int64

func NewYen(amount int) (Yen, error) {
	if amount < 0 {
		return 0, errors.New("amount cannot be negative")
	}
	return Yen(amount), nil
}

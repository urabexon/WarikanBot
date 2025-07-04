package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urabexon/WarikanBot/internal/domain/entity"
	"github.com/urabexon/WarikanBot/internal/domain/valueobject"
	"github.com/slack-go/slack"
)

func parseYen(text string) (valueobject.Yen, error) {

}

func parsePercent(text string) (valueobject.Percent, error) {

}

func botProfiles() slack.MsgOption {

}

func buildPaymentCreatedMessage(userID string, amount valueobject.Yen) slack.MsgOption {

}

func buildPayerJoinedMessage(userID string) slack.MsgOption {

}

func buildPayerAlreadyJoinedMessage(userID string) slack.MsgOption {

}

func buildSettlementMessage(settlement *usecase.Settlement) slack.MsgOption {

}

func buildHelpMessage() slack.MsgOption {

}

func buildInvalidCommandMessage(userID string) slack.MsgOption {

}

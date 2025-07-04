package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/slack-go/slack"
	"github.com/urabexon/WarikanBot/internal/domain/valueobject"
	"github.com/urabexon/WarikanBot/internal/usecase"
)

const SlackMetadataEventType = "warikan"

type SlackCommandHandler struct {
	Token string
	TeamID string
	TeamDomain string
	ChannelID string
	ChannelName string
}

type CommandHandler struct {
	paymentUsecase *usecase.PaymentUsecase
}

func (h *SlackCommandHandler) handleWarikanCommand(slash slack.SlashCommand) error {
	eventID := valueobject.NewEventID(slash.ChannelID)
	payerID := valueobject.NewPayerID(slash.UserID)

	
}
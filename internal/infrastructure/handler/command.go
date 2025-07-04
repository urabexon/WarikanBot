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


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
	signingSecret  string
	client         *slack.Client
	paymentUsecase *usecase.PaymentUsecase
	amountPattern  *regexp.Regexp
	joinPattern    *regexp.Regexp
	percentPattern *regexp.Regexp
	settlePattern  *regexp.Regexp
	helpPattern    *regexp.Regexp
}

func NewSlackCommandHandler(token string, signingSecret string, paymentUsecase *usecase.PaymentUsecase) *SlackCommandHandler {
	return &SlackCommandHandler{
		client:         slack.New(token),
		signingSecret:  signingSecret,
		paymentUsecase: paymentUsecase,
		amountPattern:  regexp.MustCompile(`\b((?:\d{1,3}(?:,\d{3})+|\d+))円?\b`),
		joinPattern:    regexp.MustCompile(`\b(?:(?i:join)|参加|払う|払います)\b`),
		percentPattern: regexp.MustCompile(`\b(\d+)(?:%)?\b`),
		settlePattern:  regexp.MustCompile(`\b(?:(?i:settle)|集計|集金|合計)\b`),
		helpPattern:    regexp.MustCompile(`\b(?:(?i:help)|(?i:h)|ヘルプ|使い方)\b`),
	}
}

func (h *SlackCommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	verifier, err := slack.NewSecretsVerifier(r.Header, h.signingSecret)
	if err != nil {
		http.Error(w, "Failed to create secrets verifier", http.StatusBadRequest)
		return
	}
	r.Body = io.NopCloser(io.TeeReader(r.Body, &verifier))
	slash, err := slack.SlashCommandParse(r)
	if err != nil {
		http.Error(w, "Failed to parse slash command", http.StatusBadRequest)
		return
	}
	if verifier.Ensure() != nil {
		http.Error(w, "Invalid request signature", http.StatusUnauthorized)
		return
	}

	err = h.handleSlashCommand(slash)
	if e := new(valueobject.ErrorNotFound); errors.As(err, &e) {
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	}
	if e := new(valueobject.ErrorAlreadyExists); errors.As(err, &e) {
		http.Error(w, e.Error(), http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, "Failed to handle slash command", http.StatusInternalServerError)
	}
}

func (h *SlackCommandHandler) handleSlashCommand(slash slack.SlashCommand) error {
	switch slash.Command {
	case "/warikan":
		return h.handleWarikanCommand(slash)
	default:
		return fmt.Errorf("unsupported command: %s", slash.Command)
	}
}

func (h *SlackCommandHandler) handleWarikanCommand(slash slack.SlashCommand) error {
	eventID := valueobject.NewEventID(slash.ChannelID)
	payerID := valueobject.NewPayerID(slash.UserID)

	if h.joinPattern.MatchString(slash.Text) {
		weight := valueobject.Percent(100)
		percentMatch := h.percentPattern.FindStringSubmatch(slash.Text)
		if percentMatch != nil {
			w, err := parsePercent(percentMatch[1])
			if err != nil {
				return err
			}
			weight = w
		}
		_, err := h.paymentUsecase.Join(eventID, payerID, weight)
		if e := new(valueobject.ErrorAlreadyExists); errors.As(err, &e) {
			_, _, err = h.client.PostMessage(slash.ChannelID, buildPayerAlreadyJoinedMessage(slash.UserID), botProfiles())
			return err
		}
		if err != nil {
			return err
		}

		_, _, err = h.client.PostMessage(slash.ChannelID, buildPayerJoinedMessage(slash.UserID), botProfiles())
		return err
	}

	match := h.amountPattern.FindStringSubmatch(slash.Text)
	if match != nil {
		amount, err := parseYen(match[1])
		if err != nil {
			return err
		}

		payment, err := h.paymentUsecase.Create(eventID, payerID, amount)
		if err != nil {
			return err
		}

		_, _, err = h.client.PostMessage(slash.ChannelID, buildPaymentCreatedMessage(slash.UserID, amount),
			slack.MsgOptionMetadata(slack.SlackMetadata{
				EventType: SlackMetadataEventType,
				EventPayload: map[string]any{
					"payment_id": payment.ID.String(),
				},
			}),
			botProfiles(),
		)

		return err
	}

	if h.settlePattern.MatchString(slash.Text) {
		settlement, err := h.paymentUsecase.Settle(eventID)
		if err != nil {
			log.Println(err)
			return err
		}
		_, _, err = h.client.PostMessage(slash.ChannelID, buildSettlementMessage(settlement), botProfiles())
		if err != nil {
			log.Println(err)
		}
		return err
	}

	if h.helpPattern.MatchString(slash.Text) {
		_, _, err := h.client.PostMessage(slash.ChannelID, buildHelpMessage(), botProfiles())
		return err
	}

	_, _, err := h.client.PostMessage(slash.ChannelID, buildInvalidCommandMessage(slash.UserID), botProfiles())
	return err
}

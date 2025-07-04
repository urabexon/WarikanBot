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
	rawYen := strings.ReplaceAll(text, ",", "")
	amount, err := strconv.Atoi(rawYen)
	if err != nil {
		return valueobject.Yen(0), fmt.Errorf("failed to parse amount: %w", err)
	}
	yen, err := valueobject.NewYen(amount)
	if err != nil {
		return valueobject.Yen(0), err
	}
	return yen, nil
}

func parsePercent(text string) (valueobject.Percent, error) {
	percent, err := strconv.Atoi(text)
	if err != nil {
		return valueobject.Percent(0), fmt.Errorf("failed to parse percent: %w", err)
	}
	percentValue, err := valueobject.NewPercent(percent)
	if err != nil {
		return valueobject.Percent(0), err
	}
	return percentValue, nil
}

func botProfiles() slack.MsgOption {
	return slack.MsgOptionCompose(
		slack.MsgOptionIconEmoji(":money_with_wings:"),
		slack.MsgOptionUsername("割り勘"),
	)
}

func buildPaymentCreatedMessage(userID string, amount valueobject.Yen) slack.MsgOption {
	return slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			slack.NewTextBlockObject("mrkdwn", fmt.Sprintf(":receipt: <@%s>さんが%s立て替えました！", userID, amount.String()), false, false),
			nil,
			nil,
		),
	)
}

func buildPayerJoinedMessage(userID string) slack.MsgOption {
	return slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			slack.NewTextBlockObject("mrkdwn", fmt.Sprintf(":purse: <@%s>さんが割り勘に参加します！", userID), false, false),
			nil,
			nil,
		),
	)
}

func buildPayerAlreadyJoinedMessage(userID string) slack.MsgOption {
	return slack.MsgOptionCompose(
		slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				slack.NewTextBlockObject("mrkdwn", fmt.Sprintf(":warning: <@%s>さんはすでに割り勘に参加しています！", userID), false, false),
				nil,
				nil,
			),
		),
		slack.MsgOptionPostEphemeral(userID),
	)
}

func buildSettlementMessage(settlement *usecase.Settlement) slack.MsgOption {
	blocks := []slack.Block{
		slack.NewHeaderBlock(
			slack.NewTextBlockObject("plain_text", ":moneybag: 集計結果", false, false),
		),
	}
	payerAmountFields := []*slack.TextBlockObject{}
	
}

func buildHelpMessage() slack.MsgOption {

}

func buildInvalidCommandMessage(userID string) slack.MsgOption {

}

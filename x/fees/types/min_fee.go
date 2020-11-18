package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinFee contains the minimum amount of coins that should be paid as a fee for each message of the specified type sent
type MinFee struct {
	MessageType string    `json:"message_type" yaml:"message_type"`
	Amount      sdk.Coins `json:"amount" yaml:"amount"`
}

func NewMinFee(messageType string, amount sdk.Coins) MinFee {
	return MinFee{
		MessageType: messageType,
		Amount:      amount,
	}
}

// Validate check if the min fee parameters are valid
func (mf MinFee) Validate() error {
	if mf.MessageType == "" {
		return fmt.Errorf("invalid minimum fee message type")
	}

	if !mf.Amount.IsValid() {
		return fmt.Errorf("invalid minimum fee amount")
	}

	return nil
}
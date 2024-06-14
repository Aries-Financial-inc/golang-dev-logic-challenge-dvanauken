package model

import (
	"testing"
)

func TestOptionsContract(t *testing.T) {
	contract := OptionsContract{
		Type:           "call",
		StrikePrice:    100,
		Bid:            10,
		Ask:            12,
		ExpirationDate: "2023-12-31",
		LongOrShort:    "long",
	}

	if contract.Type != "call" && contract.Type != "put" {
		t.Errorf("Invalid type: %s", contract.Type)
	}

	if contract.Bid >= contract.Ask {
		t.Errorf("Invalid bid-ask spread: Bid %f, Ask %f", contract.Bid, contract.Ask)
	}
}

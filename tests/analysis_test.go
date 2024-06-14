package tests

import (
	"bytes"
	"encoding/json"
	"golang-dev-logic-challenge-dvanauken/controllers" // Make sure this path is correct
	"golang-dev-logic-challenge-dvanauken/model"       // Make sure this path is correct
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOptionsContractModelValidation(t *testing.T) {
	// Example test: ensure that the model rejects invalid option types
	contract := model.OptionsContract{
		Type:           "invalid_type", // assuming valid types are "call" or "put"
		StrikePrice:    100.0,
		Bid:            5.0,
		Ask:            5.5,
		ExpirationDate: "2021-12-31",
		LongOrShort:    "long",
	}

	if contract.Type != "call" && contract.Type != "put" {
		t.Errorf("Contract type validation failed. Expected 'call' or 'put', got %s", contract.Type)
	}
}

func TestAnalysisEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.AnalysisHandler)) // Ensure AnalyzeHandler is exported
	defer server.Close()

	contracts := []model.OptionsContract{{
		Type:           "call",
		StrikePrice:    100,
		Bid:            10,
		Ask:            15,
		ExpirationDate: "2023-01-01",
		LongOrShort:    "long",
	}}
	body, _ := json.Marshal(contracts)
	req, _ := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))

	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}

	var xyValues []model.XYValue
	_ = json.NewDecoder(resp.Body).Decode(&xyValues)

	expectedY := calculateExpectedY(contracts[0])
	if xyValues[0].Y != expectedY {
		t.Errorf("Expected Y value %v; got %v", expectedY, xyValues[0].Y)
	}
}
func calculateExpectedY(contract model.OptionsContract) float64 {
	return 0 // Implement based on your specific logic
}

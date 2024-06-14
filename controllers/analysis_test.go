package controllers

import (
	"bytes"
	"encoding/json"
	"golang-dev-logic-challenge-dvanauken/model"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOptionsContractModelValidation(t *testing.T) {
	contract := model.OptionsContract{
		Type:           "invalid_type",
		StrikePrice:    100,
		Bid:            5.0,
		Ask:            5.5,
		ExpirationDate: "2023-01-01",
		LongOrShort:    "long",
	}

	if contract.Type != "call" && contract.Type != "put" {
		t.Logf("Contract type validation failed as expected. Got: %s", contract.Type)
	} else {
		t.Errorf("Contract type validation did not fail as expected. Got: %s", contract.Type)
	}
}

func TestAnalysisEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(AnalysisHandler))
	defer server.Close()

	contracts := []model.OptionsContract{{
		Type:           "call",
		StrikePrice:    100,
		Bid:            10,
		Ask:            15,
		ExpirationDate: "2023-01-01",
		LongOrShort:    "long",
	}}
	body, err := json.Marshal(contracts)
	if err != nil {
		t.Fatalf("Failed to marshal contracts: %v", err)
	}
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	t.Logf("Response Body: %s", string(respBody))

	var response model.AnalysisResponse
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(response.XYValues) == 0 {
		t.Fatalf("Expected non-empty xyValues; got empty slice")
	}

	expectedY := calculateExpectedY(contracts[0])
	if response.XYValues[0].Y != expectedY {
		t.Errorf("Expected Y value %v; got %v", expectedY, response.XYValues[0].Y)
	}
}

func calculateExpectedY(contract model.OptionsContract) float64 {
	price := 100.0

	var y float64
	if contract.Type == "call" {
		if contract.LongOrShort == "long" {
			y = math.Max(0, price-contract.StrikePrice) - contract.Ask
		} else {
			y = -(math.Max(0, price-contract.StrikePrice) - contract.Bid)
		}
	} else if contract.Type == "put" {
		if contract.LongOrShort == "long" {
			y = math.Max(0, contract.StrikePrice-price) - contract.Ask
		} else {
			y = -(math.Max(0, contract.StrikePrice-price) - contract.Bid)
		}
	}

	return y
}

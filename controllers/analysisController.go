package controllers

import (
	"encoding/json"
	"golang-dev-logic-challenge-dvanauken/model"
	"math"
	"net/http"
)

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	var contracts []model.OptionsContract
	err := json.NewDecoder(r.Body).Decode(&contracts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := model.AnalysisResponse{
		XYValues:        calculateXYValues(contracts),
		MaxProfit:       calculateMaxProfit(contracts),
		MaxLoss:         calculateMaxLoss(contracts),
		BreakEvenPoints: calculateBreakEvenPoints(contracts),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func calculateXYValues(contracts []model.OptionsContract) []model.XYValue {
	var xyValues []model.XYValue
	priceRange := 100.0
	step := 1.0

	for _, contract := range contracts {
		for price := contract.StrikePrice - priceRange; price <= contract.StrikePrice+priceRange; price += step {
			var y float64
			if contract.Type == "call" {
				if contract.LongOrShort == "long" {
					y = math.Max(0, price-contract.StrikePrice) - contract.Ask
				} else {
					y = -(math.Max(0, price-contract.StrikePrice) - contract.Bid)
				}
			} else {
				if contract.LongOrShort == "long" {
					y = math.Max(0, contract.StrikePrice-price) - contract.Ask
				} else {
					y = -(math.Max(0, contract.StrikePrice-price) - contract.Bid)
				}
			}
			xyValues = append(xyValues, model.XYValue{X: price, Y: y})
		}
	}
	return xyValues
}

func calculateMaxProfit(contracts []model.OptionsContract) float64 {
	maxProfit := 0.0
	xyValues := calculateXYValues(contracts)
	for _, xy := range xyValues {
		if xy.Y > maxProfit {
			maxProfit = xy.Y
		}
	}
	return maxProfit
}

func calculateMaxLoss(contracts []model.OptionsContract) float64 {
	maxLoss := 0.0
	xyValues := calculateXYValues(contracts)
	for _, xy := range xyValues {
		if xy.Y < maxLoss {
			maxLoss = xy.Y
		}
	}
	return maxLoss
}

func calculateBreakEvenPoints(contracts []model.OptionsContract) []float64 {
	var breakEvenPoints []float64
	xyValues := calculateXYValues(contracts)
	for _, xy := range xyValues {
		if math.Abs(xy.Y) < 0.01 {
			breakEvenPoints = append(breakEvenPoints, xy.X)
		}
	}
	return breakEvenPoints
}

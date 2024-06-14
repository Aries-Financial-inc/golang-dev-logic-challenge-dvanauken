package routes

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OptionsContract structure for the request body
type OptionsContract struct {
	Type           string  `json:"type"`            // 'call' or 'put'
	StrikePrice    float64 `json:"strike_price"`    // Strike price of the option
	Bid            float64 `json:"bid"`             // Bid price
	Ask            float64 `json:"ask"`             // Ask price
	ExpirationDate string  `json:"expiration_date"` // Expiration date of the option
	LongOrShort    string  `json:"long_or_short"`   // 'long' or 'short'
}

// AnalysisResult structure for the response body
type AnalysisResult struct {
	GraphData       []GraphPoint `json:"graph_data"`
	MaxProfit       float64      `json:"max_profit"`
	MaxLoss         float64      `json:"max_loss"`
	BreakEvenPoints []float64    `json:"break_even_points"`
}

// GraphPoint structure for X & Y values of the risk & reward graph
type GraphPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		var contracts []OptionsContract

		if err := c.ShouldBindJSON(&contracts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Your code here
		results := analyzeOptionsContracts(contracts)

		c.JSON(http.StatusOK, results)

		c.JSON(http.StatusOK, gin.H{"message": "Analysis completed successfully"})
	})

	return router
}

func analyzeOptionsContracts(contracts []OptionsContract) AnalysisResult {
	graphData := calculateXYValues(contracts)
	maxProfit := calculateMaxProfit(contracts)
	maxLoss := calculateMaxLoss(contracts)
	breakEvenPoints := calculateBreakEvenPoints(contracts)

	return AnalysisResult{
		GraphData:       graphData,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}
}

func calculateXYValues(contracts []OptionsContract) []GraphPoint {
	// Simplified logic for XY calculation
	var graphData []GraphPoint
	for _, contract := range contracts {
		for x := 0.0; x <= 200.0; x += 10.0 {
			y := 0.0
			if contract.Type == "call" {
				y = math.Max(0, x-contract.StrikePrice) - contract.Ask
			} else if contract.Type == "put" {
				y = math.Max(0, contract.StrikePrice-x) - contract.Ask
			}
			if contract.LongOrShort == "short" {
				y = -y
			}
			graphData = append(graphData, GraphPoint{X: x, Y: y})
		}
	}
	return graphData
}

func calculateMaxProfit(contracts []OptionsContract) float64 {
	maxProfit := math.Inf(-1)
	for _, contract := range contracts {
		profit := 0.0
		if contract.Type == "call" {
			profit = math.Max(0, 200-contract.StrikePrice) - contract.Ask
		} else if contract.Type == "put" {
			profit = math.Max(0, contract.StrikePrice-0) - contract.Ask
		}
		if contract.LongOrShort == "short" {
			profit = -profit
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func calculateMaxLoss(contracts []OptionsContract) float64 {
	maxLoss := 0.0
	for _, contract := range contracts {
		loss := contract.Ask
		if contract.LongOrShort == "short" {
			loss = contract.Bid
		}
		if loss > maxLoss {
			maxLoss = loss
		}
	}
	return maxLoss
}

func calculateBreakEvenPoints(contracts []OptionsContract) []float64 {
	breakEvenPoints := []float64{}
	for _, contract := range contracts {
		if contract.Type == "call" {
			breakEven := contract.StrikePrice + contract.Ask
			breakEvenPoints = append(breakEvenPoints, breakEven)
		} else if contract.Type == "put" {
			breakEven := contract.StrikePrice - contract.Ask
			breakEvenPoints = append(breakEvenPoints, breakEven)
		}
	}
	return breakEvenPoints
}

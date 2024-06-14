package model

// OptionsContract defines the structure for an option contract.
type OptionsContract struct {
	Type           string  `json:"type"`            // 'call' or 'put'
	StrikePrice    float64 `json:"strike_price"`    // Strike price of the option
	Bid            float64 `json:"bid"`             // Bid price
	Ask            float64 `json:"ask"`             // Ask price
	ExpirationDate string  `json:"expiration_date"` // Expiration date of the option
	LongOrShort    string  `json:"long_or_short"`   // 'long' or 'short'
}

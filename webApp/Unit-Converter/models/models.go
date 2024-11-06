package models

type Form struct {
	Type     string  `json:"type"`
	Value    float64 `json:"value"`
	To       string  `json:"to"`
	From     string  `json:"from"`
	Response float64 `json:"response"`
}

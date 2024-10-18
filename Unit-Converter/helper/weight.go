package helper

import "fmt"

// ConvWeight method converts weight params
func ConvWeight(value float64, from, to string) (float64, error) {
	var gram float64
	var weight float64

	switch from {
	case "mg":
		gram = value / 1000
	case "kg":
		gram = value * 1000
	case "ounce":
		gram = value * 28.3495
	case "pound":
		gram = value * 453.592
	case "g":
		gram = value
	default:
		return 0, fmt.Errorf("invalid meter: %s", from)
	}

	switch to {
	case "mg":
		weight = gram * 1000
	case "kg":
		weight = gram / 1000
	case "ounce":
		weight = gram * 0.0353
	case "pound":
		weight = gram * 0.0022
	case "g":
		weight = gram
	default:
		return 0, fmt.Errorf("invalid meter: %s", to)
	}

	return weight, nil
}

package helper

import "fmt"

// ConvTemp method converts temperature params
func ConvTemp(value float64, from, to string) (float64, error) {
	var kelvin float64
	var temp float64

	switch from {
	case "f":
		kelvin = ((value - 32) * (5.0 / 9.0)) + 273.15
	case "c":
		kelvin = value + 273.15
	case "k":
		kelvin = value
	default:
		return 0.0, fmt.Errorf("invalid meter: %s", from)
	}

	switch to {
	case "f":
		temp = ((kelvin - 273.15) * (9.0 / 5.0)) + 32
	case "c":
		temp = kelvin - 273.15
	case "k":
		temp = kelvin
	default:
		return 0.0, fmt.Errorf("invalid meter: %s", to)
	}
	return temp, nil
}

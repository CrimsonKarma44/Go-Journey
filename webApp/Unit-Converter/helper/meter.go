package helper

import "fmt"

// ConvLength method converts length params
func ConvLength(value float64, from, to string) (float64, error) {
	var meter float64
	switch from {
	case "cm":
		meter = value / 100
	case "mm":
		meter = value / 1000
	case "km":
		meter = value * 1000
	case "inch":
		meter = value * 0.0254
	case "ft":
		meter = value * 0.3048
	case "yard":
		meter = value * 0.9144
	case "mile":
		meter = value * 1609.34
	case "m":
		meter = value
	default:
		return 0, fmt.Errorf("invalid meter: %s", from)
	}

	var response float64
	switch to {
	case "m":
		response = meter
	case "cm":
		response = meter * 100
	case "mm":
		response = meter * 1000
	case "km":
		response = meter / 1000
	case "inch":
		response = meter * 39.37
	case "ft":
		response = meter * 3.28084
	case "yard":
		response = meter * 1.09361
	case "mile":
		response = meter * 0.000621371
	default:
		return 0, fmt.Errorf("invalid meter: %s", to)
	}
	return response, nil
}

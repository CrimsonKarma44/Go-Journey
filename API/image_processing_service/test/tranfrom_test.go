package utils_test

import (
	"IPS/img/models"
	"IPS/utils"
	"encoding/json"
	"testing"
)

func TestTransformJSONParser(t *testing.T) {
	input := models.TransformRequest{
		Resize:    &models.Resize{Width: 100, Height: 200},
		Crop:      &models.Crop{Width: 50, Height: 50, X: 10, Y: 10},
		Rotate:    &models.Rotate{Degree: 90},
		Format:    &models.Format{FileExt: "jpg"},
		Filter:    &models.Filter{GrayScale: true, Sepia: false},
		Watermark: &models.Watermark{Text: "Test", FontSize: 12, Opacity: 0.5, Position: "bottom-right"},
		Flip:      &models.Flip{Condition: true},
		Mirror:    &models.Mirror{Condition: true},
		Compress:  &models.Compress{Quality: 80},
	}

	result := utils.TransformJSONParser(input)

	expectedJSON, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal input: %v", err)
	}

	if string(result) != string(expectedJSON) {
		t.Errorf("Expected %s, got %s", string(expectedJSON), string(result))
	}
}

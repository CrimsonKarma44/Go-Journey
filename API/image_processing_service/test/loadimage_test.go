package utils_test

import (
	"IPS/utils"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoadImage(t *testing.T) {
	// Prepare a sample image path (ensure this file exists for the test to pass)
	imgPath := "testdata/sample.png"

	// Create a dummy image file if not present
	if _, err := os.Stat(imgPath); os.IsNotExist(err) {
		// utils.InitLocalStore(strings.Split(imgPath, "/")[0])
		fmt.Println(" Creating test image directory")
		if os.IsExist(os.MkdirAll(strings.Split(imgPath, "/")[0], os.ModePerm)) {
			t.Fatalf("Failed to create directory for test image: %v", err)
		}
	}

	imgBytes, format, err := utils.LoadImage(imgPath)
	if err != nil {
		t.Fatalf("LoadImage failed: %v", err)
	}

	if len(imgBytes) == 0 {
		t.Errorf("Expected non-empty image bytes")
	}

	if format == "" {
		t.Errorf("Expected non-empty format string")
	}

	fmt.Printf("Loaded image format: %s, size: %d bytes\n", format, len(imgBytes))
}

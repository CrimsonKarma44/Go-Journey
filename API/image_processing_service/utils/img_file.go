package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/bimg"
)

// SaveImageFromRequest saves an image from an HTTP request to the specified path.
func SaveImageFromRequest(r *http.Request, savePath string) error {
	file, _, err := r.FormFile("image")
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Detect image format from extension
	ext := strings.ToLower(filepath.Ext(savePath))
	img, format, err := image.Decode(file)
	if err != nil {
		return err
	}

	if ext == ".png" || format == "png" {
		return png.Encode(out, img)
	} else if ext == ".jpg" || ext == ".jpeg" || format == "jpeg" {
		return jpeg.Encode(out, img, nil)
	}
	return nil
}

// LoadImage loads an image from disk and returns the image.Image and its format.
func LoadImage(path string) ([]byte, string, error) {

	file, err := bimg.Read(path)
	if err != nil {
		return nil, "", err
	}

	newImage, err := bimg.NewImage(file).Resize(800, 800)
	return newImage, "format", nil
}

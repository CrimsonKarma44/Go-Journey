package models

import "C"
import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/h2non/bimg"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Filename string `gorm:"unique;not null"`
	Path     string `gorm:"not null"`
}

func (i *Image) ImageData() *ImageData {
	pwd, _ := os.Getwd()
	fileData, err := bimg.Read(pwd + i.Path + i.Filename)
	if err != nil {
		return nil
	}

	format := bimg.NewImage(fileData).Type()
	data := fileData
	return &ImageData{
		Name:      i.Filename,
		Path:      i.Path,
		Data:      data,
		Format:    format,
		CreatedAt: i.CreatedAt,
	}
}

type ImageData struct {
	Name      string
	Path      string
	Data      []byte
	Format    string
	CreatedAt time.Time
}

func (imgData *ImageData) getData() []byte {
	return imgData.Data
}

func (imgData *ImageData) getFormat() bimg.ImageType {
	switch imgData.Format {
	case "jpg", "jpeg":
		return bimg.JPEG
	case "png":
		return bimg.PNG
	case "gif":
		return bimg.GIF
	case "webp":
		return bimg.WEBP
	default:
		return bimg.JPEG
	}
}

func (imgData *ImageData) getFileSize() int {
	return len(imgData.Data)
}
func (imgData *ImageData) HttpResponse() []byte {
	js, err := json.Marshal(map[string]string{
		"filename":    imgData.Name,
		"url":         imgData.Path + imgData.Name,
		"size":        fmt.Sprintf("%d", imgData.getFileSize()),
		"contentType": fmt.Sprintf("image/%s", imgData.Format),
		"uploadedAt":  imgData.CreatedAt.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		fmt.Errorf("Error marshalling image: %s", err)
		return nil
	}
	return js
}

func (imgData *ImageData) Resize(r Resize) error {
	var err error
	imgData.Data, err = bimg.NewImage(imgData.getData()).Resize(r.Width, r.Height)
	if err != nil {
		return fmt.Errorf("error Resizing")
	}
	return nil
}

func (imgData *ImageData) Crop(c Crop) error {
	var err error

	options := bimg.Options{
		Width:  c.Width,
		Height: c.Height,
		Left:   c.X,
		Top:    c.Y,
	}

	imgData.Data, err = bimg.NewImage(imgData.getData()).Process(options)
	if err != nil {
		return fmt.Errorf("error Croping")
	}
	return nil
}

func (imgData *ImageData) Rotate(r Rotate) error {
	var err error
	imgData.Data, err = bimg.NewImage(imgData.getData()).Rotate(bimg.Angle(r.Degree))
	if err != nil {
		return fmt.Errorf("error Resizing")
	}
	return nil
}

func (imgData *ImageData) FormatChange(f Format) error {
	var err error

	if imgData.getFormat() != f.GetBIMGType() {
		imgData.Data, err = bimg.NewImage(imgData.getData()).Convert(f.GetBIMGType())
		if err != nil {
			return fmt.Errorf("error Resizing")
		}
	}
	return nil
}

func (imgData *ImageData) Filter(f Filter) error {
	return nil
}

func (imgData *ImageData) Flip(f Flip) error {
	var err error
	if f.Condition {
		imgData.Data, err = bimg.NewImage(imgData.getData()).Flip()
		if err != nil {
			return fmt.Errorf("error Fliping")
		}
	}
	return nil
}
func (imgData *ImageData) Watermark(w Watermark) error {
	if w.Text == "" {
		return nil
	}

	meta, err := bimg.NewImage(imgData.getData()).Metadata()
	if err != nil {
		return fmt.Errorf("cannot get metadata: %w", err)
	}

	if meta.Size.Width < 20 || meta.Size.Height < 20 {
		return fmt.Errorf("image too small for watermark: %dx%d", meta.Size.Width, meta.Size.Height)
	}

	originalLen := len(imgData.Data)
	_, options := w.ToBimgWatermark()

	// Sanity check on options
	if options.Watermark.DPI > 300 {
		options.Watermark.DPI = 300
	}
	if options.Watermark.Margin > meta.Size.Width/4 {
		options.Watermark.Margin = meta.Size.Width / 10
	}

	newData, err := bimg.NewImage(imgData.getData()).Process(options)
	if err != nil {
		return fmt.Errorf("error applying watermark: %w", err)
	}

	if len(newData) == 0 {
		return fmt.Errorf("processed image data is empty")
	}

	fmt.Printf("Watermark applied: Original=%d, New=%d bytes\n", originalLen, len(newData))
	imgData.Data = newData
	return nil
}

func (imgData *ImageData) Mirror(m Mirror) error {
	var err error

	if m.Condition {
		imgData.Data, err = bimg.NewImage(imgData.getData()).Flop()
		if err != nil {
			return fmt.Errorf("error Fliping")
		}
	}
	return nil
}

func (imgData *ImageData) Compress(c Compress) error {
	var err error

	options := bimg.Options{
		Quality: c.Quality,
	}
	imgData.Data, err = bimg.NewImage(imgData.getData()).Process(options)
	if err != nil {
		return fmt.Errorf("error Compressing")
	}

	return nil
}

func (imgData *ImageData) RunAll(t *TransformRequest) error {
	var err error

	if err = imgData.Resize(*t.Resize); err != nil {
		return fmt.Errorf("resize error: %s", err)
	}

	if err = imgData.Crop(*t.Crop); err != nil {
		return fmt.Errorf("crop error: %s", err)
	}

	if err = imgData.Rotate(*t.Rotate); err != nil {
		return fmt.Errorf("rotate error: %s", err)
	}

	if err = imgData.FormatChange(*t.Format); err != nil {
		return fmt.Errorf("FormatChange error: %s", err)
	}

	//if err = imgData.Filter(*t.Filter); err != nil {
	//	fmt.Println("Filter error:", err)
	//	return err
	//}

	if err = imgData.Flip(*t.Flip); err != nil {
		return fmt.Errorf("flip error: %s", err)
	}

	if err = imgData.Mirror(*t.Mirror); err != nil {
		return fmt.Errorf("mirror error: %s", err)
	}

	if err = imgData.Compress(*t.Compress); err != nil {
		return fmt.Errorf("compress error: %s", err)
	}

	//if err = imgData.Watermark(*t.Watermark); err != nil {
	//	fmt.Println("Watermark error:", err)
	//	return err
	//}

	return nil
}

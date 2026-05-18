package models

import "github.com/h2non/bimg"

type TransformRequest struct {
	Resize    *Resize    `json:"resize,omitempty"`
	Crop      *Crop      `json:"crop,omitempty"`
	Rotate    *Rotate    `json:"rotate,omitempty"`
	Format    *Format    `json:"format,omitempty"`
	Filter    *Filter    `json:"filter,omitempty"`
	Watermark *Watermark `json:"watermark,omitempty"`
	Flip      *Flip      `json:"flip,omitempty"`
	Mirror    *Mirror    `json:"mirror,omitempty"`
	Compress  *Compress  `json:"compress,omitempty"`
}

type Resize struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}
type Crop struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
	X      int `json:"x,omitempty"`
	Y      int `json:"y,omitempty"`
}

type Rotate struct {
	Degree int `json:"degree,omitempty"`
}

type Format struct {
	FileExt string `json:"file_ext,omitempty"`
}

func (f *Format) GetBIMGType() bimg.ImageType {
	switch f.FileExt {
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

type Filter struct {
	GrayScale bool `json:"gray_scale"`
	Sepia     bool `json:"sepia"`
}

type Watermark struct {
	Text     string  `json:"text"`
	FontSize int     `json:"font_size"`
	Opacity  float64 `json:"opacity"`
	Position string  `json:"position"` // e.g., "bottom-right", "center", "top-left"
}

// ToBimgWatermark converts your Watermark to bimg.Watermark and returns
// the corresponding bimg.Options (which includes Gravity)
func (w Watermark) ToBimgWatermark() (bimg.Watermark, bimg.Options) {
	// Map FontSize to DPI (approximate)
	dpi := w.FontSize * 6
	if dpi < 10 {
		dpi = 10
	}
	if dpi > 200 {
		dpi = 200
	}
	// Opacity: float64 → float32, clamp to [0,1]
	opacity := float32(w.Opacity)
	if opacity > 1.0 {
		opacity = 1.0
	} else if opacity < 0.0 {
		opacity = 0.0
	}

	// Default margin (you can make this configurable)
	margin := 10

	// Map Position to Gravity
	gravity := bimg.GravityCentre // default
	switch w.Position {
	case "top":
		gravity = bimg.GravityNorth
	case "left":
		gravity = bimg.GravityWest
	case "right":
		gravity = bimg.GravityEast
	case "bottom":
		gravity = bimg.GravitySouth
	default:
		// Keep center or log warning
	}

	bimgWM := bimg.Watermark{
		Text:        w.Text,
		DPI:         dpi,
		Opacity:     opacity,
		Margin:      margin,
		Font:        "", // use default font (or specify like "Arial")
		NoReplicate: false,
		// Background: bimg.Color{R: 255, G: 255, B: 255}, // optional background
	}

	options := bimg.Options{
		Watermark: bimgWM,
		Gravity:   gravity, // This controls position!
	}

	return bimgWM, options
}

// Flip flips Horizontally
type Flip struct {
	Condition bool `json:"condition"`
}

// Mirror flips vertically
type Mirror struct {
	Condition bool `json:"condition"`
}

type Compress struct {
	Quality int `json:"quality"`
}

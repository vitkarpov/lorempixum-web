package lorempixum

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"math/rand"
)

const (
	whiteIndex = iota
	blackIndex
)

var palette = []color.Color{color.White, color.Black}

func GetImage(width, height int) *image.Paletted {
	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			r := rand.Float32()
			if r > 0.5 {
				img.SetColorIndex(i, j, blackIndex)
			} else {
				img.SetColorIndex(i, j, whiteIndex)
			}
		}
	}

	return img
}

func StreamImage(out io.Writer, img *image.Paletted, format string) error {
	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, nil)
	}
	return errors.New("Unsupported format")
}

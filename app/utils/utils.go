package utils

import (
	"fmt"
	"image"
	"strings"

	"golang.org/x/image/draw"
)

func ImageURL(imageURL string) string {
	if strings.HasPrefix(imageURL, "https://") {
		return imageURL
	}

	return fmt.Sprintf("/public/assets/%s", imageURL)
}

// Scale will use the provided scaler to resize src to the bounds in rect
// It will always return an RGBA image
func Scale(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

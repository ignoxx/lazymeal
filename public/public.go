package public

import (
	"bytes"
	"embed"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"lazymeal/app/utils"
	"strings"

	"golang.org/x/image/draw"
)

//go:embed assets
var AssetsFS embed.FS
var PreviewImages = map[string]string{}

func init() {
	dir, err := AssetsFS.ReadDir("assets")
	if err != nil {
		fmt.Println("oopsie", err)
	}

	for _, f := range dir {
		if f.IsDir() {
			continue
		}

		fi, _ := f.Info()
		name := fi.Name()
		if strings.HasSuffix(name, "jpg") || strings.HasSuffix(name, "png") {
			imageFile, err := AssetsFS.Open("assets/" + name)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// 2. scale the image down
			img, format, err := image.Decode(imageFile)

			if err != nil {
				fmt.Println(err)
			}

			scaledImg := utils.Scale(img, image.Rect(img.Bounds().Min.X/2-10, img.Bounds().Min.Y/2-10, 10, 10), draw.BiLinear)

			rawScaledBuffer := bytes.NewBuffer([]byte{})
			if format == "png" {
				if err := png.Encode(rawScaledBuffer, scaledImg); err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				if err := jpeg.Encode(rawScaledBuffer, scaledImg, nil); err != nil {
					fmt.Println(err)
					continue
				}
			}

			if len(rawScaledBuffer.Bytes()) > 0 {
				PreviewImages[name] = fmt.Sprintf("data:image/%s;base64, %s", format, base64.StdEncoding.EncodeToString(rawScaledBuffer.Bytes()))
			}
		}
	}

	fmt.Println("Images prepared for preview: ", len(PreviewImages))
}

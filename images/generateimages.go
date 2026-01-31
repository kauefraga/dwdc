package images

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"

	"github.com/fogleman/gg"
)

func GenerateImagesFromText(text string) {
	characters := []rune(text)
	desktopPath := getUserDesktopDirectoryPath()

	for i, c := range characters {
		charString := string(c)

		width := 400
		height := 400

		img := image.NewRGBA(image.Rect(0, 0, width, height))

		cyan := color.RGBA{100, 200, 200, 0xff}

		for x := range width {
			for y := range height {
				img.Set(x, y, cyan)
			}
		}

		dc := gg.NewContextForRGBA(img)

		if err := dc.LoadFontFace("dmserifdisplay.ttf", 300); err != nil {
			log.Fatal("failed to load font")
		}

		dc.SetColor(color.Black)
		dc.DrawStringAnchored(charString, float64(dc.Width())/2, float64(dc.Height())/2, 0.5, 0.5)

		outputPath := filepath.Join(desktopPath, "dwdc-"+fmt.Sprint(i)+".png")

		if err := dc.SavePNG(outputPath); err != nil {
			log.Fatal("failed to save output image")
		}
	}
}

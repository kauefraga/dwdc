package images

import (
	"fmt"
	"image"
	"image/color"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/kauefraga/dwdc/settings"
)

func GenerateMultipleImages(s *settings.Settings) error {
	deleteDwdcImages()

	characters := []rune(s.Text)
	desktopPath := getUserDesktopDirectoryPath()

	for i, c := range characters {
		charString := string(c)

		width := 400
		height := 400

		img := image.NewRGBA(image.Rect(0, 0, width, height))

		bg := color.RGBA{
			R: s.BackgroundColor[0],
			G: s.BackgroundColor[1],
			B: s.BackgroundColor[2],
			A: s.BackgroundColor[3],
		}
		fg := color.RGBA{
			R: s.Color[0],
			G: s.Color[1],
			B: s.Color[2],
			A: s.Color[3],
		}

		for x := range width {
			for y := range height {
				img.Set(x, y, bg)
			}
		}

		dc := gg.NewContextForRGBA(img)

		if err := dc.LoadFontFace(s.FontFamily, 300); err != nil {
			return fmt.Errorf("failed to load font: %w", err)
		}

		dc.SetColor(fg)
		dc.DrawStringAnchored(charString, float64(dc.Width())/2, float64(dc.Height())/2, 0.5, 0.5)

		outputPath := filepath.Join(desktopPath, "dwdc-"+fmt.Sprint(i)+".png")

		if err := dc.SavePNG(outputPath); err != nil {
			return fmt.Errorf("failed to save output image: %w", err)
		}
	}

	return nil
}

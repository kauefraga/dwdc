package images

import (
	"fmt"
	"image/color"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/kauefraga/dwdc/settings"
)

const DefaultMargin = 300

type Coordinate struct {
	X float64
	Y float64
}

func BuildPositionMap(width, height int, paddingXPercentage, paddingYPercentage float64) map[string]Coordinate {
	w := float64(width)
	h := float64(height)
	px := w * paddingXPercentage
	py := h * paddingYPercentage

	return map[string]Coordinate{
		settings.PosTopLeft:   {px, py},
		settings.PosTopCenter: {w / 2, py},
		settings.PosTopRight:  {w - px, py},

		settings.PosCenterLeft:  {px, h / 2},
		settings.PosCenter:      {w / 2, h / 2},
		settings.PosCenterRight: {w - px, h / 2},

		settings.PosBottomLeft:   {px, h - py},
		settings.PosBottomCenter: {w / 2, h - py},
		settings.PosBottomRight:  {w - px, h - py},
	}
}

func GenerateOneImage(s *settings.Settings) error {
	desktopPath := getUserDesktopDirectoryPath()

	width := 2560
	height := 1440

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

	coords := BuildPositionMap(width, height, 0.25, 0.2)

	c, ok := coords[s.TextPosition]
	if !ok {
		c = coords[settings.PosCenter]
	}

	dc := gg.NewContext(width, height)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetColor(bg)
	dc.Fill()

	if err := dc.LoadFontFace(s.FontFamily, float64(s.FontSize)); err != nil {
		return fmt.Errorf("failed to load font: %w", err)
	}
	dc.SetColor(fg)
	dc.DrawStringAnchored(s.Text, c.X, c.Y, 0.5, 0.5)

	outputPath := filepath.Join(desktopPath, "dwdc-background.png")

	if err := dc.SavePNG(outputPath); err != nil {
		return fmt.Errorf("failed to save output image: %w", err)
	}

	return nil
}

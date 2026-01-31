package settings

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	BackgroundColor [4]uint8
	Color           [4]uint8
	FontFamily      string
	Text            string
	Template        string
}

func New() *Settings {
	var s Settings

	_, err := toml.DecodeFile("dwdc.toml", &s)
	if err != nil {
		fmt.Println("failed to load dwdc.toml:", err)
	}

	return &s
}

func (s *Settings) ApplyDefaults() *Settings {
	if s.BackgroundColor == [4]uint8{0, 0, 0, 0} {
		s.BackgroundColor = [4]uint8{100, 200, 200, 255}
	}

	if s.Color == [4]uint8{0, 0, 0, 0} {
		s.Color = [4]uint8{0, 0, 0, 255}
	}

	if s.FontFamily == "" {
		s.FontFamily = "leaguespartan.ttf"
	}

	if s.Text == "" {
		s.Template = "dayoftheweek"
	}

	return s
}

func (s *Settings) LogSettings() {
	fmt.Println("bg:", s.BackgroundColor)
	fmt.Println("fg:", s.Color)
	fmt.Println("fontfamily:", s.FontFamily)
	fmt.Println("text:", s.Text)
	fmt.Println("template:", s.Template)
}

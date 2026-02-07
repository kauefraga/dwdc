package settings

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kauefraga/dwdc/templates"
)

const (
	DayOfTheWeek    = "dayoftheweek"
	StaticText      = "text"
	BackgroundImage = "backgroundimage"
)

const (
	PosTopLeft      = "topleft"
	PosTopCenter    = "topcenter"
	PosTopRight     = "topright"
	PosCenterLeft   = "centerleft"
	PosCenter       = "center"
	PosCenterRight  = "centerright"
	PosBottomLeft   = "bottomleft"
	PosBottomCenter = "bottomcenter"
	PosBottomRight  = "bottomright"
)

type Settings struct {
	BackgroundColor [4]uint8
	Color           [4]uint8
	Template        string
	Text            string
	TextPosition    string
	FontFamily      string
	FontSize        uint
}

func New() *Settings {
	var s Settings

	// TODO check if toml file exists first

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

	if s.Text == "" && s.Template == "" {
		s.Template = DayOfTheWeek
	}

	if s.TextPosition == "" {
		s.TextPosition = PosCenter
	}

	if s.FontSize == 0 {
		s.FontSize = 250
	}

	return s
}

func (s *Settings) GetTextFromTemplate() *Settings {
	switch s.Template {
	case DayOfTheWeek:
		s.Text = templates.GetDayOfTheWeek()
	case StaticText, BackgroundImage:
		// These cases are based in text
		if s.Text == "" {
			log.Fatalln("[\033[31m error \033[0m] This template needs Text to be set. Try 'Text = \"You are worth it\"'.")
		}
	default:
		fmt.Println("[\033[33m warning \033[0m] This template does not exist, the default (day of the week) is being set.")
		s.Template = DayOfTheWeek
		s.Text = templates.GetDayOfTheWeek()
	}

	return s
}

// Returns a string prepended with ansi escape sequence for green fg and appended with reset ansi
func colorOutputValue(value any) string {
	return fmt.Sprint("\033[32m", value, "\033[0m")
}

func (s *Settings) LogSettings() {
	fmt.Println("bg:", colorOutputValue(s.BackgroundColor))
	fmt.Println("fg:", colorOutputValue(s.Color))
	fmt.Println("template:", colorOutputValue(s.Template))
	fmt.Println("fontfamily:", colorOutputValue(s.FontFamily))
	fmt.Println("fontsize:", colorOutputValue(s.FontSize))
	fmt.Println("text:", colorOutputValue(s.Text))
	fmt.Println("textposition:", colorOutputValue(s.TextPosition))
}

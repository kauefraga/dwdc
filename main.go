package main

import (
	"log"

	"github.com/kauefraga/dwdc/images"
	"github.com/kauefraga/dwdc/settings"
)

func main() {
	s := settings.
		New().
		ApplyDefaults().
		GetTextFromTemplate()

	s.LogSettings()

	var err error

	// default case is being handled in `settings.GetTextFromTemplate`, so it's not possible here
	switch s.Template {
	case settings.DayOfTheWeek, settings.StaticText:
		err = images.GenerateMultipleImages(s)
	case settings.BackgroundImage:
		err = images.GenerateOneImage(s)
	}

	if err != nil {
		log.Fatalln(err)
	}
}

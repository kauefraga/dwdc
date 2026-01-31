package main

import (
	"github.com/kauefraga/dwdc/images"
	"github.com/kauefraga/dwdc/settings"
)

func main() {
	s := settings.New().ApplyDefaults()

	s.LogSettings()

	images.DeleteDwdcImages()
	images.GenerateImagesFromSettings(s)
}

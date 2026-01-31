package main

import (
	"github.com/kauefraga/dwdc/images"
	"github.com/kauefraga/dwdc/templates"
)

func main() {
	images.DeleteDwdcImages()
	dayOfWeekToday := templates.GetDayOfWeek()
	images.GenerateImagesFromText(dayOfWeekToday)
}

package templates

import (
	"time"
)

var weekdaysInPortuguese []string = []string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}

func GetDayOfWeek() string {
	today := int(time.Now().Weekday())

	return weekdaysInPortuguese[today]
}

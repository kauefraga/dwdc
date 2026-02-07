package templates

import "time"

var weekdaysInPortuguese = []string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}

func GetDayOfTheWeek() string {
	today := int(time.Now().Weekday())

	return weekdaysInPortuguese[today]
}

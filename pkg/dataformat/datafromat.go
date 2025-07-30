package dataformat

import (
	"fmt"
	"time"
)

var monthNames = map[time.Month]string{
	time.January:   "Январь",
	time.February:  "Февраль",
	time.March:     "Март",
	time.April:     "Апрель",
	time.May:       "Май",
	time.June:      "Июнь",
	time.July:      "Июль",
	time.August:    "Август",
	time.September: "Сентябрь",
	time.October:   "Октябрь",
	time.November:  "Ноябрь",
	time.December:  "Декабрь",
}

func FormatDateRu(t time.Time) string {
	return fmt.Sprintf("%s %d, %d", monthNames[t.Month()], t.Day(), t.Year())
}

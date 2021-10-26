package utils

import "time"

func ParseDayOfWeek(number time.Weekday) string {
	switch number {
	case 1:
		return "lunedi"
	case 2:
		return "martedi"
	case 3:
		return "mercoledi"
	case 4:
		return "giovedi"
	case 5:
		return "venerdi"
	case 6:
		return "sabato"
	case 0:
		return "domenica"
	default:
		return ""
	}
}

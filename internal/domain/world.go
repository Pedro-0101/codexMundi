package domain

import "time"

// World represents the global environment and state of the simulation.
type World struct {
	Date      time.Time
	Era       *Era
	Countries []*Country
}

func NewWorld(date time.Time, era *Era, countries []*Country) *World {
	return &World{
		Date:      date,
		Era:       era,
		Countries: countries,
	}
}

// GetSeason returns the current season based on Northern Hemisphere logic.
func (w *World) GetSeason() string {
	month := w.Date.Month()
	switch month {
	case time.December, time.January, time.February:
		return "Inverno"
	case time.March, time.April, time.May:
		return "Primavera"
	case time.June, time.July, time.August:
		return "Verão"
	case time.September, time.October, time.November:
		return "Outono"
	default:
		return "Desconhecida"
	}
}

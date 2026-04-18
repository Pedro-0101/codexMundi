package main

import (
	"fmt"
	"os"
	"time"

	"codexMundi/internal/domain"
	"codexMundi/internal/engine"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize Era
	medieval := &domain.Era{Name: "Era Medieval"}

	// Initialize the Global World
	world := &domain.World{
		Date:     time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC),
		Era:      medieval,
		Speed:    1,
		IsPaused: false,
	}

	// Initialize the Country
	c := &domain.Country{
		ID:   "my_country",
		Name: "Nova Esperança",
		Politics: &domain.Politics{
			Regime: "Monarquia",
			Leader: "Pedro I",
		},
		Economy: &domain.Economy{
			GDP: 1000.0,
		},
		Population: &domain.Population{
			Total: 500,
		},
	}

	p := tea.NewProgram(engine.NewModel(c, world), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

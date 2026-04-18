package main

import (
	"fmt"
	"os"

	"codexMundi/internal/engine"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// 2. Initialize and Configure Engine
	eg := engine.GetInstance()
	eg.InitializeSimulation()

	// 3. Start TUI
	p := tea.NewProgram(engine.NewModel(eg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

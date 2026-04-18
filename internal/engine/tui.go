package engine

import (
	"fmt"
	"strings"
	"time"


	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type Model struct {
	Engine    *Engine
	Logs      []string
	TextInput textinput.Model
	Quitting  bool
}

func NewModel(e *Engine) Model {
	ti := textinput.New()
	ti.Placeholder = "Digite um comando ou ação... (/help)"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40

	return Model{
		Engine:    e,
		Logs:      []string{"Bem-vindo ao Codex Mundi. O tempo está passando..."},
		TextInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		m.tickCmd(),
	)
}

func (m Model) tickCmd() tea.Cmd {
	// The ticker always runs to keep the UI responsive, 
	// but the Engine decides if it should advance logic.
	// We use a constant UI refresh rate (e.g., 100ms) or 
	// scale it by speed for visual feedback.
	
	speed := m.Engine.Velocity
	if speed == 0 {
		speed = 1 // Min speed to keep the UI ticking/checking for unpause
	}
	
	duration := time.Second / time.Duration(speed)
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.Quitting = true
			return m, tea.Quit
		case tea.KeyEnter:
			input := m.TextInput.Value()
			if input != "" {
				m.handleInput(input)
				m.TextInput.SetValue("")
			}
			return m, nil
		}

		if m.TextInput.Value() == "" {
			switch msg.String() {
			case "p", " ":
				m.Engine.TogglePause()
				return m, m.tickCmd()
			case "1": m.Engine.SetVelocity(1); return m, m.tickCmd()
			case "2": m.Engine.SetVelocity(2); return m, m.tickCmd()
			case "4": m.Engine.SetVelocity(4); return m, m.tickCmd()
			case "8": m.Engine.SetVelocity(8); return m, m.tickCmd()
			}
		}

	case tickMsg:
		// Logic advancement is now encapsulated in the Engine
		m.Engine.UpdateTick()
		return m, m.tickCmd()
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m *Model) handleInput(input string) {
	if strings.HasPrefix(input, "/") {
		parts := strings.Fields(input)
		cmd := parts[0]
		switch cmd {
		case "/pause":
			m.Engine.TogglePause()
		case "/speed":
			if len(parts) > 1 {
				var s int
				fmt.Sscanf(parts[1], "%d", &s)
				m.Engine.SetVelocity(s)
			}
		default:
			m.Logs = append(m.Logs, "Comando desconhecido: "+cmd)
		}
		return
	}
	m.Logs = append(m.Logs, "Ação Narrativa: "+input)
}

// Styles
var (
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			Bold(true)

	statStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	logStyle = lipgloss.NewStyle().
			Height(10).
			Padding(1)
)

func (m Model) View() string {
	if m.Quitting {
		return "Até logo!\n"
	}

	w := m.Engine.World
	c := m.Engine.Country

	// Header
	worldInfo := fmt.Sprintf("Era: %s | Data: %s | Estação: %s", 
		w.Era.Name, 
		w.Date.Format("02/01/2006"), 
		w.GetSeason(),
	)
	header := headerStyle.Render("CODEX MUNDI | " + worldInfo)
	
	stats := fmt.Sprintf("\n Líder: %s | PIB: %s | Pop: %s | Velocidade: %dx %s\n",
		c.Politics.Leader,
		statStyle.Render(fmt.Sprintf("%.2f", c.Economy.GDP)),
		statStyle.Render(fmt.Sprintf("%d", c.Population.Total)),
		m.Engine.Velocity,
		func() string { if m.Engine.IsPaused() { return "(PAUSADO)" }; return "" }(),
	)

	// Logs
	start := len(m.Logs) - 10
	if start < 0 { start = 0 }
	logs := strings.Join(m.Logs[start:], "\n")
	body := logStyle.Render(logs)

	return header + stats + body + "\n" + m.TextInput.View() + "\n (p) Pausar | (1,2,4,8) Velocidade | (Esc) Sair"
}

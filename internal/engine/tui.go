package engine

import (
	"fmt"
	"strings"
	"time"

	"codexMundi/internal/domain"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type Model struct {
	Country   *domain.Country
	World     *domain.World
	Logs      []string
	TextInput textinput.Model
	Quitting  bool
}

func NewModel(c *domain.Country, w *domain.World) Model {
	ti := textinput.New()
	ti.Placeholder = "Digite um comando ou ação... (/help)"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40

	return Model{
		Country:   c,
		World:     w,
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
	if m.World.IsPaused {
		return nil
	}
	duration := time.Second / time.Duration(m.World.Speed)
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
				m.World.IsPaused = !m.World.IsPaused
				m.Logs = append(m.Logs, fmt.Sprintf("Jogo %s", func() string {
					if m.World.IsPaused { return "Pausado" }
					return "Retomado"
				}()))
				return m, m.tickCmd()
			case "1": m.World.Speed = 1; m.Logs = append(m.Logs, "Velocidade: 1x"); return m, m.tickCmd()
			case "2": m.World.Speed = 2; m.Logs = append(m.Logs, "Velocidade: 2x"); return m, m.tickCmd()
			case "4": m.World.Speed = 4; m.Logs = append(m.Logs, "Velocidade: 4x"); return m, m.tickCmd()
			case "8": m.World.Speed = 8; m.Logs = append(m.Logs, "Velocidade: 8x"); return m, m.tickCmd()
			}
		}

	case tickMsg:
		if !m.World.IsPaused {
			// Advance World Date
			m.World.AdvanceDate()
			// Deterministic Growth
			m.Country.Economy.GDP *= 1.001
			m.Country.Population.Total += int64(m.World.Speed)
			return m, m.tickCmd()
		}
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
			m.World.IsPaused = !m.World.IsPaused
			m.Logs = append(m.Logs, "Controle: Pausa alternada")
		case "/speed":
			if len(parts) > 1 {
				fmt.Sscanf(parts[1], "%d", &m.World.Speed)
				m.Logs = append(m.Logs, fmt.Sprintf("Controle: Velocidade ajustada para %dx", m.World.Speed))
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

	// Header Logic
	worldInfo := fmt.Sprintf("Era: %s | Data: %s | Estação: %s", 
		m.World.Era.Name, 
		m.World.Date.Format("02/01/2006"), 
		m.World.GetSeason(),
	)
	header := headerStyle.Render("CODEX MUNDI | " + worldInfo)
	
	stats := fmt.Sprintf("\n Líder: %s | PIB: %s | Pop: %s | Velocidade: %dx %s\n",
		m.Country.Politics.Leader,
		statStyle.Render(fmt.Sprintf("%.2f", m.Country.Economy.GDP)),
		statStyle.Render(fmt.Sprintf("%d", m.Country.Population.Total)),
		m.World.Speed,
		func() string { if m.World.IsPaused { return "(PAUSADO)" }; return "" }(),
	)

	// Logs
	start := len(m.Logs) - 10
	if start < 0 { start = 0 }
	logs := strings.Join(m.Logs[start:], "\n")
	body := logStyle.Render(logs)

	return header + stats + body + "\n" + m.TextInput.View() + "\n (p) Pausar | (1,2,4,8) Velocidade | (Esc) Sair"
}

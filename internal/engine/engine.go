package engine

import (
	"codexMundi/internal/domain"
	"time"
)

type Engine struct {
	Clock *Clock
	World *domain.World
}

var instance *Engine

func GetInstance() *Engine {
	if instance == nil {
		startTime := time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC)
		instance = &Engine{
			Clock: NewClock(startTime),
		}
		// Start the clock in a background goroutine
		go instance.Clock.Start()
	}
	return instance
}

func (e *Engine) SetState(w *domain.World) {
	e.World = w
}

func (e *Engine) InitializeSimulation() {
	medieval := &domain.Era{Name: "Era Medieval"}

	countries := []*domain.Country{
		domain.NewCountry("Brasil", &domain.Politics{Leader: "D. Pedro II"}, &domain.Economy{GDP: 5000}, &domain.Population{Total: 1000000}),
		domain.NewCountry("Japão", &domain.Politics{Leader: "Meiji"}, &domain.Economy{GDP: 7000}, &domain.Population{Total: 2000000}),
		domain.NewCountry("Alemanha", &domain.Politics{Leader: "Bismarck"}, &domain.Economy{GDP: 8000}, &domain.Population{Total: 1500000}),
	}

	world := domain.NewWorld(e.Clock.GetCurrentTime(), medieval, countries)
	e.World = world
}

func (e *Engine) TogglePause() {
	e.Clock.TogglePause()
}

func (e *Engine) SetVelocity(v int) {
	e.Clock.SetVelocity(int8(v))
}

func (e *Engine) IsPaused() bool {
	return e.Clock.IsPaused()
}

func (e *Engine) GetVelocity() int {
	return int(e.Clock.GetVelocity())
}

// UpdateTick process a single tick from the clock.
func (e *Engine) UpdateTick(t time.Time) []string {
	var logs []string
	if e.World != nil {
		e.World.Date = t
		for _, c := range e.World.Countries {
			logs = append(logs, c.Update())
		}
	}
	return logs
}

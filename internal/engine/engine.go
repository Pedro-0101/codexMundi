package engine

import (
	"codexMundi/internal/domain"
	"time"
)

type Engine struct {
	Clock   *Clock
	World   *domain.World
	Country *domain.Country
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

func (e *Engine) SetState(w *domain.World, c *domain.Country) {
	e.World = w
	e.Country = c
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
func (e *Engine) UpdateTick(t time.Time) {
	if e.World != nil {
		e.World.Date = t
		// Here we would trigger more domain logic
	}
}

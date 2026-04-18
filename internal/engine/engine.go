package engine

import (
	"codexMundi/internal/domain"
)

type Engine struct {
	Velocity     int
	LastVelocity int
	Tick         int
	World        *domain.World
	Country      *domain.Country
}

var instance *Engine

func GetInstance() *Engine {
	if instance == nil {
		instance = &Engine{
			Velocity:     1,
			LastVelocity: 1,
			Tick:         0,
		}
	}
	return instance
}

func (e *Engine) SetState(w *domain.World, c *domain.Country) {
	e.World = w
	e.Country = c
}

func (e *Engine) TogglePause() {
	if e.Velocity == 0 {
		e.Velocity = e.LastVelocity
	} else {
		e.LastVelocity = e.Velocity
		e.Velocity = 0
	}
}

func (e *Engine) SetVelocity(v int) {
	e.Velocity = v
	if v > 0 {
		e.LastVelocity = v
	}
}

func (e *Engine) IsPaused() bool {
	return e.Velocity == 0
}

func (e *Engine) UpdateTick() {
	if e.Velocity == 0 {
		return
	}

	e.Tick++
	
	// Advance Date (1 tick = 1 day in MVP)
	e.World.AdvanceDate(1)

	// Deterministic Growth
	// GDP grows by 0.1% per tick when things are normal
	e.Country.Economy.GDP *= 1.001
	
	// Population increases based on a base rate scaled by world speed (relative to UI)
	// For now, simple linear growth
	e.Country.Population.Total += int64(e.Velocity)
}

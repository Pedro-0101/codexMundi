package engine

import (
	"sync"
	"time"
)

// Clock manages the game time and generates ticks at a rate determined by velocity.
type Clock struct {
	currentTime time.Time
	velocity    int8
	isPaused    bool
	tickChan    chan time.Time
	mu          sync.RWMutex
}

func NewClock(startTime time.Time) *Clock {
	return &Clock{
		currentTime: startTime,
		velocity:    1,
		isPaused:    true,
		tickChan:    make(chan time.Time, 1),
	}
}

// Start initiates the clock loop. It should be called in a goroutine.
func (c *Clock) Start() {
	ticker := time.NewTicker(time.Second) // Initial value, will be updated
	defer ticker.Stop()

	for {
		c.mu.RLock()
		paused := c.isPaused
		vel := c.velocity
		c.mu.RUnlock()

		if paused {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		// Adjust ticker interval based on velocity
		interval := time.Second / time.Duration(vel)
		ticker.Reset(interval)

		<-ticker.C
		c.mu.Lock()
		c.currentTime = c.currentTime.AddDate(0, 0, 1) // Advance 1 day per tick
		newTime := c.currentTime
		c.mu.Unlock()

		// Non-blocking send to tickChan
		select {
		case c.tickChan <- newTime:
		default:
		}

	}
}

func (c *Clock) GetTickChan() <-chan time.Time {
	return c.tickChan
}

func (c *Clock) SetVelocity(v int8) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v != 1 && v != 2 && v != 4 && v != 8 {
		return
	}
	c.velocity = v
}

func (c *Clock) TogglePause() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.isPaused = !c.isPaused
}

func (c *Clock) IsPaused() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.isPaused
}

func (c *Clock) GetVelocity() int8 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.velocity
}

func (c *Clock) GetCurrentTime() time.Time {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.currentTime
}

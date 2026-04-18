package dice

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Roll returns a random integer between 1 and sides (inclusive).
func Roll(sides int) int {
	if sides <= 0 {
		return 0
	}
	return rand.Intn(sides) + 1
}

// Chance returns true if a random roll (0.0 to 1.0) is less than threshold.
func Chance(threshold float64) bool {
	return rand.Float64() < threshold
}

// Factor returns a random multiplier around 1.0 based on variance.
// e.g., Factor(0.1) returns a value between 0.9 and 1.1.
func Factor(variance float64) float64 {
	return 1.0 + (rand.Float64()*2-1)*variance
}

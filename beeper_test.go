package beeper

import(
	"testing"
)

// Beep one time
func TestBeepOneTime (t *testing.T) {
	Beep()
}

// Beep three times
func TestBeepThreeTimes (t *testing.T) {
	Beep(3)
}

// Beep a melody
func TestBeepMelodicTimes (t *testing.T) {
	Melody("****-*-*")
}
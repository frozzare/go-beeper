package beeper

import(
	"testing"
	"github.com/bmizerany/assert"
)

// Beep one time
func TestBeepOneTime (t *testing.T) {
	assert.Equal(t, Beep(), true)
}

// Beep three times
func TestBeepThreeTimes (t *testing.T) {
	assert.Equal(t, Beep(3), true)
}

// Beep a melody
func TestBeepMelodicTimes (t *testing.T) {
	assert.Equal(t, Melody("****-*-*"), true)
}
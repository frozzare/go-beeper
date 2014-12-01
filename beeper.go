package beeper

import (
	"os"
	"time"
)

const delay = 100;

// Just beep!
func beep() {
	os.Stdout.Write([]byte("\u0007"))
}

// Beep a melody.
func Melody(params ...string, ) {
	if len(params) == 0 {
		return
	}
	
	for _, note := range params[0] {
		if note == 42 {
			beep()
		}
		
		time.Sleep(time.Second + delay)
	}
}

// Beep numbers of times.
func Beep(params ...int) {
	if (len(params) == 0) {
		params = make([]int, 1)
		params[0] = 1
	}
	
	for _, num := range params {
		for i := 0; i < num; i++ {
			beep()
			if i + 1 < num {
				time.Sleep(time.Second + delay)
			}
		}
	}
}
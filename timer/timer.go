package timer

import (
	"time"
	"fmt"
)

// Timer defined an struct to hold on to the start and finish times
type Timer struct {
	Start time.Time
	Finish time.Time
}

// StartTimer sets the start time of a Timer
func StartTimer(t * Timer, msg string) {
	if len(msg) > 0 {
		printMessage("START: " + msg)
	}

	t.Start = time.Now()
}

// StopTimer sets the finish time of a Timer
func StopTimer(t * Timer, msg string) {
	t.Finish = time.Now()

	if len(msg) > 0 {
		printMessage("END: " + msg)
	}
}

// ElapsedTime obtain a duration type from a Timer
// in case start or finish are not valid, zro is returned
func ElapsedTime(t * Timer) time.Duration {
	if t.Start.IsZero() || t.Finish.IsZero() {
		return time.Duration(0)
	}

	return t.Finish.Sub(t.Start)
}

// printMessage internal helper to print start or end timer
// to the screen
func printMessage(msg string) {
	msgLen := len(msg)

	if msgLen > 0 {
		for i := 0; i < msgLen + 4; i++ {
			fmt.Print("-")
		}

		fmt.Println("")
		fmt.Print("| ")
		fmt.Print(msg)
		fmt.Print(" |")
		fmt.Println("")

		for i := 0; i < msgLen + 4; i++ {
			fmt.Print("-")
		}

		fmt.Println("")
	}
}
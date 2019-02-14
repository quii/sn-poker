package poker

import (
	"fmt"
	"io"
	"time"
)

// BlindAlerter schedules alerts for blind amounts
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc allows you to implement BlindAlerter with a function
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

// Alerter will schedule alerts and print them to "to"
func Alerter(when time.Duration, amount int, to io.Writer) {
	time.AfterFunc(when, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})

	// not first alert
	if when > 0 {
		time.AfterFunc(when-30*time.Second, func() {
			fmt.Fprintf(to, "Blind will be increasing to %d shortly", amount)
		})
	}
}

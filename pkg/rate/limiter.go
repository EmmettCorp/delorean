/*
Package rate provides rate limiter functionality.
*/
package rate

import "time"

// Limiter is a client that helps you to manage rate limits.
type Limiter struct {
	expired bool
	dur     time.Duration
}

func NewLimiter(d time.Duration) *Limiter {
	return &Limiter{
		expired: true,
		dur:     d,
	}
}

// Allow checks if time == `Limiter.dur` passed from the last Allow call.
// If time passed it returns true, otherwise false.
func (l *Limiter) Allow() bool {
	if l.expired {
		l.expired = false
		go l.setNewTimer()

		return true
	}

	return false
}

func (l *Limiter) setNewTimer() {
	timer := time.NewTimer(l.dur)

	go func() {
		<-timer.C
		l.expired = true
	}()
}

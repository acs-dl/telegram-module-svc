package running

import "time"

type incrementalTimer struct {
	initialPeriod time.Duration
	maxPeriod     time.Duration
	multiplier    time.Duration

	currentPeriod time.Duration
	iteration     int
}

func newIncrementalTimer(initialPeriod, maxPeriod time.Duration, multiplier int) *incrementalTimer {
	return &incrementalTimer{
		initialPeriod: initialPeriod,
		maxPeriod:     maxPeriod,
		multiplier:    time.Duration(multiplier),

		currentPeriod: initialPeriod,
		iteration:     0,
	}
}

func (t *incrementalTimer) next() <-chan time.Time {
	result := time.After(t.currentPeriod)

	t.currentPeriod = t.currentPeriod * t.multiplier

	if t.currentPeriod > t.maxPeriod {
		t.currentPeriod = t.maxPeriod
	}

	t.iteration += 1

	return result
}


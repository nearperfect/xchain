package core

import "time"

type Ticker struct {
	Period time.Duration
	Ticker time.Ticker
}

func CreateTicker(period time.Duration) *Ticker {
	return &Ticker{period, *time.NewTicker(period)}
}

func (t *Ticker) ResetTicker() {
	t.Ticker = *time.NewTicker(t.Period)
}

func (t *Ticker) Stop() {
	t.Ticker.Stop()
}

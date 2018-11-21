package ws

import (
	"time"
)

const TFREQ = time.Second * 6

type Heart struct {
	cnt int
	// ticker *time.Ticker
	timer *time.Timer
}

func NewHeart() *Heart {
	return &Heart{
		0,
		time.NewTimer(TFREQ), //time.NewTicker(TFREQ),
	}
}

func (c *Heart) Reset() {
	c.cnt = 0
	// c.ticker = time.NewTicker(TFREQ)
	// c.timer.Reset(TFREQ)
}

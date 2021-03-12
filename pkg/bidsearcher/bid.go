package bidsearcher

import "time"

type Bid struct {
	unit     string
	caseName string
	vendor   string
	award    int
	date     time.Time
}

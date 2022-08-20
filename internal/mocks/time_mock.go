package mocks

import "time"

type Clock interface {
	Now() time.Time
}

type ClockMock struct{}

func (ClockMock) Now() time.Time {
	return time.Date(1970, time.January, 1, 1, 2, 3, 4, time.UTC)
}

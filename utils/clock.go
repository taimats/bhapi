package utils

import "time"

type Clock interface {
	Now() time.Time
}

type Clocker struct{}

func NewClocker() *Clocker {
	return &Clocker{}
}

func (*Clocker) Now() time.Time {
	return time.Now()
}

type TestClocker struct{}

func NewTestClocker() *TestClocker {
	return &TestClocker{}
}

func (*TestClocker) Now() time.Time {
	return time.Date(2024, 2, 5, 14, 43, 0, 0, time.UTC)
}

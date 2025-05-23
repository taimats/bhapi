package utils

import "time"

type Clock interface {
	Now() time.Time
}

type Clocker struct{}

func NewClocker() Clocker {
	return Clocker{}
}

//タイムゾーンがローカルに依存した現在時刻を返す
func (Clocker) Now() time.Time {
	return time.Now()
}

type TestClocker struct{}

func NewTestClocker() TestClocker {
	return TestClocker{}
}

//タイムゾーンがローカルに依存した固定の現在時刻を返す
func (TestClocker) Now() time.Time {
	return time.Date(2024, 2, 5, 14, 43, 0, 0, time.Local)
}

//タイムゾーンがローカルに依存した固定の現在時刻をRFC3339形式で返す
func (TestClocker) NowString() string {
	return time.Date(2024, 2, 5, 14, 43, 0, 0, time.Local).Format(time.RFC3339)
}

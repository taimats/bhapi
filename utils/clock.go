package utils

import "time"

var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

type Clock interface {
	Now() time.Time
}

type Clocker struct{}

func NewClocker() Clocker {
	return Clocker{}
}

//TZがJSTの現在時刻を返す
func (Clocker) Now() time.Time {
	return time.Now().In(JST)
}

type TestClocker struct{}

func NewTestClocker() TestClocker {
	return TestClocker{}
}

//TZがJSTで一定の現在時刻を返す
func (TestClocker) Now() time.Time {
	return time.Date(2024, 2, 5, 14, 43, 0, 0, JST)
}

//TZがJSTで一定の現在時刻をRFC3339形式で返す
func (TestClocker) NowString() string {
	return time.Date(2024, 2, 5, 14, 43, 0, 0, JST).Format(time.RFC3339)
}

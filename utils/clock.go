package utils

import "time"

type Clock interface {
	Now() time.Time
}

type Clocker struct{}

func NewClocker() Clocker {
	return Clocker{}
}

//TZがJSTの現在時刻を返す
func (Clocker) Now() time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Now().In(jst)
}

type TestClocker struct{}

func NewTestClocker() TestClocker {
	return TestClocker{}
}

//UTCの固定の現在時刻を返す
func (TestClocker) Now() time.Time {
	return time.Date(2024, 2, 5, 14, 43, 0, 0, time.UTC)
}

//UTCの固定の現在時刻をRFC3339形式で返す
func (TestClocker) NowString() string {
	return time.Date(2024, 2, 5, 14, 43, 0, 0, time.UTC).Format(time.RFC3339)
}

//TZがJSTの固定の現在時刻を返す
func (TestClocker) NowJST() time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Date(2024, 2, 5, 14, 43, 0, 0, jst)
}

//TZがJSTの固定の現在時刻をRFC3339形式で返す
func (TestClocker) NowJSTString() string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return time.Date(2024, 2, 5, 14, 43, 0, 0, jst).Format(time.RFC3339)
}

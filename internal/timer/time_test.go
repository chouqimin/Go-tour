package timer

import (
	"testing"
	"time"
)

type before struct {
	currentTimer time.Time
	d            string
}
type testData struct {
	before
	after time.Time
}

// 根據字串加減給定時間
func TestGetCalculateTime(t *testing.T) {
	tests := []testData{
		// testData{
		// 	before: before{currentTimer: time.Time{}, d: "-2h"},
		// 	after:  time.Time{}},
		testData{
			before: before{currentTimer: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), d: "-2h"},
			after:  time.Date(2009, 11, 17, 18, 34, 58, 651387237, time.UTC)},
		testData{
			before: before{currentTimer: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), d: ""},
			after:  time.Time{}},
	}
	for _, test := range tests {
		got, _ := GetCalculateTime(test.before.currentTimer, test.before.d)
		if got != test.after {
			t.Errorf("給的值為\"%v\"跟\"%s\"，想要的值為\"%v\", 但得到的值為\"%v\"", test.currentTimer, test.d, test.after, got)
		}
	}
}

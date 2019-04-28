package timer

import (
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
	counter := 0
	interval := SetInterval(func(args ...interface{}) {
		counter = counter + 1
	}, 300, &counter)

	time.AfterFunc(time.Duration(1)*time.Second, func() {
		if counter < 3 {
			t.Errorf("Expected interval to run 3 times within a second. Only ran %v", counter)
		}

		if counter != interval.Count() {
			t.Errorf("Expected %v iterations, only counted %v", counter, interval.Count())
		}
	})
}

func TestTimeout(t *testing.T) {
	counter := 0
	interval := SetTimeout(func(args ...interface{}) {
		counter = counter + 1
	}, 300, &counter)

	time.AfterFunc(time.Duration(1)*time.Second, func() {
		if counter > 1 {
			t.Errorf("Expected timeout to run once within a second. # Times Run: %v", counter)
		}

		if counter != interval.Count() {
			t.Errorf("Expected %v iterations, only counted %v", counter, interval.Count())
		}
	})
}

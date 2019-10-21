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

	timer := time.AfterFunc(time.Duration(1)*time.Second, func() {
		if counter < 3 {
			t.Errorf("Expected interval to run 3 times within a second. Only ran %v", counter)
		}

		if counter != interval.Count() {
			t.Errorf("Expected %v iterations, only counted %v", counter, interval.Count())
		}

		interval.Stop()
	})

	defer timer.Stop()

	time.Sleep(time.Duration(1400) * time.Millisecond)
}

func TestTimeout(t *testing.T) {
	counter := 0
	interval := SetTimeout(func(args ...interface{}) {
		counter = counter + 1
	}, 300, &counter)

	timer := time.AfterFunc(time.Duration(1)*time.Second, func() {
		if counter > 1 {
			t.Errorf("Expected timeout to run once within a second. # Times Run: %v", counter)
		}

		if counter != interval.Count() {
			t.Errorf("Expected %v iterations, only counted %v", counter, interval.Count())
		}
	})

	defer timer.Stop()

	time.Sleep(time.Duration(500) * time.Millisecond)
}

func TestMultipleTimers(t *testing.T) {
	counter := 0
	var counterFn func(args ...interface{})

	counterFn = func(args ...interface{}) {
		counter = counter + 1

		if counter < 3 {
			_ = SetTimeout(counterFn, 300)
		}
	}

	_ = SetTimeout(counterFn, 300)

	timer := time.AfterFunc(time.Duration(2)*time.Second, func() {
		if counter != 3 {
			t.Errorf("Expected timeout to run 3 times within a second. # Times Run: %v", counter)
			t.Fail()
		}
	})

	defer timer.Stop()

	time.Sleep(time.Duration(2500) * time.Millisecond)
}

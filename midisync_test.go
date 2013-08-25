package main

import (
	"testing"
	"time"
)

func TestMicrosecondsPerPulse(t *testing.T) {
	res := microsecondsPerPulse(120.0)
	if res != time.Duration(20833333) {
		t.Error("Unexpected Microscend interval for 120bpm:", res)
	}
}

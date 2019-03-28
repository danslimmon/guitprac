package main

import (
	"testing"
)

func TestDistance(t *testing.T) {
	majorScale := MajorScale{}
	majorScaleDegrees := majorScale.Degrees()

	// a perfect 4th is 5 half steps up from the root
	func() {
		expected := IntervalByHalfsteps(5)
		got := majorScale.DistanceUp(majorScaleDegrees[0], majorScaleDegrees[3])
		if expected != got {
			t.Logf("expected a perfect 4th to be %d half steps up from root; got %d", expected.Halfsteps, got.Halfsteps)
			t.Fail()
		}
	}()

	// a perfect 4th is 7 half steps down from the root
	func() {
		expected := IntervalByHalfsteps(7)
		got := majorScale.DistanceDown(majorScaleDegrees[0], majorScaleDegrees[3])
		if expected != got {
			t.Logf("expected a perfect 4th to be %d half steps down from root; got %d", expected.Halfsteps, got.Halfsteps)
			t.Fail()
		}
	}()

	// an octave is 12 half steps up from the root
	func() {
		expected := IntervalByHalfsteps(12)
		got := majorScale.DistanceUp(majorScaleDegrees[0], majorScaleDegrees[0])
		if expected != got {
			t.Logf("expected an octave to be %d half steps up from root; got %d", expected.Halfsteps, got.Halfsteps)
			t.Fail()
		}
	}()

	// an octave is 12 half steps down from the root
	func() {
		expected := IntervalByHalfsteps(12)
		got := majorScale.DistanceDown(majorScaleDegrees[0], majorScaleDegrees[0])
		if expected != got {
			t.Logf("expected an octave to be %d half steps down from root; got %d", expected.Halfsteps, got.Halfsteps)
			t.Fail()
		}
	}()
}

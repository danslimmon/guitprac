package main

import (
	"bytes"
	"testing"
)

func TestDegreeStepsFormatter(t *testing.T) {
	majorScale := MajorScale{}
	majorScaleDegrees := majorScale.Degrees()

	func() {
		fmtr := &DegreeStepsFormatter{}
		fmtr.Start(majorScaleDegrees[0])
		fmtr.AddStep("up", majorScaleDegrees[5])
		fmtr.AddStep("down", majorScaleDegrees[0])

		var outBuf bytes.Buffer
		got := make([]byte, 0)
		fmtr.Flush(&outBuf)
		got = outBuf.Bytes()
		expected := []byte(" 6\n1 1\n")
		if !bytes.Equal(expected, got) {
			t.Logf("DegreeStepsFormatter failed to produce correct output: expected %v; got %v", expected, got)
			t.Fail()
		}
	}()
}

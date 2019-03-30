package main

import (
	"io"
	"strings"
)

type degreeStepsFormatterStep struct {
	Direction string
	Degree    ScaleDegree
}

type DegreeStepsFormatter struct {
	steps []degreeStepsFormatterStep
}

func (fmtr *DegreeStepsFormatter) Start(deg ScaleDegree) {
	fmtr.steps = []degreeStepsFormatterStep{{"", deg}}
}

func (fmtr *DegreeStepsFormatter) AddStep(direction string, degree ScaleDegree) {
	fmtr.steps = append(fmtr.steps, degreeStepsFormatterStep{direction, degree})
}

func (fmtr *DegreeStepsFormatter) Flush(output io.Writer) {
	rsltRows := fmtr.getRows()
	rslt := strings.Join(rsltRows, "\n")
	rslt = strings.Join([]string{rslt, "\n"}, "")
	_, _ = output.Write([]byte(rslt))
}

func (fmtr *DegreeStepsFormatter) getRows() []string {
	if len(fmtr.steps) == 0 {
		return []string{}
	}

	rows := make([]string, 1)
	rows[0] = fmtr.steps[0].Degree.Name
	width := len(rows[0])
	lastRowInd := 0

	for _, step := range fmtr.steps[1:] {
		if step.Direction == "up" {
			if lastRowInd == 0 {
				// add a new row, since we're going up from the highest row
				rows = append([]string{""}, rows...)
				lastRowInd++
			}
			newRowInd := lastRowInd - 1
			rows[newRowInd] = fmtr.pad(rows[newRowInd], width)
			rows[newRowInd] = strings.Join([]string{rows[newRowInd], step.Degree.Name}, "")
			lastRowInd = newRowInd
		} else {
			if lastRowInd == len(rows)-1 {
				// add a new row, since we're going down from the lowest row
				rows = append(rows, "")
			}
			newRowInd := lastRowInd + 1
			rows[newRowInd] = fmtr.pad(rows[newRowInd], width)
			rows[newRowInd] = strings.Join([]string{rows[newRowInd], step.Degree.Name}, "")
			lastRowInd = newRowInd
		}
		width = len(rows[lastRowInd])
	}
	return rows
}

func (fmtr *DegreeStepsFormatter) pad(s string, w int) string {
	if w < len(s) {
		return s
	}
	spaces := strings.Repeat(" ", w-len(s))
	return strings.Join([]string{s, spaces}, "")
}

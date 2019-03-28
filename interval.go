package main

import (
	"fmt"
)

type Interval struct {
	Name      string
	Halfsteps int
}

func IntervalByHalfsteps(halfsteps int) Interval {
	switch halfsteps {
	case 0:
		return Interval{"unison", 0}
	case 1:
		return Interval{"half step", 1}
	case 2:
		return Interval{"whole step", 2}
	case 3:
		return Interval{"minor 3rd", 3}
	case 4:
		return Interval{"major 3rd", 4}
	case 5:
		return Interval{"perfect 4th", 5}
	case 6:
		return Interval{"diminished 5th", 6}
	case 7:
		return Interval{"perfect 5th", 7}
	case 8:
		return Interval{"minor 6th", 8}
	case 9:
		return Interval{"major 6th", 9}
	case 10:
		return Interval{"minor 7th", 10}
	case 11:
		return Interval{"major 7th", 11}
	case 12:
		return Interval{"octave", 12}
	default:
		panic(fmt.Sprintf("no interval with %d half steps", halfsteps))
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	type position struct {
		Halfsteps int
		Degree    ScaleDegree
	}

	bottomLimit := -15
	topLimit := 24
	majorScale := MajorScale{}
	majorScaleDegrees := majorScale.Degrees()
	pos := position{0, majorScaleDegrees[0]}

	rand.Seed(time.Now().UTC().UnixNano())

	degDir := func() (ScaleDegree, int) {
		// pick a random degree of the scale
		degree := majorScaleDegrees[rand.Int()%len(majorScaleDegrees)]
		// pick a random direction to move (either 1 or -1)
		direction := -1 + (2 * (rand.Int() % 2))
		return degree, direction
	}

	fmt.Println("start one octave above the lowest G")
	for i := 0; i < 32; i++ {
		newDegree, direction := degDir()
		var newPos position
		var moveInterval Interval
		var directionName string

		if direction == 1 {
			directionName = "UP"
			moveInterval = majorScale.DistanceUp(pos.Degree, newDegree)
			newPos = position{pos.Halfsteps + moveInterval.Halfsteps, newDegree}
		} else {
			directionName = "DOWN"
			moveInterval = majorScale.DistanceDown(pos.Degree, newDegree)
			newPos = position{pos.Halfsteps - moveInterval.Halfsteps, newDegree}
		}

		if newPos.Halfsteps < bottomLimit || newPos.Halfsteps > topLimit {
			i--
			continue
		}
		pos = newPos
		fmt.Printf("%-4s to the nearest %-4s (comes to %2d half-steps up from original note)\n", directionName, pos.Degree.Name, pos.Halfsteps)
		if i%4 == 3 && i < 31 {
			fmt.Println()
		}
	}
}

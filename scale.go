package main

type ScaleDegree struct {
	Name     string
	Interval Interval
}

type Scale interface {
	Degrees() []ScaleDegree
	DistanceUp(from, to ScaleDegree) Interval
	DistanceDown(from, to ScaleDegree) Interval
}

type MajorScale struct{}

func (scale MajorScale) Degrees() []ScaleDegree {
	return []ScaleDegree{
		{"1", IntervalByHalfsteps(0)},
		{"2", IntervalByHalfsteps(2)},
		{"3", IntervalByHalfsteps(4)},
		{"4", IntervalByHalfsteps(5)},
		{"5", IntervalByHalfsteps(7)},
		{"6", IntervalByHalfsteps(9)},
		{"7", IntervalByHalfsteps(11)},
	}
}

func (scale MajorScale) DistanceUp(from, to ScaleDegree) Interval {
	if from.Interval.Halfsteps < to.Interval.Halfsteps {
		return IntervalByHalfsteps(to.Interval.Halfsteps - from.Interval.Halfsteps)
	} else if from.Interval.Halfsteps > to.Interval.Halfsteps {
		return IntervalByHalfsteps(12 - (from.Interval.Halfsteps - to.Interval.Halfsteps))
	}
	// octave
	return IntervalByHalfsteps(12)
}

func (scale MajorScale) DistanceDown(from, to ScaleDegree) Interval {
	if from.Interval.Halfsteps > to.Interval.Halfsteps {
		return IntervalByHalfsteps(from.Interval.Halfsteps - to.Interval.Halfsteps)
	} else if from.Interval.Halfsteps < to.Interval.Halfsteps {
		return IntervalByHalfsteps(12 - (to.Interval.Halfsteps - from.Interval.Halfsteps))
	}
	// octave
	return IntervalByHalfsteps(12)
}

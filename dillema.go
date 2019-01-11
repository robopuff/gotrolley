package main

import (
	"fmt"
	"math"
	"math/rand"
)

type dillema interface {
	printDillema()
	printDecision(move bool)
	updatePoints(move bool) points
}

type trolleyDillema struct {
	out            output
	lowerTrack     int
	upperTrack     int
	lowerTrackText string
	upperTrackText string
}

func (d trolleyDillema) printDillema() {
	d.out.printf("A runaway trolley is barrelling towards %s\n", d.lowerTrackText)
	d.out.printf("You can pull a lever to divert the trolley to another track, containing %s\n", d.upperTrackText)

}

func (d trolleyDillema) printDecision(move bool) {
	if move {
		d.out.println("You moved a lever")
	} else {
		d.out.println("You left lever where it was")
	}
}

func (d trolleyDillema) updatePoints(move bool) points {
	pts := points{0, 0}
	diff := d.lowerTrack - d.upperTrack

	abs := func(in int) int {
		return int(math.Abs(float64(in)))
	}

	if (diff < 0 && move) || (diff > 0 && !move) {
		d.out.printf("You have made the wrong utilitarian decision. Lose %d utils!\n", abs(diff))
		pts.utils = -1 * abs(diff)
	} else if diff == 0 {
		d.out.println("Your choice is neutral on utilitarian grounds")
	} else if (diff > 0 && move) || (diff < 0 && !move) {
		d.out.printf("You have made the correct utilitarian decision and saved a net %d lives.\n", abs(diff))
		d.out.printf("Gain %d utils!\n", abs(diff))
		pts.utils = abs(diff)
	}

	if move && diff > 0 {
		d.out.println("You have a hypothetical imperative to save lives, but not a categorical one.")
		d.out.println("Gain 1 Kant point")
		pts.kant = 1
	} else {
		d.out.println("You are a MURDERER who has violated the categorical imperative!")
		d.out.println("Lose 10 Kant points.")
		pts.kant = -10
	}

	return pts
}

func newTrolleyDillema(out output) trolleyDillema {
	lower := randInt(1, 10)
	upper := randInt(0, 5)

	var lowerText, upperText string

	if lower == 1 {
		lowerText = "one worker who is mysteriously tied up. "
	} else {
		lowerText = fmt.Sprintf("%d workers who are mysteriously tied up. ", lower)
	}

	if upper == 1 {
		upperText = "one worker who is also tied up."
	} else {
		upperText = fmt.Sprintf("%d workers who are also tied up. ", upper)
	}

	return trolleyDillema{
		out:            out,
		lowerTrack:     lower,
		upperTrack:     upper,
		lowerTrackText: lowerText,
		upperTrackText: upperText,
	}
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

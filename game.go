package main

import (
	"math/rand"
	"os"
	"time"
)

type points struct {
	kant  int
	utils int
}

func (p *points) add(pts points) {
	p.kant += pts.kant
	p.utils += pts.utils
}

type game struct {
	out    output
	count  int
	points *points
}

func (g *game) play() {
	g.count++
	d := g.createDillema()
	d.printDillema()
	time.Sleep(400 * time.Millisecond)

	move := g.out.pullLever()
	d.printDecision(move)
	g.points.add(d.updatePoints(move))

	g.showResults()
	g.continuePlaying()
	time.Sleep(200 * time.Millisecond)
	g.play()
}

func (g game) continuePlaying() {
	g.out.print("Continue? [Y/N] ")
	if !g.out.booleanInput() {
		g.out.println("Expect the unexpected\n - Trolleys for life (or death)")
		os.Exit(0)
	}
}

func (g game) createDillema() dillema {
	return newTrolleyDillema(g.out)
}

func (g game) showResults() {
	g.out.printf("You've encoutered %d ethical dillemas\n", g.count)
	g.out.printf("Your points\n %d kant pts\n %d util pts\n", g.points.kant, g.points.utils)
}

func newGame(out output) *game {
	pts := &points{0, 0}
	return &game{
		out:    out,
		count:  0,
		points: pts,
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	out := newTerminal()
	out.println("Go, Trolley Game")
	g := newGame(out)
	g.showResults()
	g.play()
}

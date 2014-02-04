package main

import (
	"flag"
	"fmt"
	eg "github.com/trumae/evolvegamelib"
	"math/rand"
	"time"
)

var numCycles int
var seed int64

func main() {
	fmt.Print("EvolveGame\n\n")
	flag.IntVar(&numCycles, "numCycles", 100, "Number of evolution cycles (default: 100)")
	flag.Int64Var(&seed, "seed", time.Now().UnixNano(), "seed for random number (default: now)")
	flag.Parse()
	rand.Seed(seed)
	a := eg.NewArenaSample()
	a.Evolution(numCycles)
	a.DrawWorld()
}

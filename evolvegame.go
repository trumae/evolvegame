package main

import (
 "fmt"
 "flag"
 "math/rand"
 "time"
 eg "github.com/trumae/evolvegamelib"
)

var numCycles int

func main() {
        fmt.Print("EvolveGame\n\n")
	flag.IntVar(&numCycles, "numCycles", 100, "Number of evolution cycles (default: 100)")
	flag.Parse()
        rand.Seed(time.Now().UnixNano())
	a := eg.NewArenaSample()
	a.Evolution(numCycles)
        a.DrawWorld()
}


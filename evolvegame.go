package main

import "fmt"
import "math/rand"
import "time"
import eg "github.com/trumae/evolvegamelib"

func main() {
	fmt.Print("Hello 2\n")
        rand.Seed(time.Now().UnixNano())
	a := eg.NewArenaSample()
	a.Evolution(200)
        a.DrawWorld()
}


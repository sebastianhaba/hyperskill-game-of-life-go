package main

import (
	"flag"
	"github.com/inancgumus/screen"
	"github.com/sebastianhaba/hyperskiill-game-of-life-go/universe"
	"time"
)

func main() {
	size := flag.Int("size", 10, "Size of the universe")
	seed := flag.Int64("seed", 0, "Seed for the random number generator")
	generations := flag.Int("generations", 10, "Number of generations to run")
	pauseTime := flag.Int("pause", 500, "Pause time between generations in milliseconds")
	flag.Parse()

	if *size <= 0 {
		*size = 4
	}

	if *seed <= 0 {
		*seed = time.Now().Unix()
	}

	if *generations <= 0 {
		*generations = 10
	}

	screen.Clear()

	u := universe.NewUniverse(*size, *seed)

	for i := 0; i < *generations; i++ {
		screen.MoveTopLeft()
		u.Display()
		u.NextGeneration()
		time.Sleep(time.Duration(*pauseTime) * time.Millisecond)
	}
}

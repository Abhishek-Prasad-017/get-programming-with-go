package unit_7

import (
	"fmt"
	"image"
	"time"
)

/*
Using listing 31.5 as a starting point, change the code so that the delay time gets a half a second longer with each move.
 */

func PositionWorker() {
	go Worker()
	time.Sleep(3 * time.Second)
}

func Worker() {
	delay := time.Second

	pos := image.Point{
		X: 5, Y: 5,
	}
	direction := image.Point{
		X: 2, Y: 0,
	}

	next := time.After(delay)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("Current position is ", pos)
			delay = delay + time.Second / 2
			next = time.After(delay)
		}
	}
}

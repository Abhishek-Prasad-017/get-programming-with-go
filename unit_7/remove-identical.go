package unit_7

/*
It’s boring to see the same line repeated over and over again.
Write a pipeline element (a goroutine) that remembers the previous value and only sends the value to the next stage of the pipeline if it’s different from the one that came before.
To make things a little simpler, you may assume that the first value is never the empty string.
*/

import (
	"fmt"
)

func RemoveIdentical() {

	c0 := make(chan string)
	c1 := make(chan string)
	go StartSource1(c0)
	go RemDuplicates(c0, c1)
	Display1(c1)
}

func StartSource1(downstream chan string) {
	s := []string{"a", "v", "v", "p", "p", "e", "m", "e"}
	for _, v := range s {
		downstream <- v
	}
	close(downstream)
}

func Display1(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func RemDuplicates(upstream, downstream chan string) {
	prev := ""
	for v := range upstream {
		if v != prev {
		downstream <- v
		prev = v
		}
	}
	close(downstream)
}



package unit_7

/*
Sometimes it’s easier to operate on words than on sentences.
Write a pipeline element that takes strings, splits them up into words (you can use the Fields function from the strings package), and sends all the words, one by one, to the next pipeline stage.
 */

import (
	"fmt"
	"strings" )

func Split() {

	c0 := make(chan string)
	c1 := make(chan string)
	go StartSource(c0)
	go SplitWords(c0, c1)
	Display(c1)
}

func StartSource(downstream chan string) {
	s := []string{"Bye", "Apple is red", "Application"}
	for _, v := range s{
		downstream <- v
	}
		close(downstream)
}

func Display(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}


func SplitWords(upstream, downstream chan string) {
	for v := range upstream {
		for _, word := range strings.Fields(v) {
			downstream <- word
		}
	}
		close(downstream)
}

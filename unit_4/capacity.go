package unit_4

/*
Write a program that uses a loop to continuously append an element to a slice.
Print out the capacity of the slice whenever it changes. Does append always double the capacity when the underlying array runs out of room?
*/

import (
	"fmt"
	"math/rand"
)

func AppendRandomNum(s []int) {
	capTest := cap(s)
	for i := 0; i <= 20; i++ {
		s = append(s, rand.Intn(50))
		if cap(s) != capTest {
			fmt.Println("Capacity is ", cap(s))
			capTest = cap(s)
		}
	}

}

func Appending() {
	slice := make([]int, 0, 10)
	AppendRandomNum(slice)
}

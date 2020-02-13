package unit_6
/*
Write a program with a turtle that can move up, down, left, or right. The turtle should store an (x, y) location where positive values go down and to the right.
Use methods to increment/decrement the appropriate variable. A main function should exercise the methods you’ve written and print the final location.
 */

import (
	"fmt"
)

func GoUp(x *int, y *int) {
	*y = *y + 5
}

func GoLeft(x *int, y *int) {
	*x = *x - 5
}

func GoRight(x *int, y *int) {
	*x = *x + 5
}

func GoDown(x *int, y *int) {
	*y = *y - 5
}

func StartTurtle(x *int, y *int) {
	GoUp(x,y)
	GoLeft(x,y)
	GoRight(x,y)
	GoDown(x,y)
}

func Turtle() {
	x := 5
	y := 5
	//fmt.Printf("Addres 1 = %v", &x)
	StartTurtle(&x,&y)
	fmt.Printf("Co - ordinates of the turtle now is %v %v \n",x,y)
}
package unit_6

import (
	"fmt"
)

/*
Write a program with a turtle that can move up, down, left, or right. The turtle should store an (x, y) location where positive values go down and to the right.
Use methods to increment/decrement the appropriate variable. A main function should exercise the methods you’ve written and print the final location.
*/


type turtle struct{
	x int
	y int
}

func (t *turtle) GoUp() {
	t.y = t.y + 5
}

func (t *turtle) GoLeft() {
	t.x = t.x - 5
}

func (t *turtle) GoRight() {
	t.x = t.x + 5
}

func (t *turtle) GoDown() {
	t.y = t.y - 5
}


func Turtle2() {
	t := &turtle{
		5,5,
	}
	//t := turt
	//fmt.Printf("Addres 1 = %v", &x)
	//StartTurtle(t)
	t.GoUp()
	t.GoLeft()
	t.GoRight()
	t.GoDown()
	fmt.Printf("Co - ordinates of the turtle now is %v %v \n",t.x,t.y)
}
package unit_1

/*
Write a guess-the-number program. Make the computer pick random numbers between 1–100 until it guesses your number, which you declare at the top of the program.
Dis- play each guess and whether it was too big or too small.
*/

import (
	"fmt"
	"math/rand"
)

func GuessNumber() {
	var myNumber = 20
	var num = rand.Intn(100) + 1
	for num != myNumber {
		fmt.Println("Number guessed is ", num)
		if num > myNumber {
			fmt.Println("The number guessed was bigger!!!")
		} else {
			fmt.Println("The number guessed was smaller!!!")
		}
		num = rand.Intn(100) + 1
	}
	fmt.Println("Your Computer Won")

}

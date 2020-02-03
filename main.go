package main

//go mod init github.com/get-programming-with-go

import (
	u1 "github.com/get-programming-with-go/unit_1"
	u2 "github.com/get-programming-with-go/unit_2"
)

func main() {
	u1.PrintName("Abhishek Prasad")
	u1.PrintNameInHindi("अभिषेक प्रसाद")
	u1.CalculateSpeed()
	u1.Health()
	u1.GuessNumber()
	u1.LeapYear()
	u2.Piggy()
	u2.Piggy2()
}

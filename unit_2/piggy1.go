package unit_2

import (
	"fmt"
	"math/rand"
)

/*
Save some money to buy a gift for your friend.
Write a program that randomly places nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it con- tains at least $20.00.
 Display the running balance of the piggy bank after each deposit, formatting it with an appropriate width and precision.
*/

func Piggy() {
	var nickles = 0.05
	var dimes = 0.10
	var quarters = 0.25
	var amt = 0.0
	for amt <= 20 {
		switch choice := rand.Intn(3); choice {
		case 0:
			amt = amt + nickles
		case 1:
			amt = amt + dimes
		case 2:
			amt = amt + quarters
		}
		fmt.Printf("Amt is now %.2f\n", amt)
	}
	//fmt.Printf("Amt is now %.2f\n", amt)
}

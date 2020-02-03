//Program to show usage of if condition
//The program shows your health accordig to your BMI

package unit_1

import "fmt"

func health() {
	var bmi = 23.3
	//salary = salary - 500000
	if bmi <= 25 {
		fmt.Println("You are Fit!!!")
	} else {
		fmt.Println("You need to exercise more!!!")
	}
}

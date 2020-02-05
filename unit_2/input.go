package unit_2

/*
Write a program that converts strings to Booleans:
 The strings “true”, “yes”, or “1” are true.
 The strings “false”, “no”, or “0” are false.
 Display an error message for any other values.
*/

import (
	"fmt"
)

func StringsToBooleans() {
	value := "1"
	var boolValue bool
	switch value {
	case "true", "yes", "1":
		boolValue = true
		fmt.Println(boolValue)
	case "false", "no", "0":
		boolValue = false
		fmt.Println(boolValue)
	default:
		fmt.Println("Error Occured!!!")
	}

	//fmt.Println(boolValue)

}

/*
Modify listing 4.7 to handle leap years:
 Generate a random year instead of always using 2018.
 For February, assign daysInMonth to 29 for leap years and 28 for other years.
 Use a for loop to generate and display 10 random dates.
*/

package unit_1

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func leapYear() {
	year := rand.Intn(2020) + 1
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		if year%4 == 0 {
			if year%100 == 0 {
				if year%400 == 0 {
					daysInMonth = 29
				} else {
					daysInMonth = 28
				}
			} else {
				daysInMonth = 29
			}
		} else {
			daysInMonth = 28
		}
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	for count := 10; count > 0; count-- {
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}

}

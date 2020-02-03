package unit_2

import (
	"fmt"
)

/*
Canis Major Dwarf is the closest known galaxy to Earth at 236,000,000,000,000,000 km from our Sun (though some dispute that it is a galaxy).
Use constants to convert this dis-tance to light years.
*/

func LightYears() {
	const distance = 236000000000000000
	const light = 9.461e+15
	fmt.Println("No. of light years  =  ", distance/light, "\n")
}

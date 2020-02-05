package unit_3

/*
Type listing 14.6 into the Go Playground to see it in action:
 Rather than passing 5 as an argument to calibrate, declare and pass a variable. Modify the variable and you’ll notice that calls to sensor() still result in 5.
That’s because the offset parameter is a copy of the argument (pass by value).
 Use calibrate with the fakeSensor function from listing 14.2 to create a new sensor function.
Call the new sensor function multiple times and notice that the original fakeSensor is still being called each time, resulting in random values.
*/

import (
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

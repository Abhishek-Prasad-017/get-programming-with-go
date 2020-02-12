package unit_3

/*
Use the Go Playground at play.golang.org to type in listing 12.1 and declare additional temperature conversion functions:
 Reuse the kelvinToCelsius function to convert 233 K to Celsius.
 Write and use a celsiusToFahrenheit temperature conversion function. Hint: the for-
mula for converting to Fahrenheit is: (c * 9.0 / 5.0) + 32.0.
 Write a kelvinToFahrenheit function and verify that it converts 0 K to approximately –459.67° F.
Did you use kelvinToCelsius and celsiusToFahrenheit in your new function or write an independent function with a new formula? Both approaches are perfectly valid.
 */

import(
	"fmt"
)

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func TemperatureConversion(){
	fmt.Printf("0 degree K is %.2f degree F\n", kelvinToFahrenheit(0))
	fmt.Printf("233 degree K is %.2f degree C\n", kelvinToCelsius(233))
}

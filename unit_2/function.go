package unit_2

/*
Use the Go Playground at play.golang.org to type in listing 12.1 and declare additional temperature conversion functions:
 Reuse the kelvinToCelsius function to convert 233 K to Celsius.
 Write and use a celsiusToFahrenheit temperature conversion function. Hint: the for-
mula for converting to Fahrenheit is: (c * 9.0 / 5.0) + 32.0.
 Write a kelvinToFahrenheit function and verify that it converts 0 K to approxi-
mately –459.67° F.
*/

func KelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func CelsiusToFahrenheit(c float64) float64 {
	fahrenheit := (c * 9.0 / 5.0) + 32.0
	return fahrenheit
}

func KelvinToFahrenheit(k float64) float64 {
	fahrenheit := (k - 459.67)
	return fahrenheit
}

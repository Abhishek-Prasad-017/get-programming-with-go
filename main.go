package main

//go mod init github.com/get-programming-with-go

import (
	u1 "github.com/get-programming-with-go/unit_1"
	u2 "github.com/get-programming-with-go/unit_2"
	//u3 "github.com/get-programming-with-go/unit_3"
	//u4 "github.com/get-programming-with-go/unit_4"
)

type celsius float64
type fahrenheit float64
type kelvin float64

func main() {

	u1.PrintName("Abhishek Prasad")
	/*
		u1.PrintNameInHindi("अभिषेक प्रसाद")
		u1.CalculateSpeed()
		u1.Health()
		u1.GuessNumber()
		u1.LeapYear()
		u2.Piggy()
		u2.Piggy2()
	*/
	//u2.CeaserCipher("L fdph, L vdz, L frqtxhuhg.")
	//u2.Internatioanl("Hola Estación Espacial Internacional")
	u2.StringsToBooleans()
	//kelvin := 294.0
	//cel := 24.0
	//celsius := u2.KelvinToCelsius(kelvin)
	//fahrenheit := u2.CelsiusToFahrenheit(cel)
	//fah := u2.KelvinToFahrenheit(kelvin)
	//fmt.Print(kelvin, " o K is ", celsius, "o C", "\n")
	//fmt.Print(cel, " o C is ", fahrenheit, "o F", "\n")
	//fmt.Print(kelvin, " o K is ", fah, "o F", "\n")
	//var c celsius
	//var f fahrenheit
	//c = f.Celcius()
	//u3.method()
	//u4.CallChess()
	//u4.Appending()
}

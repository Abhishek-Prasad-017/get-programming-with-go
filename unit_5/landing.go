package unit_5

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program that displays the JSON encoding of the three rover landing sites in list- ing 21.8.
The JSON should include the name of each landing site and use struct tags as shown in listing 21.10.
To make the output friendlier, make use of the MarshalIndent function from the json package.
*/
type location struct {
	Name string  `json:"name"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func Json() {
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	bytes, err := json.MarshalIndent(locations, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}

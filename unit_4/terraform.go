package unit_4

import (
	"fmt"
)

/*
Write a program to terraform a slice of strings by prepending each planet with "New ". Use your program to terraform Mars, Uranus, and Neptune.
Your first iteration can use a terraform function, but your final implementation should introduce a Planets type with a terraform method, similar to sort.StringSlice.
*/

type Planets []string

func (planets Planets) Terra() {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}

func TransformPlanets() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
	}
	Planets(planets[3:4]).Terra()
	Planets(planets[6:7]).Terra()
	fmt.Println(planets)
}

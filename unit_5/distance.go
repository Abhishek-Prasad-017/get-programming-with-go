package unit_5
import (
	"fmt"
)

func reclassify(planets []string) {
	//fmt.Printf("Old Address  = %p",planets[0])
	//delete(planets,"Mars")
	planets[0] = "A"
	//fmt.Printf("Old Address  = %p",planets)
}
func CallMe() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto",
	}
	//fmt.Printf("Old Address  = %p",planets)
	reclassify(planets)
	fmt.Println(planets)
}


package unit_5
/*

Write a program with a gps structure for a Global Positioning System (GPS). This struct should be composed of a current location, destination location, and a world.
Implement a description method for the location type that returns a string containing the name, latitude, and longitude.
The world type should implement a distance method using the math from lesson 22.
Attach two methods to the gps type. First, attach a distance method that finds the distance between the current and destination locations.
Then implement a message method that returns a string describing how many kilometers remain to the destination.
As a final step, create a rover structure that embeds the gps and write a main function to test everything out.
Initialize a GPS for Mars with a current location of Bradbury Land- ing (-4.5895, 137.4417) and a destination of Elysium Planitia (4.5, 135.9).
Then create a curiosity rover and print out its message (which forwards to the gps).

 */

import (
	"fmt"
	"math"
)
 type world struct {
 	radius float64
 }

 type location1 struct {
 	name string
 	lat , long float64
 }

 type gps struct {
 	world world
 	currLocation location1
 	destLocation location1

 }

func (loc location1) description() string {
	return fmt.Sprintf("%v (%f %f)", loc.name, loc.lat, loc.long)
}

func (w world) distance(a1 , a2 location1) float64 {
	s1, c1 := math.Sincos(rad(a1.lat))
	s2, c2 := math.Sincos(rad(a2.lat))
	clong := math.Cos(rad(a1.long - a2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (g gps) distance() float64 {
	return g.world.distance(g.currLocation , g.destLocation)
}

func rad(deg float64) float64 { return deg * math.Pi / 180
}

func (g gps) message() string {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.destLocation.description())
}

type rover struct {
	gps
}

func UseGps() {
	mars := world{radius: 3389.5}
	bradbury := location1{"Bradbury Landing", -4.5895, 137.4417}
	elysium := location1{"Elysium Planitia", 4.5, 135.9}
	gps := gps{
		world:        mars,
		currLocation: bradbury,
		destLocation: elysium,
	}
	curiosity := rover {
		gps: gps,
	}
	fmt.Println(curiosity.message())
}



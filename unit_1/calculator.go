package unit_1

/*
Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis.
Write a program to determine how fast a ship would need to travel (in km/h) in order to reach Malacan- dra in 28 days. Assume a distance of 56,000,000 km.
*/

import "fmt"

// calculates-the-distance
func calculateSpeed() {
	const days = 28
	const hoursInDay = 24
	var dist = 56000000
	var speed = dist / (days * hoursInDay)
	fmt.Printf("Speed of the ship should be %-2v Km/Hr \n", speed)
}

package unit_1

/*
Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis.
Write a program to determine how fast a ship would need to travel (in km/h) in order to reach Malacan- dra in 28 days. Assume a distance of 56,000,000 km.
*/

import "fmt"

// calculates-the-distance
func CalculateSpeed() {
	const days = 28
	var dist = 56000000
	var speed = dist / (days * 24)
	fmt.Printf("Speed of the ship should be %-2v Km/Hr ", speed)
}

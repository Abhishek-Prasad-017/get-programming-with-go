package unit_1

/*
The table should have four columns:
 The spaceline company providing the service
 The duration in days for the trip to Mars (one-way)  Whether the price covers a return trip
 The price in millions of dollars
For each ticket, randomly select one of the following spacelines: Space Adventures, SpaceX, or Virgin Galactic.
Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km away from Earth at the time.
Randomly choose the speed the ship will travel, from 16 to 30 km/s.
This will determine the duration for the trip to Mars and also the ticket price.
Make faster ships more expen- sive, ranging in price from $36 million to $50 million. Double the price for round trips.
 */

import (
	"fmt"
	"math/rand"
	)


func GetTickets() {
	const secPerDay = 86400
	distance := 62100000
	company := ""
	trip := ""
	fmt.Printf("Hello Future Traveller!!!\n")
	fmt.Println("Spaceline Days Trip type Price")
	fmt.Println("Spaceline ======================================")
	for c := 0; c < 10; c++ {
		switch choice := rand.Intn(3);choice {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}
	}
	speed := rand.Intn(15) + 16
	price := 20.0 + speed
	duration := distance / speed / secPerDay
	// For round/single trip
	if rand.Intn(2) == 1 {
		trip = "Round-trip"
		price = price * 2
	} else {
		trip = "One-way"
	}
	fmt.Printf("%-15v %2v %-15v $%2v\n", company, duration, trip, price)
}



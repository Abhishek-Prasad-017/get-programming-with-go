package unit_5
/*
Your task is to create a simulation of the first animal sanctuary on Mars. Make a few types of animals.
Each animal should have a name and adhere to the Stringer interface to return their name.
Every animal should have methods to move and eat. The move method should return a description of the movement.
The eat method should return the name of a random food that the animal likes.
Implement a day/night cycle and run the simulation for three 24-hour sols (72 hours).
All the animals should sleep from sunset until sunrise. For every hour of the day, pick an animal at random to perform a random action (move or eat).
For every action, print out a description of what the animal did.
Your implementation should make use of structures and interfaces.
 */
import (
	"fmt"
	"math/rand"
)

type animal interface{
	eat() string
	move() string
}

// Animal "Lion"
type lion struct{
	name string
}

func (l lion) String () string {
	return l.name
}

func (l lion) move() string {
	switch choice :=  rand.Intn(2); choice {
	case 0 :
		return "Roar!!!"
	case 1:
		return "Hunt!!!"
	default:
		return "Walk!!!"
	}
}

func (l lion) eat() string {
	switch choice :=  rand.Intn(2); choice {
	case 0 :
		return "Deer"
	case 1:
		return "Wildebeest"
	default:
		return "gazelle"
	}
}

// Animal "Elephant"
type elephant struct {
	name string
}

func (e elephant) String() string {
	return e.name
}

func (e elephant) move() string {
	switch choice :=  rand.Intn(2); choice {
	case 0 :
		return "Chill!!!"
	case 1:
		return "Slog!!!"
	default:
		return "Dance!!!"
	}
}

func (e elephant) eat() string {
	switch choice :=  rand.Intn(2); choice {
	case 0 :
		return "Grass"
	case 1:
		return "Leaves"
	default :
		return "root"
	}
}

func choose(a animal) {
	switch choice := rand.Intn(2); choice {
	case 0:
		fmt.Printf("%v %v \n", a , a.move)
	default :
		fmt.Printf("%v eats %v \n" , a , a.eat())
	}
}

func RunAnimalSimulator() {
	animals :=[]animal {
		lion{name : "Kong of the Jungle"} ,
		elephant{name : "Ele"},
	}
	c := rand.Intn(2)
	choose(animals[c])
}





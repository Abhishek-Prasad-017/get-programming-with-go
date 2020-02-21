package unit_6
/*
A knight blocks Arthur’s path. Our hero is empty-handed, represented by a nil value for leftHand *item.
Implement a character struct with methods such as pickup(i *item) and give(to *character).
Then use what you’ve learned in this lesson to write a script that has Arthur pick up an item and give it to the knight, displaying an appropriate description for each action.
 */

import(
	"fmt"
)

type item struct {
	name string
}

type character struct {
	name string
	leftHand *item
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("Person %v is carrying nothing", c.name)
	}
	return fmt.Sprintf("Person %v is carrying a %v", c.name, c.leftHand.name)
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("Person %v picks up an item  %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}

	if c.leftHand == nil {
		fmt.Printf("Person %v has nothing to give\n", c.name)
		return
	}

	if to.leftHand != nil {
		fmt.Printf("Person's %v hands are full\n", to.name)
		return
	}

	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("Person %v gives %v a %v\n", c.name, to.name, to.leftHand.name)
}




func Knight() {
	arthur := &character{name: "Arthur"}
	knight := &character{name: "Knight"}
	gold := &item{name: "Gold"}
	arthur.pickup(gold)
	arthur.give(knight)
	fmt.Println(arthur)
	fmt.Println(knight)
}
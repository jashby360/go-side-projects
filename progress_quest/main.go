package main

import (
	"fmt"
	test "progress_quest/functions"
	"time"
)

func main() {
	var name, class string

	fmt.Print("Input name and class: ")
	fmt.Scanln(&name, &class)
	fmt.Println("Your name is ", name, "and you're an ", class)

	// Take user input
	character := test.CreateCharacter(name, class)
	// Hard coded for testing purposes
	// character := test.CreateCharacter("Hero", "Warrior")
	fmt.Printf("Welcome, %s!\n", character.Name)
	for !character.Dead {
		fmt.Println("Current character death status:", character.Dead)

		quest := test.GenerateRandomQuest()
		// fmt.Println(character)
		test.SimulateCombat(character, quest)
		fmt.Printf("%s is now level %d. \n", character.Name, character.Level)
		time.Sleep(time.Second * 2) // Simulate idle time
	}

	fmt.Println("Character final death status:", character.Dead)

	if character.Dead {
		fmt.Println("Your character is dead, this is the end of the road...")
	} else {
		fmt.Println("Congrats you've made it through without your character dying")
	}

}

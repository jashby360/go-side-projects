package main

import (
	"fmt"
	test "progress_quest/functions"
	"time"
)

func main() {
	character := test.CreateCharacter("Hero", "Warrior")
	fmt.Printf("Welcome, %s!\n", character.Name)
	for !character.Dead {
		quest := test.GenerateRandomQuest()
		fmt.Println(character)
		test.SimulateCombat(character, quest)
		fmt.Printf("%s is now level %d. \n", character.Name, character.Level)
		time.Sleep(time.Second * 2) // Simulate idle time
	}

}

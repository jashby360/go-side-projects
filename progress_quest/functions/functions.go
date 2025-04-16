package functions

import (
	"fmt"
	"math/rand"
	"time"
)

// Character Struct
type Character struct {
	Name  string
	Level int
	Class string
	HP    int
	// Other stats
	EXP int
}

// Quest Struct
type Quest struct {
	Description string
	Reward      int
	Completed   bool
	// Other Quest Information
	Difficulty int
}

func CreateCharacter(name string, class string) *Character {
	return &Character{
		Name:  name,
		Level: 1,
		Class: class,
		HP:    100,
		EXP:   0,
	}
}

// Function
func GenerateQuest() *Quest {
	quests := []string{"Fetch", "Placate", "Deliver"}
	// items := []string{"Shiny Stone", "Rare Flower", "Ancient Artifact"}
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	randomElement := quests[rand.Intn(len(quests))]
	randomDifficulty := rand.Intn(5) + 1

	// fmt.Println(randomElement)

	// A map of different types of quests for various difficulties
	myMap := map[int]map[string][]string{
		1: {
			"Fetch":   []string{"Stone", "Flowers", "Dairy"},
			"Placate": []string{"Cat", "Dog", "Sheep"},
			"Deliver": []string{"Parcel", "Food", "Milk"},
		},
		2: {
			"Fetch":   []string{"Shiny Stone", "Rare Flower", "Artifact"},
			"Placate": []string{"Horse", "Wolf", "Herd"},
			"Deliver": []string{"Heavy Package", "Gilded Comb", "Monster Parts"},
		},
		3: {
			"Fetch":   []string{"Silver", "Exotic Flower", "Ancient Artifact"},
			"Placate": []string{"Undead", "Berserk Animals", "Disturbed Spirits"},
			"Deliver": []string{"Secret Package", "Noble Package", "Governor Package"},
		},
		4: {
			"Fetch":   []string{"Gold", "Moonlight Flower", "Mystical Artifact"},
			"Placate": []string{"Treants", "Evil Spirits", "Forest Spirits"},
			"Deliver": []string{"Exotic Monster Parts", "Nobles Family Heirloom", "Church's Artifact"},
		},
		5: {
			"Fetch":   []string{"Mythril", "Umbra Flower", "Godly Artifact"},
			"Placate": []string{"God", "Demon", "Dragon"},
			"Deliver": []string{"Mystical Artifact", "Legendary Sword", "Alchemist Ingredients"},
		},
	}

	randomIndex := rand.Intn(len(myMap[randomDifficulty][randomElement]))

	return &Quest{
		Description: fmt.Sprintf("%s %s", randomElement, myMap[randomDifficulty][randomElement][randomIndex]),
		Reward:      rand.Intn(100) + 50,
		Completed:   false,
		Difficulty:  randomDifficulty,
	}

}

func SimulateCombat(character *Character, quest *Quest) {
	fmt.Printf("%s is on a quest to %s which is a level %d Quest.\n", character.Name, quest.Description, quest.Difficulty)
	time.Sleep(time.Duration(rand.Intn(5) + 1))
	// add a function that will calculate the difficulty of quest and return if completed
	fmt.Printf("%s has completed the quest!\n", character.Name)
	// if completed set to true and add to exp until, level is hit then restore ecp to zero.
	quest.Completed = true
	character.Level++
}

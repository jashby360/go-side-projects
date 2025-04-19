package functions

import (
	"fmt"
	"math/rand"
	s "progress_quest/structures"
	"time"
)

func CreateCharacter(name string, class string) *s.Character {
	return &s.Character{
		Name:  name,
		Level: 1,
		Class: class,
		HP:    100,
		EXP:   0,
		Dead:  false,
	}
}

// Function
func GenerateRandomQuest() *s.Quest {
	quests := []string{"Fetch", "Placate", "Deliver"}
	// items := []string{"Shiny Stone", "Rare Flower", "Ancient Artifact"}
	rand.New(rand.NewSource(time.Now().UnixNano()))

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

	return &s.Quest{
		Description: fmt.Sprintf("%s %s", randomElement, myMap[randomDifficulty][randomElement][randomIndex]),
		Reward:      rand.Intn(100) + 50,
		Completed:   false,
		Difficulty:  randomDifficulty,
	}

}

func SimulateCombat(character *s.Character, quest *s.Quest) {
	fmt.Printf("%s is on a quest to %s which is a level %d Quest.\n", character.Name, quest.Description, quest.Difficulty)
	time.Sleep(time.Duration(rand.Intn(5) + 1))
	// add a function that will calculate the difficulty of quest and return if completed
	// if completed set to true and add to exp until, level is hit then restore ecp to zero.
	determineQuestCompleted(character, quest)
}

func allotEXP(character *s.Character, reward int) {
	fmt.Println("Before EXP earned:", character.EXP)

	character.EXP += reward

	fmt.Println("After EXP earned:", character.EXP)

	if character.EXP >= 100 {
		character.Level++
		character.EXP = 0
	}

}

func determineDeath(character *s.Character, quest *s.Quest) {

}

func determineQuestCompleted(character *s.Character, quest *s.Quest) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	x := generateRandomNumber(character.Level, float64(quest.Difficulty))

	fmt.Println("Quest complete:", x)

	if x < 0.5 {
		quest.Completed = false
		fmt.Printf("%s has failed the quest!\n", character.Name)
		// determineDeath(character, quest)
	} else {
		quest.Completed = true
		fmt.Printf("%s has completed the quest!\n", character.Name)
		allotEXP(character, quest.Reward)
	}
}

// Function to generate a random number between 0 and 1
// influenced by level and difficulty
func generateRandomNumber(level int, difficulty float64) float64 {
	// Seed the random number generator to ensure different results each time
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Adjust the random number based on level and difficulty
	// Higher level or difficulty might narrow the range or shift the distribution
	modifier := float64(level) * difficulty * 0.1 // Adjust the multiplier as needed

	// Generate a random float between 0 and 1
	randomNumber := rand.Float64()

	// Apply the modifier
	adjustedRandomNumber := randomNumber - modifier

	// Ensure the result is within the range of 0 and 1
	if adjustedRandomNumber < 0 {
		adjustedRandomNumber = 0
	} else if adjustedRandomNumber > 1 {
		adjustedRandomNumber = 1
	}
	return adjustedRandomNumber
}

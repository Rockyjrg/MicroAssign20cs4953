package main

import (
	"fmt"
	"math/rand/v2"
)

type Potion struct {
	effects string
	power   float32
	uses    int
}

func (p *Potion) Generate(seed1, seed2 uint64) {
	pcg := rand.NewPCG(seed1, seed2)
	newRandom := rand.New(pcg)

	listOfEffects := []string{"Invisibility", "Height Boost", "Strength Boost", "Poision Debuff", "Death", "Healing", "God Mode", "Time Stop", "Nothing"}

	randomEffect := newRandom.IntN(len(listOfEffects))

	p.effects = listOfEffects[randomEffect]

	p.power = newRandom.Float32() * 25

	p.uses = newRandom.IntN(10)
}

func main() {
	firstPotion := Potion{}
	firstPotion.Generate(420, 69)
	secondPotion := Potion{}
	secondPotion.Generate(69, 69)
	thirdPotion := Potion{}
	thirdPotion.Generate(25, 25)
	fmt.Printf("Potion Effect: %s | Time Limit: %.2f | Uses: %d\n", firstPotion.effects, firstPotion.power, firstPotion.uses)
	fmt.Printf("Potion Effect: %s | Time Limit: %.2f | Uses: %d\n", secondPotion.effects, secondPotion.power, secondPotion.uses)
	fmt.Printf("Potion Effect: %s | Time Limit: %.2f | Uses: %d\n", thirdPotion.effects, thirdPotion.power, thirdPotion.uses)
}

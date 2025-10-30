package main

import "fmt"

func main() {

}

type attaquant struct {
	name  string
	hp    int
	power int
}

type equipe struct {
	attaquant attaquant
}

func launchGame() {
	var laHorde equipe
	var lalliance equipe

	laHorde.attaquant = attaquant{
		"orc",
		20,
		10,
	}

	lalliance.attaquant = attaquant{
		"elf",
		30,
		5,
	}

	laHorde.attaquant.hp = lalliance.attaquant.power
	if laHorde.attaquant.hp > lalliance.attaquant.power || laHorde.attaquant.hp > 0 {
		laHorde.attaquant.hp--
	} else {
		fmt.Println("l'équipe2 a gagné")
	}

	lalliance.attaquant.hp = laHorde.attaquant.power
	if lalliance.attaquant.hp > laHorde.attaquant.power || laHorde.attaquant.hp < 0 {
		lalliance.attaquant.hp--
	} else {
		fmt.Println("l'équipe1 a gagné")
	}
}

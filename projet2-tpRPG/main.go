package main

import "math/rand"

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

func (a *attaquant) LooseHP(dmg int) {
	a.hp -= dmg
	if a.hp < 0 {
		a.hp = 0
	}
}

func (a *attaquant) Attack() bool {
	return rand.Intn(100)%2 == 0
}

func (a attaquant) Isdead() bool {
	return a.hp <= 0
}

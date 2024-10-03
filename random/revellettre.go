package random

import (
	"fmt"
	"math/rand"
	"time"
)

var Mot = Motsaleatoire()
var DifficultyCounter int

func Revelelettres() []rune {
	var hiddenword []rune
	var runeSlice = []rune(Mot)
	n := rand.Intn(len(Mot) - 1)
	x := rand.Intn(len(Mot) - 1)
	for n == x {
		x = rand.Intn(len(Mot) - 1)
	}

	if DifficultyCounter == 1 {
		for i, v := range runeSlice {
			if i == n {
				hiddenword = append(hiddenword, v)
			} else if v == 13 {
				break
			} else {
				hiddenword = append(hiddenword, '_')
			}
		}
	}

	if DifficultyCounter == 2 {
		for i, v := range runeSlice {
			if i == n {
				hiddenword = append(hiddenword, v)
			} else if i == x {
				hiddenword = append(hiddenword, v)
			} else if v == 13 {
				break
			} else {
				hiddenword = append(hiddenword, '_')
			}
		}
	}
	return hiddenword
}
func Difficultychoix() {
	var input int

	fmt.Println("Choisissez votre difficulté entre 1 ou 2 :")
	fmt.Println("1. Difficulté Facile (vous révélez 2 lettres)")
	fmt.Println("2. Difficulté Difficile (vous révélez 1 lettre)")
	fmt.Scan(&input)
	switch input {
	case 1:
		DifficultyCounter = 2
		Revelelettres()
	case 2:
		DifficultyCounter = 1
		Revelelettres()
	}
	fmt.Print("Très bien, votre difficulté est bien enregistrée.")
	time.Sleep(5 * time.Second)

}

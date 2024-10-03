package random

import (
	"fmt"
	"math/rand"
	"time"
)

var Mot = string(CorrectionMot())
var DifficultyCounter int
var Motcache = Difficultychoix()
var Counter int = 6
var lettrefausse []string

func CorrectionMot() []rune {
	mottemp := Motsaleatoire()
	var motfinal []rune

	for _, v := range mottemp {
		if v == 13 {
			break
		} else {
			motfinal = append(motfinal, v)
		}
	}
	return motfinal
}

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
func Difficultychoix() []rune {
	var input int
	var motchoisis []rune

	fmt.Println("Choisissez votre difficulté entre 1 ou 2 :")
	fmt.Println("1. Difficulté Facile (vous révélez 2 lettres)")
	fmt.Println("2. Difficulté Difficile (vous révélez 1 lettre)")
	fmt.Scan(&input)
	switch input {
	case 1:
		DifficultyCounter = 2
		motchoisis = Revelelettres()
	case 2:
		DifficultyCounter = 1
		motchoisis = Revelelettres()
	}
	fmt.Println("Très bien, votre difficulté est bien enregistrée.")
	time.Sleep(3 * time.Second)
	return motchoisis
}

func Game() {
	var lettre string
	lettretrouve := false

	fmt.Println("\n\nProposez une lettre: ")
	fmt.Scan(&lettre)
	if len(lettre) > 1 {
		if lettre == Mot {
		} else {
			Counter -= 2
		}
	} else {
		for indexmot, v := range Mot {
			if lettre == string(v) {
				lettretrouve = true
				for indexmotcache := range Motcache {
					if indexmot == indexmotcache {
						Motcache[indexmotcache] = v
						break
					}
				}
			}
		}
		if !lettretrouve {
			lettrefausse = append(lettrefausse, lettre)
			Counter--
		}
	}
}

package random

import (
	"fmt"
	"math/rand"
	"time"
)

var Mot = string(CorrectionMot())
var DifficultyCounter int
var Motcache = Difficultychoix()
var Vies int = 6
var Lettrefausse []string

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
	fmt.Print("\033[H\033[2J")
	fmt.Println("\033[32mBienvenue dans le jeu du pendu, amusez-vous bien!\033[0m")
	fmt.Printf("Choisissez votre difficulté entre 1 ou 2 :")
	fmt.Printf("\n\n\033[34m1. Difficulté Facile (vous révélez 2 lettres)\033[0m")
	fmt.Printf("\n\033[34m2. Difficulté Difficile (vous révélez 1 lettre)\033[0m\n")
	fmt.Println("Votre choix : ")
	fmt.Scan(&input)

	for input != 1 && input != 2 {
		fmt.Println("Choix invalide réessayez ! ")
		fmt.Scan(&input)
	}

	switch input {
	case 1:
		DifficultyCounter = 2
		motchoisis = Revelelettres()
	case 2:
		DifficultyCounter = 1
		motchoisis = Revelelettres()
	}
	fmt.Println("Très bien, votre difficulté est bien enregistrée.")
	time.Sleep(2 * time.Second)
	return motchoisis
}

func Game() {
	var lettre string
	lettretrouve := false
	fmt.Println("\n\nProposez une lettre: ")
	fmt.Scan(&lettre)

	for _, v := range Lettrefausse {
		if v == lettre {
			fmt.Println("Lettre déjà trouvé ! ")
			fmt.Scan(&lettre)
		}
	}

	if len(lettre) > 1 {
		if lettre == Mot {
			Motcache = []rune(Mot)
		}

		if lettre != Mot {
			fmt.Println("Votre mot est incorrect. Vous perdez 2 vies.")
			time.Sleep(2 * time.Second)
			Vies -= 2
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
			fmt.Print("\033[31mLettre incorrecte. Vous perdez 1 vie.\033[0m\n")
			time.Sleep(2 * time.Second)
			Lettrefausse = append(Lettrefausse, lettre)
			Vies--
		}
	}
}

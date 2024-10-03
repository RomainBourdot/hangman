package jeu

import (
	"fmt"
	"hangman/random"
)

func Jeu() {
	for random.Mot != string(random.Motcache) && random.Counter > 0 {
		fmt.Println(random.Mot)
		fmt.Printf("\nCompteur: %v", random.Counter)
		fmt.Printf("\nMot: %v", string(random.Motcache))
		random.Game()
	}
}

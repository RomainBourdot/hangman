package jeu

import (
	"fmt"
	"hangman/random"
)

func Jeu() {
	fmt.Printf("Mot:%v", string(random.Revelelettres()))
}

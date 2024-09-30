package main

import (
	"fmt"
	"hangman/random"
)

func main() {
	mot := random.Motsaleatoire()
	fmt.Println(random.Revelelettres(mot, 2))
}

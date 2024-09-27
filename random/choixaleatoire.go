package random

import (
	"math/rand"
	"os"
	"strings"
)

func Motsaleatoire() string {
	file, _ := os.ReadFile("mots.txt")
	data := string(file)

	words := strings.Split(data, "\n")
	n := rand.Intn(19)
	word := ""
	for i, elem := range words {
		if i == n {
			word = elem
		}
	}
	return word
}

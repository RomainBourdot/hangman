package random

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func Motsaleatoire() string {
	file, err := os.ReadFile("./random/mots.txt")
	data := string(file)
	if err != nil {
		fmt.Println(err)
	}

	words := strings.Split(data, "\n")
	n := rand.Intn(83)
	word := ""
	for i, elem := range words {
		if i == n {
			word = elem
		}
	}
	return word
}

package jeu

import (
	"fmt"
	"hangman/dessin"
	"hangman/random"
	"strings"
)

func Jeu() {
	for random.Mot != string(random.Motcache) && random.Vies > 0 {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("\nCompteur:\033[96m %v\033[0m", random.Vies)
		fmt.Println(dessin.Dessinparvie())
		fmt.Printf("\n\033[90mMot\033[0m: %v", string(random.Motcache))
		lettre := strings.Join(random.Lettrefausse, ", ")
		fmt.Printf("\n\nLettres déjà utilisés :%v", lettre)
		random.Game()
	}
	if random.Mot == string(random.Motcache) {
		fmt.Printf("\n\033[32mBien joué vous avez gagné ! Le mot était :\033[0m\033[95m%s\033[0m", random.Mot)
	} else {
		fmt.Println("\n\033[91mDésolé, vous avez perdu toutes vos vies.\033[0m")
		fmt.Printf("Le mot correct était : \033[95m%v\033[0m\n", random.Mot)
	}
}

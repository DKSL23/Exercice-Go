package main

import (
	"fmt"
	"os"
)

const MaxDisplay = 10

func main() {

	taille := len(os.Args[1:])

	fmt.Println("Le nombre total de mots est de", taille)

	if taille == 0 {
		fmt.Println("Aucun mot n'a été entré")
		os.Exit(1)
	}
	for i := 0; i < taille; i++ {
		if len(os.Args[i+1]) >= 4 {
			fmt.Println(os.Args[i+1])
		}
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

const colonnes = 7
const lignes = 6

type Grille [colonnes + 2][lignes + 2]string

func main() {
	keyboard.Open()
	defer keyboard.Close()

	var grille = remplirDeVide()
	afficher(grille)

	for {
		grille = choisirEtJouer(grille, "X")
		grille = choisirEtJouer(grille, "O")
		if matchNul(grille) {
			fmt.Println("Match nul")
			return
		}
	}
}

func afficher(grille Grille) {
	for x := 1; x <= colonnes; x++ {
		fmt.Print(" ")
		fmt.Print(x)
	}
	fmt.Println()
	for y := 1; y <= lignes; y++ {
		for x := 1; x <= colonnes; x++ {
			fmt.Print("|")
			fmt.Print(grille[x][y])
		}
		fmt.Println("|")
	}
	fmt.Println()
}

func remplirDeVide() Grille {
	var grille Grille
	for y := 1; y <= lignes; y++ {
		for x := 1; x <= colonnes; x++ {
			grille[x][y] = " "
		}
	}

	return grille
}

func choisirEtJouer(grille Grille, joueur string) Grille {
	for {
		touche, _, _ := keyboard.GetKey()
		var colonne = int(touche - '0')
		if peutJouer(grille, colonne) {
			var grille, ligne = jouer(grille, colonne, joueur)
			afficher(grille)
			if gagnant(grille, colonne, ligne) {
				fmt.Println(joueur, "gagne !")
				os.Exit(0)
			}

			return grille
		} else {
			fmt.Println("Choisissez une autre colonne")
		}
	}
}

func jouer(grille Grille, colonne int, joueur string) (Grille, int) {
	for y := lignes; y >= 1; y-- {
		if grille[colonne][y] == " " {
			grille[colonne][y] = joueur
			return grille, y
		}
	}

	panic("ne peut pas jouer")
}

func peutJouer(grille Grille, colonne int) bool {
	if grille[colonne][1] == " " {
		return true
	}

	return false
}

func matchNul(grille Grille) bool {
	for x := 1; x <= colonnes; x++ {
		if grille[x][1] == " " {
			return false
		}
	}
	return true
}

func gagnant(grille Grille, colonne int, ligne int) bool {
	if mesurer(grille, colonne, ligne, 1, 0) >= 4 {
		return true
	}
	if mesurer(grille, colonne, ligne, 0, 1) >= 4 {
		return true
	}
	if mesurer(grille, colonne, ligne, 1, -1) >= 4 {
		return true
	}
	if mesurer(grille, colonne, ligne, 1, 1) >= 4 {
		return true
	}

	return false
}

func mesurer(grille Grille, colonne int, ligne int, sensColonne int, sensLigne int) int {
	var joueur = grille[colonne][ligne]
	for grille[colonne+sensColonne][ligne+sensLigne] == joueur {
		colonne = colonne + sensColonne
		ligne = ligne + sensLigne
	}
	var longueur = 1
	for grille[colonne-sensColonne][ligne-sensLigne] == joueur {
		colonne = colonne - sensColonne
		ligne = ligne - sensLigne
		longueur = longueur + 1
	}

	return longueur
}

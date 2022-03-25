package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

const colonnes = 7
const lignes = 6

var grille [colonnes + 2][lignes + 2]string

func main() {
	keyboard.Open()
	defer keyboard.Close()
	remplirDeVide()
	afficher()

	for {
		choisirEtJouer("X")
		choisirEtJouer("O")
		if matchNul() {
			fmt.Println("Match nul")
			return
		}
	}
}

func afficher() {
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

func remplirDeVide() {
	for y := 1; y <= lignes; y++ {
		for x := 1; x <= colonnes; x++ {
			grille[x][y] = " "
		}
	}
}

func choisirEtJouer(joueur string) {
	for {
		touche, _, _ := keyboard.GetKey()
		var colonne = int(touche - '0')
		if peutJouer(colonne) {
			var ligne = jouer(colonne, joueur)
			afficher()
			if gagnant(colonne, ligne) {
				fmt.Println(joueur, "gagne !")
				return
			}
			if matchNul() {
				fmt.Println("Match nul")
				return
			}
			break
		} else {
			fmt.Println("Choisissez une autre colonne")
		}
	}
}

func jouer(colonne int, joueur string) int {
	for y := lignes; y >= 1; y-- {
		if grille[colonne][y] == " " {
			grille[colonne][y] = joueur
			return y
		}
	}

	return -1
}

func peutJouer(colonne int) bool {
	if grille[colonne][1] == " " {
		return true
	}

	return false
}

func matchNul() bool {
	for x := 1; x <= colonnes; x++ {
		if grille[x][1] == " " {
			return false
		}
	}
	return true
}

func gagnant(colonne int, ligne int) bool {
	if mesurer(colonne, ligne, 1, 0) >= 4 {
		return true
	}
	if mesurer(colonne, ligne, 0, 1) >= 4 {
		return true
	}
	if mesurer(colonne, ligne, 1, -1) >= 4 {
		return true
	}
	if mesurer(colonne, ligne, 1, 1) >= 4 {
		return true
	}

	return false
}

func mesurer(colonne int, ligne int, sensColonne int, sensLigne int) int {
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

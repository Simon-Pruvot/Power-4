package main

import (
	"net/http"
)

type pageData struct {
	Grille        [][]string
	Colonnes      []int
	joueur        []string
	indiceJoueur  int
	botverif      [][]int
	botProfondeur int
}

var data pageData

func (data *pageData) verif(ligne int, col int, pion string, w http.ResponseWriter, r *http.Request) int {
	stop := true
	for i := 0; i < len(data.Colonnes) && stop; i++ {
		if data.Grille[0][i] == "/images/pion0.png" {
			stop = false
		}
	}
	if stop {
		http.Redirect(w, r, "/egalite", http.StatusSeeOther)
	}
	max := 1
	compteur := 1
	for i := ligne - 1; i > -1; i-- {
		if data.Grille[col][i] == pion {
			compteur += 1
		} else {
			break
		}
	}
	for i := ligne + 1; i < len(data.Grille[0])-1; i++ {
		if data.Grille[col][i] == pion {
			compteur += 1
		} else {
			break
		}
	}
	if max < compteur {
		max = compteur
	}
	if compteur >= 4 {
		return max
	} else {
		compteur = 1
		for i := col + 1; i < len(data.Grille); i++ {
			if data.Grille[i][ligne] == pion {
				compteur += 1
			} else {
				break
			}
		}
		if max < compteur {
			max = compteur
		}
		if compteur >= 4 {
			return compteur
		} else {
			compteur = 1
			for i := 1; col-i >= 0 && ligne-i >= 0; i++ {
				if data.Grille[col-i][ligne-i] == pion {
					compteur += 1
				} else {
					break
				}
			}
			for i := 1; col+i <= len(data.Grille)-1 && ligne+i <= len(data.Grille[0])-1; i++ {
				if data.Grille[col+i][ligne+i] == pion {
					compteur += 1
				} else {
					break
				}
			}
			if max < compteur {
				max = compteur
			}
			if compteur >= 4 {
				return compteur
			} else {
				compteur = 1
				for i := 1; col+i <= len(data.Grille)-1 && ligne-i >= 0; i++ {
					if data.Grille[col+i][ligne-i] == pion {
						compteur += 1
					} else {
						break
					}
				}
				for i := 1; col-i >= 0 && ligne+i <= len(data.Grille[0])-1; i++ {
					if data.Grille[col-i][ligne+i] == pion {
						compteur += 1
					} else {
						break
					}
				}
				if max < compteur {
					max = compteur
				}
				if compteur >= 4 {
					return compteur
				}
			}

		}
	}
	return max
}

func chercherLigne(index int) int {
	for i := len(data.Grille) - 1; i >= 0; i-- {
		if data.Grille[i][index] == "/images/pion0.png" {
			return i
		}
	}
	return -1
}

func (data *pageData) ajouterPion(index int, grille [][]string, w http.ResponseWriter, r *http.Request) bool {
	ligne := chercherLigne(index)
	if ligne != -1 {
		grille[ligne][index] = data.joueur[data.indiceJoueur]
		if data.verif(index, ligne, data.joueur[data.indiceJoueur], w, r) >= 4 {
			return true
		}
		data.indiceJoueur = (data.indiceJoueur + 1) % 2
		temp := make([][]int, data.botProfondeur)
		for i := range temp {
			temp[i] = make([]int, 2)
		}
		//RechercheBot(0, data.Grille, temp)
	}
	return false
}

func nouvelleGrille(rows, cols int) [][]string {
	grille := make([][]string, rows)
	for i := range grille {
		grille[i] = make([]string, cols)
		for j := range grille[i] {
			grille[i][j] = "/images/pion0.png"
		}
	}
	return grille
}

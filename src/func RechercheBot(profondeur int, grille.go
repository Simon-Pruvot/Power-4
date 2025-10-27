func RechercheBot(profondeur int, grille [][]string, temp [][]int) {
	copie := make([][]string, len(grille))
	for i := range grille {
		copie[i] = append([]string(nil), grille[i]...)
	}
	for i := 0; i < len(grille[0]); i++ {
		ligne := chercherLigne(i)
		if ligne != -1 {
			verif := data.verif(i, ligne, data.joueur[profondeur%2])
			if verif > temp[i][1] && profondeur%2 == 0 {
				temp[profondeur%2][1] = i
				temp[profondeur%2][0] = verif
			} else if verif < temp[i][1] && profondeur%2 == 1 {
				temp[profondeur%2][1] = i
				temp[profondeur%2][0] = verif
			}
		}
	}

	if temp[1][1] == 4 && temp[0][1] != 4 {
		temp[0] = temp[1]
	}

	ligne := chercherLigne(temp[profondeur][1])
	if ligne != -1 {
		copie[ligne][temp[profondeur][1]] = data.joueur[profondeur%2]
	}
	if profondeur == data.botProfondeur {
		fmt.Println("choix du bot:", temp)
		return
	} else {
		RechercheBot(profondeur+1, data.Grille, temp)
	}
}

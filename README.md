# Power-4 — Puissance 4 en Go

Une application web en Go qui permet de jouer au Puissance 4 à deux joueurs, avec personnalisation de la taille de la grille et des jetons (galerie d’images et upload). Interface HTML/CSS, rendu avec `html/template`, serveur HTTP standard (`net/http`).

Auteurs (issus de la page d’accueil) : Simon Pruvot, Kevin Bura, Simon Beausire

## Fonctionnalités

- Jeu Puissance 4 jouable à deux en local (même navigateur)
- Grille personnalisable (lignes/colonnes via la page « Personalisation »)
- Choix/Personnalisation des jetons via une galerie d’images locale (`/merch`)
- Upload d’images (ex. via la page caméra) pour enrichir la galerie
- Pages dédiées : accueil, règles, victoire, égalité
- Feuilles de style par page, base CSS commune

## Pile technique

- Go (net/http, html/template)
- HTML/CSS statiques servis depuis `src/CSS` et `src/template`
- Aucune dépendance externe côté Go

## Prérequis

- Go 1.20+ recommandé
	- Si vous voyez une erreur « invalid go version » liée à la directive `go 1.25.0` dans `go.mod`, remplacez-la par une version valide disponible sur votre machine (ex. `go 1.22`).
- Windows, macOS ou Linux

## Installation

1. Cloner le dépôt
2. Placer le terminal à la racine du repo

## Lancement (développement)

Le serveur écoute actuellement sur le port 80 (voir `src/main.go`). Sur Windows, le port 80 peut nécessiter les droits administrateur.

Options :

- Option simple (recommandée) : modifier `http.ListenAndServe(":80", nil)` en `":8080"`, puis lancer l’appli et ouvrir http://localhost:8080
- Option administrateur : conserver `:80` et lancer le binaire/commande avec des droits élevés

Points d’entrée possibles (selon votre configuration Go) :

- go run
	- `go run ./src`
	- ou `go run src/main.go`

- build puis exécution
	- `go build -o power4 ./src`
	- `./power4` (ou `power4.exe` sur Windows)

Si vous rencontrez une erreur liée à `go.mod` (directive `go 1.25.0`), ouvrez le fichier et remplacez simplement la ligne `go 1.25.0` par `go 1.22` (ou la version installée), puis relancez la commande.

## Routes principales

- `/` : page d’accueil
- `/diff` : choix du mode (classique vs personnalisation)
- `/temp` : initialise une grille classique 6x7 et redirige vers `/play`
- `/personalisation` : sélection du nombre de lignes/colonnes, soumet vers `/play`
- `/play` :
	- GET avec `rows`/`cols` pour générer la grille
	- POST avec `col` pour « lâcher » un jeton dans une colonne
- `/merch` : galerie d’images locales (répertoire `images/`) et sélection de skins des joueurs
- `/camera` : page caméra (form d’upload)
- `/uploadphoto` : POST d’une image, ajoutée dans `images/`
- `/victoire` : page de victoire (utilise le paramètre `winner`)
- `/egalite` : page d’égalité
- `/regle` : page des règles

## Structure du projet

```
go.mod
src/
	main.go            # démarrage du serveur, routes, fichiers statiques
	handlers.go        # gestionnaires HTTP (pages, upload, logique de navigation)
	game.go            # logique du jeu (grille, vérifications, coups)
	CSS/               # styles par page + base.css
	template/          # templates HTML (accueil, jeu, règles, etc.)
	images/            # images des pions et ressources (écritures de fichiers upload)
```

## Astuces et notes

- Port d’écoute : si le port 80 est bloqué, utilisez `:8080` pour le développement local.
- Images/skins : placez des images dans le dossier `src/images` pour qu’elles apparaissent dans la galerie `/merch`.
- Upload : les photos envoyées sont stockées dans `src/images` avec un nom unique.

## Contributions

Les PRs et issues sont bienvenues. Pour une contribution :

1. Créez une branche à partir de `main`
2. Faites des changements atomiques et testables
3. Ouvrez une PR avec une description claire (contexte, changements, tests)

## Licence

Non spécifiée pour le moment.

---

## CSS structure (EN)

Styles were modularized to keep the design identical while making the code easier to navigate:

- src/CSS/base.css: shared base (body, buttons, header, common utilities)
- src/CSS/index.css: styles used by template/index.html
- src/CSS/diff.css: styles used by template/diff.html
- src/CSS/play.css: styles used by template/play.html
- src/CSS/regle.css: styles used by template/regle.html
- src/CSS/merch.css: styles used by template/merch.html (skins gallery + camera link)
- src/CSS/victoire.css: styles used by template/victoire.html
- src/CSS/egalite.css: styles used by template/egalite.html
- src/CSS/personalisation.css: styles used by template/personalisation.html
- src/CSS/camera.css: page-specific overrides for template/camera.html

Each HTML template now links to base.css plus its own page stylesheet. The original src/CSS/style.css remains in the repo for reference but is no longer used by templates.

Tip: Add new global utilities to base.css, and keep page-only rules in that page’s CSS to avoid regressions.


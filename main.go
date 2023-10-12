package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Initialisation de raylib
	rl.InitWindow(800, 600, "Hangman Game")
	rl.SetTargetFPS(60)

	rand.Seed(time.Now().UnixNano())

	dico := []string{"AVION", "ARBRE", "CAMIONS", "CHEVEUX", "CAMPING", "CHAUSSURE", "FOUET", "CHAMPS", "HORLOGE", "CHIPS"} // liste de mot possible
	wordToGuess := dico[rand.Intn(len(dico))] // le mot à deviner choisi aléatoirement dans la taille du tableau
	guessedWord := make([]string, len(wordToGuess))

	for i := range guessedWord {
		guessedWord[i] = "_"
	}

	attemptsLeft := 6
	guessedLetters := make([]string, 0)

	for !rl.WindowShouldClose() { // La boucle principale du jeu

		if attemptsLeft == 0 || strings.Join(guessedWord, "") == wordToGuess {
			break
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		// Dessiner l'interface utilisateur
		rl.DrawText(fmt.Sprintf("Vie restantes : %d", attemptsLeft), 10, 10, 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Lettres utilisé : %s", strings.Join(guessedLetters, " ")), 10, 40, 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Mot caché : %s", strings.Join(guessedWord, " ")), 10, 70, 20, rl.Black)

		var guess string // Je code ca comme un golmon mais je sais pas comment le faire autrement ( j'obtimiserais le code plus tard )
		if rl.IsKeyDown(rl.KeyQ) {
			guess = "A"
		} else if rl.IsKeyDown(rl.KeyB) {
			guess = "B"
		} else if rl.IsKeyDown(rl.KeyC) {
            guess = "C"
        } else if rl.IsKeyDown(rl.KeyD) {
            guess = "D"
        } else if rl.IsKeyDown(rl.KeyE) {
            guess = "E"
        } else if rl.IsKeyDown(rl.KeyF) {
            guess = "F"
        } else if rl.IsKeyDown(rl.KeyG) {
            guess = "G"
        } else if rl.IsKeyDown(rl.KeyH) {
            guess = "H"
        } else if rl.IsKeyDown(rl.KeyI) {
            guess = "I"
        } else if rl.IsKeyDown(rl.KeyJ) {
            guess = "J"
        } else if rl.IsKeyDown(rl.KeyK) {
            guess = "K"
        } else if rl.IsKeyDown(rl.KeyL) {
            guess = "L"
        } else if rl.IsKeyDown(rl.KeyM) {
            guess = "M"
        } else if rl.IsKeyDown(rl.KeyN) {
            guess = "N"
        } else if rl.IsKeyDown(rl.KeyO) {
            guess = "O"
        } else if rl.IsKeyDown(rl.KeyP) {
            guess = "P"
        } else if rl.IsKeyDown(rl.KeyA) {
            guess = "Q"
        } else if rl.IsKeyDown(rl.KeyR) {
            guess = "R"
        } else if rl.IsKeyDown(rl.KeyS) {
            guess = "S"
        } else if rl.IsKeyDown(rl.KeyT) {
            guess = "T"
        } else if rl.IsKeyDown(rl.KeyU) {
            guess = "U"
        } else if rl.IsKeyDown(rl.KeyV) {
            guess = "V"
        } else if rl.IsKeyDown(rl.KeyZ) {
            guess = "W"
        } else if rl.IsKeyDown(rl.KeyX) {
            guess = "X"
        } else if rl.IsKeyDown(rl.KeyY) {
            guess = "Y"
        } else if rl.IsKeyDown(rl.KeyW) {
            guess ="Z"
        }
		if guess != "" {
			guess = strings.ToUpper(guess)

			if strings.Contains(wordToGuess, guess) {
				for i, letter := range wordToGuess {
					if string(letter) == guess {
						guessedWord[i] = guess
					}
				}
			} else {
				fmt.Printf("La lettre %s n'est pas dans le mot.\n", guess)
				attemptsLeft--
			}

			guessedLetters = append(guessedLetters, guess)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()

	if attemptsLeft == 0 {
		fmt.Println("Perdu ! le mot était :", wordToGuess)
	} else {
		fmt.Println("GG tu as trouvé le mot:", wordToGuess)
	}
}

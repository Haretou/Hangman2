package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "Hangman Game")
	rl.SetTargetFPS(60)

	rand.Seed(time.Now().UnixNano())

	dico := []string{"AVION", "ARBRE", "CAMIONS", "CHEVEUX", "CAMPING", "CHAUSSURE", "FOUET", "CHAMPS", "HORLOGE", "CHIPS", "ORDINATEUR", "TABLE", "APERO", "TABLEAU", "WAGON", "CAILLOUX"}
	wordToGuess := dico[rand.Intn(len(dico))]
	guessedWord := make([]string, len(wordToGuess))

	for i := range guessedWord {
		guessedWord[i] = "_"
	}

	attemptsLeft := 6
	guessedLetters := make([]string, 0)

	for !rl.WindowShouldClose() {
		if attemptsLeft == 0 || strings.Join(guessedWord, "") == wordToGuess {
			break
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		rl.DrawText(fmt.Sprintf("Vies restantes : %d", attemptsLeft), 10, 10, 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Lettres utilisées : %s", strings.Join(guessedLetters, " ")), 10, 40, 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Mot caché : %s", strings.Join(guessedWord, " ")), 10, 70, 20, rl.Black)
        // Cette parti est récupée d'un cours de Gamecodeur - L'école de jeux vidéo en Ligne ( une chaine Youtube )
        // Elle sert a initié chaque lettre 
		if rl.IsKeyPressed(rl.KeyQ) {
			makeGuess("A", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyB) {
			makeGuess("B", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyC) {
			makeGuess("C", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyD) {
			makeGuess("D", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyE) {
			makeGuess("E", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyF) {
			makeGuess("F", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyG) {
			makeGuess("G", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyH) {
			makeGuess("H", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyI) {
			makeGuess("I", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyJ) {
			makeGuess("J", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyK) {
			makeGuess("K", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyL) {
			makeGuess("L", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyM) {
			makeGuess("M", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyN) {
			makeGuess("N", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyO) {
			makeGuess("O", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyP) {
			makeGuess("P", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyA) {
			makeGuess("Q", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyR) {
			makeGuess("R", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyS) {
			makeGuess("S", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyT) {
			makeGuess("T", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyU) {
			makeGuess("U", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyV) {
			makeGuess("V", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyZ) {
			makeGuess("W", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyX) {
			makeGuess("X", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyY) {
			makeGuess("Y", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		} else if rl.IsKeyPressed(rl.KeyW) {
			makeGuess("Z", wordToGuess, guessedWord, &attemptsLeft, &guessedLetters)
		}
        // Fin de la partie récupée 
		rl.EndDrawing()
	}

	rl.CloseWindow()

	if attemptsLeft == 0 {
		fmt.Println("Perdu ! Le mot était :", wordToGuess)
	} else {
		fmt.Println("GG ! Tu as trouvé le mot :", wordToGuess)
	}
}

func makeGuess(guess string, wordToGuess string, guessedWord []string, attemptsLeft *int, guessedLetters *[]string) {
	guess = strings.ToUpper(guess)

	if strings.Contains(wordToGuess, guess) {
		for i, letter := range wordToGuess {
			if string(letter) == guess {
				guessedWord[i] = guess
			}
		}
	} else {
		fmt.Printf("La lettre %s n'est pas dans le mot.\n", guess)
		*attemptsLeft--
	}

	*guessedLetters = append(*guessedLetters, guess)
}

package main

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    dico := []string{"AVION", "ARBRE", "CAMIONS", "CHEVEUX", "CAMPING", "CHAUSSURE"}
    wordToGuess := dico[rand.Intn(len(dico))] // le mot à deviner
    guessedWord := make([]string, len(wordToGuess))

    for i := range guessedWord {
        guessedWord[i] = "_"
    }

    attemptsLeft := 6
    guessedLetters := make([]string, 0)

    fmt.Println("Welcome to the Hangman game!")
    fmt.Printf("Guess the word: %s\n", strings.Join(guessedWord, " "))

    for attemptsLeft > 0 { // Tant qu'il reste des "vies" au joueur
        fmt.Printf("Attempts left: %d\n", attemptsLeft)
        fmt.Print("Guessed letters: ")
        fmt.Println(strings.Join(guessedLetters, " "))

        var guess string
        fmt.Print("Guess a letter: ")
        fmt.Scanf("%s", &guess)

        if len(guess) != 1 {
            fmt.Println("Please enter a single letter.")
            continue
        }

        guess = strings.ToUpper(guess) // Convertir la lettre en majuscule pour correspondre au mot en majuscule

        if strings.Contains(wordToGuess, guess) {
            for i, letter := range wordToGuess {
                if string(letter) == guess {
                    guessedWord[i] = guess
                }
            }
        } else {
            fmt.Printf("The letter %s is not in the word.\n", guess)
            attemptsLeft--
        }

        guessedLetters = append(guessedLetters, guess)

        if strings.Join(guessedWord, "") == wordToGuess {
            fmt.Println("Congratulations! You have guessed the word:", wordToGuess)
            break
        }

        fmt.Printf("Current word: %s\n", strings.Join(guessedWord, " "))
    }

    if attemptsLeft == 0 {
        fmt.Println("You lost! The word was:", wordToGuess)
    }
}

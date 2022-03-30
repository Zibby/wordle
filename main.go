package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

var max_guesses int = 5
var guesses_count int
var words []string
var game_word string
var game_word_chars []string
var guess string
var won bool = false

func loadWords() {
	file, err := os.Open("/usr/share/dict/british-english")
	defer file.Close()
	if err != nil {
		log.Fatalf("Could not open words list, is the words package installed?")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		match, _ := regexp.MatchString("^[a-z]{5}$", scanner.Text())
		if match {
			words = append(words, scanner.Text())
		}
	}
}

func pickWord() {
	rand.Seed(time.Now().UnixNano())
	game_word = words[rand.Intn(len(words))]
	fmt.Println(game_word)
	game_word_chars = strings.Split(game_word, "")
	fmt.Println(game_word_chars)
}

func init() {
	loadWords()
	pickWord()
}

func compareGuess() {
	if guess == game_word {
		fmt.Println("You win")
		won = true
		return
	}

	for i, letter := range strings.Split(guess, "") {
		if game_word_chars[i] == letter {
			d := color.New(color.FgGreen, color.Bold)
			d.Printf(letter)

		} else {
			d := color.New(color.FgRed, color.Underline)
			d.Printf(letter)
		}
	}
	println("")
}

func readGuess() {
	fmt.Printf("Guess a 5 letter word, you are on guess %v of %v\n", guesses_count, max_guesses)
	fmt.Scanln(&guess)
	compareGuess()
}

func main() {
	for guesses_count = 0; guesses_count < max_guesses; guesses_count++ {
		readGuess()
		if won == true {
			break
		}
	}
}

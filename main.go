package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
)

func getPic(index int) string {
	HANGMANPICS := []string{`
+---+
|   |
|
|
|
|
=========`,
		`
+---+
|   |
|   O  
|  
|  
|
=========`, `
+---+
|   |
|   O  
|  /
|  
|
=========`, `
+---+
|   |
|   O  
|  /|
|  
|
=========`, `
+---+
|   |
|   O  
|  /|\
|  
|
=========`, `
+---+
|   |
|   O  
|  /|\
|  / 
|
=========`, `
+---+
|   |
|   O  
|  /|\
|  / \
|
=========`,
	}
	return HANGMANPICS[index]
}

func getRandomWord() string {
	words := []string{"hello", "world"}
	randomIndex := rand.Intn(len(words))
	pick := words[randomIndex]
	return pick
}

func main() {
	hangCount := 0
	word := getRandomWord()
	m := map[string]bool{}
	w := getWordMap(word)
	remaining := len(word)
	fmt.Println("word has " + strconv.Itoa(remaining) + " characters")

	for hangCount < 6 {
		input := getInput("")
		_, ok := m[input]
		isValid := isValidInput(input)
		for isValid == false || ok {
			if isValid == false {
				input = getInput("Input should be a single letter\n")
			} else {
				input = getInput("Duplicate character.\n")
			}
			_, ok = m[input]
			isValid = isValidInput(input)
		}
		// string is unique
		m[input] = true
		keys := reflect.ValueOf(m).MapKeys()
		fmt.Println("Your guesses:", keys)

		printUnderlineWord(word, m)
		fmt.Println()

		_, ok = w[input]
		if ok {
			fmt.Println(input + " is in the word.")
			remaining -= int(w[input])
			if remaining == 0 {
				fmt.Println("You won.")
				break
			} else {
				fmt.Println(strconv.Itoa(remaining) + " characters remaining")
			}
		} else {
			fmt.Println(input + " is not in the word.")
			hangCount++
		}

		fmt.Println(getPic(hangCount))
	}
	fmt.Println("It was " + word)
}

func isValidInput(input string) bool {
	if len(input) > 1 {
		return false
	}
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(input) {
		return false
	}
	return true
}

func printUnderlineWord(word string, m map[string]bool) {
	for _, c := range word {
		_, ok := m[string(c)]
		s := string(c)
		if ok == false {
			s = "_"
		}
		fmt.Print(s + " ")
	}
}

func getWordMap(word string) map[string]int32 {
	w := map[string]int32{}
	for _, c := range word {
		w[string(c)]++
	}
	return w
}

func getInput(message string) string {
	fmt.Println(message + "Enter character: ")
	var input string
	fmt.Scanln(&input)
	return input
}

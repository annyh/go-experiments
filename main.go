package main

import (
	"fmt"
	"strconv"
)

// TODO: randomize word
func main() {
	count := 0
	word := "hello"
	m := map[string]bool{}
	w := getWordMap(word)
	remaining := len(word)
	fmt.Println("word has " + strconv.Itoa(remaining) + " characters")

	for count < 6 {
		input := getInput("")
		_, ok := m[input]
		for ok {
			input = getInput("Duplicate character. ")
			_, ok = m[input]
		}
		// string is unique
		m[input] = true
		// if string is one of the characters, say got character right
		_, ok = w[input]
		if ok {
			fmt.Println(input + " is in the word.")
			remaining -= int(w[input])
			fmt.Println(strconv.Itoa(remaining) + " characters remaining")
			if remaining == 0 {
				fmt.Println("You won.")
				break
			}
		}
		count++
	}
	fmt.Println("It was " + word)
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

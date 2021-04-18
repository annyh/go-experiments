package main

import "fmt"

func main() {
	count := 0
	word := "hey"
	m := map[string]bool{}
	for count < 6 {
		input := getInput("")
		_, ok := m[input]
		for ok {
			input = getInput("Duplicate character. ")
			_, ok = m[input]
		}
		// string is unique
		m[input] = true
		count++
	}
	fmt.Println("It was " + word)
}

func getInput(message string) string {
	fmt.Println(message + "Enter character: ")
	var input string
	fmt.Scanln(&input)
	return input
}

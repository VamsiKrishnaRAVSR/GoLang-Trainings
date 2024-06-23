package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"
	entries := map[string]bool{}
	placeholder := []string{}
	// slice := make([]string, len(word), len(word))

	i := 0
	for i = 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	chances := 8
	for {
		userInput := strings.Join(placeholder, "")

		if chances == 0 && userInput != word {
			fmt.Println("Game Over! Try again")
			break
		}
		if userInput == word {
			fmt.Println("You Win!!")
			break
		}

		fmt.Println("\n")
		fmt.Println(placeholder)  // render the placeholder
		fmt.Printf("%d", chances) // render the chances left
		keys := []string{}
		for k, _ := range entries {
			keys = append(keys, k)
		}
		fmt.Println(keys) // show the letters or words guessed till now.
		fmt.Printf("Guess a letter or the word: ")

		// take the input
		str := ""
		fmt.Scanln(&str)
		// compare and  update entries placeholder and chances
		_, ok := entries[str]

		if ok {
			continue
		}
		entries[str] = true

		found := false
		temp := strings.Split(word, "")
		for i, v := range temp {
			if v == str {
				placeholder[i] = v
				found = true
			}
		}
		if !found {
			chances -= 1
		}

	}
}

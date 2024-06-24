package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func get_word() string {
	resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
	if err != nil {
		return "elephant"
	}
	defer resp.Body.Close()
	words := []string{}

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &words)
	// fmt.Printf("%s", body)
	return words[0]
}

func main() {
	word := get_word()
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
			fmt.Printf("%s \n", word)
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

		// if str length <2
		// compare and  update entries placeholder and chances
		if len(str) > 2 {
			if str == word {
				fmt.Println("You win!")
				break
			} else {
				chances -= 1
				entries[str] = true
				continue
			}
		}
		_, ok := entries[str]

		if ok {
			continue
		}

		found := false
		for i, v := range word {
			if str == string(v) {
				placeholder[i] = string(v)
				found = true
			}
		}
		// temp := strings.Split(word, "")
		// for i, v := range temp {
		// 	if  {
		// placeholder[i] = v
		// found = true
		// 	}
		// }
		if !found {
			entries[str] = true
			chances -= 1
		}

	}
}

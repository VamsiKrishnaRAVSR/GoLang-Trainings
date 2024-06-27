package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var dev = flag.Bool("dev", false, "dev mode")

func get_word() string {
	if *dev {
		return "elephant"
	}
	resp, err := http.Get("https://random-word-api.herokuapp.com/all")
	if err != nil {
		return "elephant"
	}
	defer resp.Body.Close()
	words := []string{}

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &words)
	return words[0]
}

func main() {
	flag.Parse()
	word := get_word()
	entries := map[string]bool{}
	placeholder := []string{}
	// slice := make([]string, len(word), len(word))

	i := 0
	for i = 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	chances := 8

	// t := time.NewTimer(2 * time.Minute)
	t := time.NewTimer(5 * time.Second)

	result := make(chan bool)
	go func(chances int) {
		for {
			userInput := strings.Join(placeholder, "")

			if chances == 0 && userInput != word {
				fmt.Printf("%s \n", word)
				result <- false
				fmt.Println("Game Over! Try again")
				return
			}
			if userInput == word {
				result <- true
				fmt.Println("You Win!!")
				return
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
					result <- true
					fmt.Println("You win!")
					return
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
	}(chances)

	for {
		select {
		case <-result:
			// if r {
			// 	// win
			// } else {
			// 	// lose
			// }
			fmt.Println("Congratulations!!")
			goto END
		case <-t.C:
			fmt.Println("Timeout....., Too bad!")
			goto END
		}
	}
END:
	fmt.Println("Let's try again")
}

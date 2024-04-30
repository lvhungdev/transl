package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("provide a word to continue")
		return
	}

	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + args[0]

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.StatusCode != 200 {
        // TODO maybe integrate with spell checker api to determine if there is typo
		fmt.Println("sorry, no definition found")
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var translations []Translation
	err = json.Unmarshal(data, &translations)
	if err != nil {
		panic(err)
	}

	for _, t := range translations {
		fmt.Println(t)
	}
}

type Translation struct {
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
		} `json:"definitions"`
	} `json:"meanings"`
}

func (r Translation) String() string {
	var result string

	for _, m := range r.Meanings {
		if len(m.Definitions) == 0 {
			continue
		}

		result += "usage:      " + m.PartOfSpeech + "\n"
		result += "definition: " + m.Definitions[0].Definition + "\n"
		result += "example:    " + m.Definitions[0].Example + "\n"
	}

	return result
}

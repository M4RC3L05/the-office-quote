package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed quotes.json
var quotesData []byte

type Quotes []struct {
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func main() {
	var quotes Quotes

	json.Unmarshal(quotesData, &quotes)

	quote := quotes[rand.Intn(len(quotes))]

	text := ""

	for _, fragment := range strings.Split(quote.Quote, ".") {
		if len(fragment) <= 0 {
			continue
		}

		text += fmt.Sprintf("%s.\n", strings.Trim(fragment, " "))
	}

	fmt.Printf("%s\n- %s\n", text, quote.Character)
}

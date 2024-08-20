package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
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

	re := regexp.MustCompile(`[.!?]\s+`)
	sentences := strings.Join(re.Split(quote.Quote, -1), ".\n")

	fmt.Printf("%s\n\n- %s\n", sentences, quote.Character)
}

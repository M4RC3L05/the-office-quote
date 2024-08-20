package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

//go:embed quotes.json
var quotesData []byte
var version = "dev-build"

type Quotes []struct {
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func main() {
	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "Show version")

	flag.Parse()

	if versionFlag {
		fmt.Println(version)

		return
	}

	var quotes Quotes

	json.Unmarshal(quotesData, &quotes)

	quote := quotes[rand.Intn(len(quotes))]

	re := regexp.MustCompile(`[.!?]\s+`)
	sentences := strings.Join(re.Split(quote.Quote, -1), ".\n")

	fmt.Printf("%s\n\n- %s\n", sentences, quote.Character)
}

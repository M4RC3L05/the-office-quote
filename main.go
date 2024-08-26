package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"

	"gopkg.in/yaml.v3"
)

//go:embed data/quotes.yaml
var quotesData []byte
var version = "dev-build"

type Quotes []struct {
	Text string `yaml:"text"`
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

	if err := yaml.Unmarshal(quotesData, &quotes); err != nil {
		panic(err)
	}

	quote := quotes[rand.Intn(len(quotes))]

	fmt.Printf("%s\n", quote.Text)
}

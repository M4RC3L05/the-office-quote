package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/goccy/go-yaml"
)

//go:embed data/quotes.yaml
var quotes []byte

var Version string = "v1.1.0"

type Quote struct {
	Text string
}

func main() {
	flag.CommandLine.SetOutput(os.Stdout)

	helpFlag := flag.Bool("help", false, "Show help menu")
	versionFlag := flag.Bool("version", false, "Show version")

	flag.Usage = func() {
		fmt.Printf("The Office Quote\n\n")
		fmt.Printf("Usage %s [OPTIONS]\n\n", os.Args[0])
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}

	if *versionFlag {
		println(Version)
		return
	}

	var data []Quote

	if err := yaml.Unmarshal(quotes, &data); err != nil {
		panic(err)
	}

	println(data[rand.Intn(len(data))].Text)
}

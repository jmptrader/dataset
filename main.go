package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/alexandercampbell/dataset/statlib"
)

func main() {
	width := flag.Int("width", 80, "Output wrap width")
	flag.Parse()

	args := flag.Args()

	var input io.Reader
	var err error

	switch len(args) {
	case 0:
		input = os.Stdin
	case 1:
		if args[0] == "-" {
			input = os.Stdin
		} else if input, err = os.Open(args[0]); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("dataset: too many arguments")
	}

	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	dataset, err := statlib.ReadDataset(bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(dataset.Histogram(*width))
	fmt.Println()
	fmt.Println(dataset.FiveNumberSummary().AsTable(*width))
}

package main

import (
	"fmt"

	"github.com/FlamesX-128/monas-chinas-cli/src/controllers/questions"
	"github.com/FlamesX-128/monas-chinas-cli/src/solvers"
)

func main() {
	var options = []string{}

	for key := range solvers.Translations {
		options = append(options, key)

	}

	language := questions.Selection("Select the preferred language.", options)

	fmt.Println(language)
}

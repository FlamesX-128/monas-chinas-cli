package tools

import (
	"fmt"
	"log"
	"os"

	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func ErrIfEmpty[T any](arr []T, em string) []T {
	if len(arr) == 0 {
		fmt.Println(em)

		os.Exit(1)
	}

	return arr
}

func NewScraper[T types.Anime | types.Service | string](
	url string, query string, fn func(h *colly.HTMLElement) T,
) (res []T) {
	c := colly.NewCollector()

	c.OnHTML(query, func(h *colly.HTMLElement) {
		res = append(res, fn(h))

	})

	if err := c.Visit(url); err != nil {
		log.Panicln(err.Error())

	}

	return
}

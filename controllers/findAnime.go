package controllers

import (
	"log"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/utils"
	"github.com/gocolly/colly"
)

func FindAnime(toSearch string) []utils.Anime {
	var animeList []utils.Anime

	c := colly.NewCollector()

	c.OnHTML("div.row a[href]", func(h *colly.HTMLElement) {
		name := h.ChildText("h5.seristitles")
		url := h.Attr("href")

		animeList = append(animeList, utils.Anime{Name: name, URL: url})
	})

	err := c.Visit(utils.BaseURL + "buscar?q=" + (strings.Join(strings.Split(toSearch, " "), "+")))

	if err != nil {
		log.Panicln(err)

	}

	return animeList
}

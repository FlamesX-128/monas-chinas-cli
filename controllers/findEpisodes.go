package controllers

import (
	"log"

	"github.com/gocolly/colly"
)

func FindEpisodes(url string) []string {
	var episodeList []string

	c := colly.NewCollector()

	c.OnHTML("div.allanimes a", func(h *colly.HTMLElement) {
		episodeList = append(episodeList, h.Attr("href"))

	})

	err := c.Visit(url)

	if err != nil {
		log.Panicln(err)

	}

	return episodeList
}

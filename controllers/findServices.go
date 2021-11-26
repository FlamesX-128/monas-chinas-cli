package controllers

import (
	"log"

	"github.com/gocolly/colly"
)

func FindServices(url string) []string {
	var servicesList []string

	c := colly.NewCollector()

	c.OnHTML("div.heromain div.playother p", func(h *colly.HTMLElement) {
		servicesList = append(servicesList, h.Text)

	})

	err := c.Visit(url)

	if err != nil {
		log.Panicln(err)

	}

	return servicesList
}

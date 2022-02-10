package tools

import (
	"log"

	"github.com/FlamesX-128/monas-chinas-cli/src/interfaces"
	"github.com/gocolly/colly"
)

func FindServices(url string) (services []interfaces.Anime) {
	c := colly.NewCollector()

	c.OnHTML("div.heromain div.playother p", func(h *colly.HTMLElement) {
		name := h.Text
		url := h.Attr("data-player")

		services = append(services, interfaces.Anime{Name: name, Url: url})
	})

	err := c.Visit(url)

	if err != nil {
		log.Panicln(err)

	}

	return
}

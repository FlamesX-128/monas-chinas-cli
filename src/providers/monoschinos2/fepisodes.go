package monoschinos2

import (
	"log"

	"github.com/gocolly/colly"
)

func FetchEpisodes(url string) (episodes []string) {
	c := colly.NewCollector()

	c.OnHTML("div.allanimes a", func(h *colly.HTMLElement) {
		episodes = append(episodes, h.Attr("href"))

	})

	if err := c.Visit(url); err != nil {
		log.Panicln(err.Error())

	}

	return
}

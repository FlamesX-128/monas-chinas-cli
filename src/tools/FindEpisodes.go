package tools

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func FindEpisodes(url string) (episodes []string) {
	c := colly.NewCollector()

	c.OnHTML("div.allanimes a", func(h *colly.HTMLElement) {
		episodes = append(episodes, h.Attr("href"))

	})

	err := c.Visit(url)

	if err != nil {
		fmt.Print(err)

		os.Exit(0)
	}

	return
}

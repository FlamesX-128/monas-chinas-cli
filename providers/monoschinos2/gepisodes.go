package monoschinos2

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
	"github.com/gocolly/colly"
)

func GetEpisodes(u string) []string {
	q := "div.allanimes a"

	f := func(h *colly.HTMLElement) string {

		return h.Attr("href")
	}

	return tools.NewScraper(u, q, f)
}

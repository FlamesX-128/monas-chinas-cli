package gogoanime

import (
	"log"

	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func FetchAnimes(name string) (animes []types.Anime) {
	c := colly.NewCollector()

	c.OnHTML("div.last_episodes ul.items li", func(h *colly.HTMLElement) {
		animes = append(animes, types.Anime{
			Image: h.ChildAttr("img", "src"),
			Name:  h.ChildAttr("a", "title"),
			Url:   h.ChildAttr("a", "href"),
		})

	})

	if err := c.Visit(base_url + "/search.html?keyword=" + name); err != nil {
		log.Panicln(err.Error())

	}

	return
}

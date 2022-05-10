package monoschinos2

import (
	"log"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func FetchAnimes(name string) (animes []types.Anime) {
	c := colly.NewCollector()

	name = strings.Join(strings.Split(name, " "), "+")

	c.OnHTML("div.row a[href]", func(h *colly.HTMLElement) {
		animes = append(animes, types.Anime{
			Image: h.ChildAttr("img.animemainimg", "src"),
			Name:  h.ChildText("h5.seristitles"),
			Url:   h.Attr("href"),
		})

	})

	if err := c.Visit(url + "buscar?q=" + name); err != nil {
		log.Panicln(err.Error())

	}

	return
}

package providers

import (
	"log"

	"github.com/FlamesX-128/monas-chinas-cli/src/controllers"
	"github.com/gocolly/colly"
)

const monos_url = "https://monoschinos2.com/"

var Monoschinos = &controllers.Service{
	SearchAnimes: func(name string) (animes []controllers.Anime) {
		c := colly.NewCollector()

		c.OnHTML("div.row a[href]", func(h *colly.HTMLElement) {
			animes = append(animes, controllers.Anime{
				Image: h.ChildAttr("img.animemainimg", "src"),
				Name:  h.ChildText("h5.seristitles"),
				Url:   h.Attr("href"),
			})

		})

		if err := c.Visit(monos_url + "buscar?q=" + name); err != nil {
			log.Panicln(err.Error())

		}

		return
	},
}

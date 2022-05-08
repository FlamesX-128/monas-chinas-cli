package providers

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/controllers"
	"github.com/gocolly/colly"
)

const monos_url = "https://monoschinos2.com/"

var Monoschinos = &controllers.Provider{
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
	SearchEpisodes: func(url string) (episodes []string) {
		c := colly.NewCollector()

		c.OnHTML("div.allanimes a", func(h *colly.HTMLElement) {
			episodes = append(episodes, h.Attr("href"))

		})

		if err := c.Visit(url); err != nil {
			log.Panicln(err.Error())

		}

		return
	},
	SearchServices: func(url string) (services []controllers.Service) {
		c := colly.NewCollector()

		c.OnHTML("div.heromain div.playother p", func(h *colly.HTMLElement) {
			var url string = ""

			link, _ := base64.StdEncoding.DecodeString(h.Attr("data-player"))
			urls := strings.Split(string(link), "?url=")

			if len(urls) == 2 {
				url = urls[1]

			} else {
				url = urls[0]

			}

			services = append(services, controllers.Service{
				Name: h.Text,
				Url:  url,
			})
		})

		if err := c.Visit(url); err != nil {
			log.Panicln(err)

		}

		return
	},
}

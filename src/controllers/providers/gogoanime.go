package providers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/controllers"
	"github.com/gocolly/colly"
)

const gogo_url = "https://gogoanime.sk/"

var Gogoanime = &controllers.Provider{
	SearchAnimes: func(name string) (animes []controllers.Anime) {
		c := colly.NewCollector()

		c.OnHTML("div.last_episodes ul.items li", func(h *colly.HTMLElement) {
			animes = append(animes, controllers.Anime{
				Image: h.ChildAttr("img", "src"),
				Name:  h.ChildAttr("a", "title"),
				Url:   h.ChildAttr("a", "href"),
			})

		})

		if err := c.Visit(gogo_url + "/search.html?keyword=" + name); err != nil {
			log.Panicln(err.Error())

		}

		return
	},
	SearchEpisodes: func(url string) (episodes []string) {
		var amount int = 1

		c := colly.NewCollector()

		c.OnHTML("ul#episode_page a[href]", func(h *colly.HTMLElement) {
			i, err := strconv.Atoi(h.Attr("ep_end"))

			if err != nil {
				log.Panicln(err.Error())

			}

			amount = i
		})

		if err := c.Visit(gogo_url + url); err != nil {
			log.Panicln(err.Error())

		}

		for i := 1; i <= amount; i++ {
			episodes = append(episodes, gogo_url+url[strings.LastIndex(url, "/")+1:]+"-episode-"+strconv.Itoa(i))

		}

		return
	},
	SearchServices: func(url string) (services []controllers.Service) {
		c := colly.NewCollector()

		fmt.Println(url)

		c.OnHTML("div.anime_muti_link ul li", func(h *colly.HTMLElement) {
			name := h.Attr("class")
			url := h.ChildAttr("li[class] a", "data-video")

			if url[0:2] == "//" {
				url = "https:" + url

			}

			services = append(services, controllers.Service{Name: name, Url: url})
		})

		if err := c.Visit(url); err != nil {
			log.Panicln(err)

		}

		return
	},
}

package gogoanime

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func FetchEpisodes(url string) (episodes []string) {
	var (
		amount int
		err    error
	)

	c := colly.NewCollector()

	c.OnHTML("ul#episode_page a[href]", func(h *colly.HTMLElement) {
		amount, err = strconv.Atoi(h.Attr("ep_end"))

		if err != nil {
			log.Panicln(err.Error())

		}

	})

	if err := c.Visit(base_url + url); err != nil {
		log.Panicln(err.Error())

	}

	for i := 1; i <= amount; i++ {
		link := base_url + url[strings.LastIndex(url, "/")+1:] +
			"-episode-" + strconv.Itoa(i)

		episodes = append(episodes, link)

	}

	return
}

package monoschinos2

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func FetchServices(url string) (services []types.Service) {
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

		services = append(services, types.Service{
			Name: h.Text,
			Url:  url,
		})
	})

	if err := c.Visit(url); err != nil {
		log.Panicln(err)

	}

	return
}

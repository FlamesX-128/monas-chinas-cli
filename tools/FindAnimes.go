package tools

import (
	"fmt"
	"os"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/interfaces"
	"github.com/gocolly/colly"
)

func FindAnime(name string) (animes []interfaces.Anime) {
	c := colly.NewCollector()

	c.OnHTML("div.row a[href]", func(h *colly.HTMLElement) {
		image := h.ChildAttr("img.animemainimg", "src")
		name := h.ChildText("h5.seristitles")
		url := h.Attr("href")

		animes = append(animes, interfaces.Anime{Name: name, Url: url, Image: image})
	})

	err := c.Visit(interfaces.BaseURL + "buscar?q=" + (strings.Join(strings.Split(name, " "), "+")))

	if err != nil {
		fmt.Print(err)

		os.Exit(0)
	}

	return
}

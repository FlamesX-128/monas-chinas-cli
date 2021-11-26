package controllers

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type serviceHandler struct {
	Data string
	ID   string
	Tag  string
}

func GetPage(url string) []serviceHandler {
	var handlers []serviceHandler

	c := colly.NewCollector()

	c.OnHTML("div.playother p.play-video", func(h *colly.HTMLElement) {
		data := h.Attr("data-player")

		handlers = append(handlers, serviceHandler{
			Data: data,
			ID:   h.Text,
			Tag: fmt.Sprintf(`
				<p class="play-video" data-player="%s">
					%s
				</p>
			`, data, h.Text),
		})
	})

	err := c.Visit(url)

	if err != nil {
		log.Panicln(err)

	}

	return handlers
}

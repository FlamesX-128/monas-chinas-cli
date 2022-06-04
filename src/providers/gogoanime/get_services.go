package gogoanime

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func (p *provider) GetServices(u string) []types.Service {
	q := "div.anime_muti_link ul li"

	f := func(h *colly.HTMLElement) types.Service {
		u := h.ChildAttr("li[class] a", "data-video")

		if u[0] == '/' {
			u = "https:" + u

		}

		return types.Service{
			Name: h.Attr("class"),
			Url:  u,
		}

	}

	return tools.NewScraper(u, q, f)
}

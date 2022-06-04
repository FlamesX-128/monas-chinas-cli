package monoschinos

import (
	"encoding/base64"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func (p *provider) GetServices(u string) []types.Service {
	q := "div.heromain div.playother p"

	f := func(h *colly.HTMLElement) types.Service {
		url := h.Attr("data-player")

		u, _ := base64.StdEncoding.DecodeString(url)

		if u := strings.Split(string(u), "?url="); len(u) == 2 {
			url = u[1]

		}

		return types.Service{
			Name: h.Text,
			Url:  url,
		}

	}

	return tools.NewScraper(u, q, f)
}

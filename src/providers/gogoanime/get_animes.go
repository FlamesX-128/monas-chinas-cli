package gogoanime

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func (p *provider) GetAnimes(kw string) []types.Anime {
	u := url + "/search.html?keyword=" + kw
	q := "div.last_episodes ul.items li"

	f := func(h *colly.HTMLElement) types.Anime {
		return types.Anime{
			Image: h.ChildAttr("img", "src"),
			Name:  h.ChildAttr("a", "title"),
			Url:   h.ChildAttr("a", "href"),
		}

	}

	return tools.NewScraper(u, q, f)
}

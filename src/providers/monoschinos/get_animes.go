package monoschinos

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
	"github.com/gocolly/colly"
)

func (p *provider) GetAnimes(kw string) []types.Anime {
	u := url + "buscar?q=" + kw
	q := "div.row a[href]"

	f := func(h *colly.HTMLElement) types.Anime {
		return types.Anime{
			Image: h.ChildAttr("img.animemainimg", "src"),
			Name:  h.ChildText("h5.seristitles"),
			Url:   h.Attr("href"),
		}

	}

	return tools.NewScraper(u, q, f)
}

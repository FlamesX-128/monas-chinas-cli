package main

import (
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/providers"
	"github.com/FlamesX-128/monas-chinas-cli/src/controllers/survey"
	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
)

func ProviderToUse() types.Provider {
	return providers.Handler[survey.Selection(
		tools.GetKeys(providers.Handler), "What provider do you want to use?",
	)]

}

func AnimeToSearch(m string) string {
	return strings.ReplaceAll(survey.String(m), " ", "+")

}

func AnimeToWatch(p types.Provider, k string) types.Anime {
	la := p.GetTranslation()

	l := tools.ErrIfEmpty(p.GetAnimes(k), la.AnimesNF)
	o := tools.FindFields[string](l, "Name")

	return l[tools.FindIndex(
		o, survey.Selection(o, la.AnimeTW),
	)]

}

func EpisodeToWatch(p types.Provider, u string) string {
	la := p.GetTranslation()

	l := tools.ErrIfEmpty(p.GetEpisodes(u), la.EpisodesNF)

	return l[survey.Number(
		la.EpisodeTW, la.InvalidNS, len(l),
	)]

}

func ServiceToUse(p types.Provider, u string) types.Service {
	la := p.GetTranslation()

	l := tools.ErrIfEmpty(p.GetServices(u), la.ServicesNF)
	o := tools.FindFields[string](l, "Name")

	return l[tools.FindIndex(
		o, survey.Selection(o, la.ServiceTU),
	)]

}

func main() {

}

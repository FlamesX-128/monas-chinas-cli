package main

import (
	"fmt"

	"github.com/FlamesX-128/monas-chinas-cli/src/controllers/questions"
	"github.com/FlamesX-128/monas-chinas-cli/src/solvers"
	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
)

func main() {
	var options = []string{}

	for key := range solvers.Translations {
		options = append(options, key)

	}

	language := questions.Selection("Select the preferred language.", options)

	translation := solvers.Translations[language]
	provider := tools.GetProvider(language)

	anime_list := provider.SearchAnimes(questions.String(translation.AnimeToSearch))
	options = []string{}

	for _, anime := range anime_list {
		options = append(options, anime.Name)

	}

	anime := anime_list[tools.FindIndex(options, questions.Selection(translation.AnimeToWatch, options))]

	episode_list := provider.SearchEpisodes(anime.Url)
	episode := episode_list[questions.Number(translation.EpisodeToWatch, "", len(episode_list))-1]

	service_list := provider.SearchServices(episode)

	fmt.Println(service_list)
}

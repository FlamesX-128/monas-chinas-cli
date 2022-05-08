package main

import (
	"fmt"

	"github.com/FlamesX-128/monas-chinas-cli/src/controllers"
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

	if len(options) == 0 {
		fmt.Println(translation.AnimesNotFound)

		return
	}

	anime := anime_list[tools.FindIndex(options, questions.Selection(translation.AnimeToWatch, options))]

	episode_list := provider.SearchEpisodes(anime.Url)

	if len(episode_list) == 0 {
		fmt.Println(translation.EpisodesNotFound)

		return
	}

	episode_number := questions.Number(translation.EpisodeToWatch, "", len(episode_list)) - 1
	episode := episode_list[episode_number]

	service_list := provider.SearchServices(episode)
	options = []string{}

	for _, service := range service_list {
		options = append(options, service.Name)

	}

	if len(options) == 0 {
		fmt.Println(translation.EpisodesNotFound)

		return
	}

	service := service_list[tools.FindIndex(options, questions.Selection(translation.ServiceToUse, options))]

	go controllers.Presence(anime.Name, anime.Image, episode_number)
	controllers.WebView(service.Url)
}

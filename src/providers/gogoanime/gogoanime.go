package gogoanime

import "github.com/FlamesX-128/monas-chinas-cli/src/types"

const base_url string = "https://gogoanime.sk/"

var Handler = types.Provider{
	Translation: types.Translation{
		AnimeToSearch:  "What anime do you want to see?",
		AnimeToWatch:   "The following animes were found:",
		AnimesNotFound: "No anime found.",

		EpisodeToWatch:   "Enter a number between 1 to %v to select the episode:",
		EpisodesNotFound: "No episodes found.",

		ServiceToUse:     "The following services were found:",
		ServicesNotFound: "No services found.",
	},
	Fetch: types.Fetch{
		Animes:   FetchAnimes,
		Episodes: FetchEpisodes,
		Services: FetchServices,
	},
}

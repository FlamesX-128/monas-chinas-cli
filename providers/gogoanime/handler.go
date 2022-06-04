package gogoanime

import "github.com/FlamesX-128/monas-chinas-cli/src/types"

const Url string = "https://gogoanime.sk/"

var Handler = types.Provider{
	Transl: types.Transl{
		AnimeTS:  "What anime do you want to see?",
		AnimeTW:  "The following animes were found:",
		AnimesNF: "No animes found.",

		EpisodeTW:  "Enter a number between 1 to %v:",
		EpisodesNF: "No episodes found.",

		InvalidNS: "You must enter a number between 1 to %v!",

		ServiceTU:  "The following services were found:",
		ServicesNF: "No services found.",

		QualityTU: "The following resolutions were found:",
	},
	Funcs: types.Funcs{
		GetAnimes:   GetAnimes,
		GetEpisodes: GetEpisodes,
		GetServices: GetServices,
	},
}

package monoschinos2

import "github.com/FlamesX-128/monas-chinas-cli/src/types"

const url string = "https://monoschinos2.com/"

var Handler = types.Provider{
	Translation: types.Translation{
		AnimeToSearch:  "¿Qué anime quieres ver?",
		AnimeToWatch:   "Se encontraron los siguientes animes:",
		AnimesNotFound: "No se encontraron animes.",

		EpisodeToWatch:   "Introduzca un numero entre 1 a %v para seleccionar el episodio:",
		EpisodesNotFound: "No se encontraron episodios",

		ServiceToUse:     "Se encontraron los siguientes servicios:",
		ServicesNotFound: "No se encontraron servicios.",
	},
	Fetch: types.Fetch{
		Animes:   FetchAnimes,
		Episodes: FetchEpisodes,
		Services: FetchServices,
	},
}

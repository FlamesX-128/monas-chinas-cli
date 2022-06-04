package monoschinos2

import "github.com/FlamesX-128/monas-chinas-cli/src/types"

const Url string = "https://monoschinos2.com/"

var Handler = types.Provider{
	Transl: types.Transl{
		AnimeTS:  "¿Qué anime quieres ver?",
		AnimeTW:  "Se encontraron los siguientes animes:",
		AnimesNF: "No se encontraron animes.",

		EpisodeTW:  "Introduzca un número entre 1 y %v:",
		EpisodesNF: "No se encontraron episodios.",

		InvalidNS: "Debes introducir un numero entre 1 a %v!",

		ServiceTU:  "¿Qué servicio quieres ver?",
		ServicesNF: "No se encontraron servicios.",

		QualityTU: "Se encontraron las siguientes resoluciones:",
	},
	Funcs: types.Funcs{
		GetAnimes:   GetAnimes,
		GetEpisodes: GetEpisodes,
		GetServices: GetServices,
	},
}

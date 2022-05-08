package solvers

type Translation struct {
	AnimeToSearch  string
	AnimeToWatch   string
	AnimesNotFound string

	EpisodeToWatch   string
	EpisodesNotFound string

	ServiceToUse     string
	ServicesNotFound string
}

var Translations = map[string]Translation{
	"english": {
		AnimeToSearch:  "What anime do you want to see?",
		AnimeToWatch:   "The following animes were found:",
		AnimesNotFound: "No anime found.",

		EpisodeToWatch:   "Enter a number between 1 to %v to select the episode:",
		EpisodesNotFound: "No episodes found.",

		ServiceToUse:     "The following services were found:",
		ServicesNotFound: "No services found.",
	},
	"español": {
		AnimeToSearch:  "¿Qué anime quieres ver?",
		AnimeToWatch:   "Se encontraron los siguientes animes:",
		AnimesNotFound: "No se encontraron animes.",

		EpisodeToWatch:   "Introduzca un numero entre 1 a %v para seleccionar el episodio:",
		EpisodesNotFound: "No se encontraron episodios",

		ServiceToUse:     "Se encontraron los siguientes servicios:",
		ServicesNotFound: "No se encontraron servicios.",
	},
}

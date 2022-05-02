package solvers

type Translation struct {
	AnimeToSearch  string
	AnimeToWatch   string
	EpisodeToWatch string
}

var Translations = map[string]Translation{
	"english": {
		AnimeToSearch:  "What anime do you want to see?",
		AnimeToWatch:   "The following animes were found:",
		EpisodeToWatch: "Enter a number between 1 to %v to select the episode:",
	},
	"español": {
		AnimeToSearch:  "¿Qué anime quieres ver?",
		AnimeToWatch:   "Se encontraron los siguientes animes:",
		EpisodeToWatch: "introduzca un numero entre 1 a %v para seleccionar el episodio:",
	},
}

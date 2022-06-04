package types

// Contains provider translations.
type Translation struct {
	AnimeTS  string // Anime to search.
	AnimeTW  string // Anime to watch.
	AnimesNF string // Animes not found.

	EpisodeTW  string // Episode to watch.
	EpisodesNF string // Episodes not found.

	ServiceTU  string // Service to use.
	ServicesNF string // Services not found.

	InvalidNS string // Invalid numeric survey.
	QualityTU string // Quality to use.
}

type Provider interface {
	GetAnimes(string) []Anime
	GetEpisodes(string) []string
	GetServices(string) []Service
	GetTranslation() Translation
}

package types

type Translation struct {
	AnimeToSearch  string
	AnimeToWatch   string
	AnimesNotFound string

	EpisodeToWatch   string
	EpisodesNotFound string

	ServiceToUse     string
	ServicesNotFound string
}

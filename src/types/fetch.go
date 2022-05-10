package types

type Fetch struct {
	Animes func(name string) (animes []Anime)

	Episodes func(url string) (episodes []string)
	Services func(url string) (services []Service)
}

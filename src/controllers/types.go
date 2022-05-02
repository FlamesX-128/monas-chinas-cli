package controllers

type Anime struct {
	Image string
	Name  string
	Url   string
}

type Service struct {
	Name string
	Url  string
}

type Provider struct {
	SearchAnimes   func(name string) (animes []Anime)
	SearchEpisodes func(url string) (episodes []string)
	SearchServices func(url string) (services []Service)
}

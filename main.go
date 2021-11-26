package main

import (
	_ "embed"
	"fmt"

	"github.com/FlamesX-128/monas-chinas-cli/controllers"
	"github.com/FlamesX-128/monas-chinas-cli/handlers"
	"github.com/FlamesX-128/monas-chinas-cli/utils"
	"github.com/wailsapp/wails"
)

func main() {
	toSearch := handlers.ToSearch()
	animesList := controllers.FindAnime(toSearch)

	if len(animesList) == 0 {
		fmt.Println("No results found")

		return
	}

	selectedAnime := handlers.ToSelectAnime(animesList)
	episodesList := controllers.FindEpisodes(selectedAnime)

	episode := handlers.ToSelectEpisode(episodesList)

	servicesList := controllers.FindServices(episode)
	service := handlers.ToSelectService(servicesList)

	serviceHandler := controllers.GetPage(episode)
	page := controllers.BuildPage(utils.HTML, selectedAnime, serviceHandler, service)

	app := wails.CreateApp(&wails.AppConfig{
		Colour:           "#010409",
		DisableInspector: true,
		Height:           480,
		HTML:             page,
		Resizable:        false,
		MinWidth:         640,
		MinHeight:        480,
		MaxWidth:         640,
		MaxHeight:        480,
		Title:            "monas-chinas-cli",
		Width:            640,
	})

	app.Run()
}

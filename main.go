package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gocolly/colly"
	"github.com/webview/webview"
)

const baseURL = "https://monoschinos2.com/"

type anime struct {
	Name string
	Url  string
}

func survey_string(msg string) (input string) {
	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: msg,
			},
		},
	}, &input)

	if err != nil {
		fmt.Print(err)

		os.Exit(0)
	}

	return
}

func survey_number(max int, msg string) (input int) {
	var Input string

	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: msg,
			},
			Validate: func(ans interface{}) error {
				res, err := strconv.Atoi(ans.(string))

				if err != nil {
					return fmt.Errorf("enter a number")

				}

				if res < 1 || res > max {
					return fmt.Errorf("enter a number greater than 1 or less than %d", max)

				}

				return nil
			},
		},
	}, &Input)

	input, _ = strconv.Atoi(Input)

	if err != nil {
		fmt.Print(err)

		os.Exit(0)
	}

	return
}

func find_anime(name string) (animes []anime) {
	c := colly.NewCollector()

	c.OnHTML("div.row a[href]", func(h *colly.HTMLElement) {
		name := h.ChildText("h5.seristitles")
		url := h.Attr("href")

		animes = append(animes, anime{Name: name, Url: url})
	})

	err := c.Visit(baseURL + "buscar?q=" + (strings.Join(strings.Split(name, " "), "+")))

	if err != nil {
		fmt.Print(err)

		os.Exit(0)
	}

	return
}

func find_episodes(url string) (episodes []string) {
	c := colly.NewCollector()

	c.OnHTML("div.allanimes a", func(h *colly.HTMLElement) {
		episodes = append(episodes, h.Attr("href"))

	})

	err := c.Visit(url)

	if err != nil {
		fmt.Print(err)

		os.Exit(0)
	}

	return
}

func find_services(url string) (services []anime) {
	c := colly.NewCollector()

	c.OnHTML("div.heromain div.playother p", func(h *colly.HTMLElement) {
		name := h.Text
		url := h.Attr("data-player")

		services = append(services, anime{Name: name, Url: url})
	})

	err := c.Visit(url)

	if err != nil {
		log.Panicln(err)

	}

	return
}

func main() {
	animeName := survey_string("¿Qué anime quieres ver?")
	animeList := find_anime(animeName)

	if len(animeList) == 0 {
		fmt.Println("No se encontró el anime que buscabas.")

		os.Exit(1)
	}

	for i, anime := range animeList {
		log.Printf("[%d]: %s", i+1, anime.Name)

	}

	animeTarget := survey_number(len(animeList), "Se encontraron los siguientes animes, ¿Cual quieres ver?")
	episodeList := find_episodes(animeList[animeTarget-1].Url)

	if len(episodeList) == 0 {
		fmt.Println("No se encontraron episodios.")

		os.Exit(1)
	}

	for i, episode := range episodeList {
		log.Printf("[%d]: %s", i+1, episode[strings.Index(episode, "ver/")+4:])

	}

	episodeTarget := survey_number(len(episodeList), "Se encontraron los siguientes episodios, ¿Cual quieres ver?")
	serviceList := find_services(episodeList[episodeTarget-1])

	for i, service := range serviceList {
		log.Printf("[%d]: %s", i+1, service.Name)

	}

	serviceTarget := survey_number(len(serviceList), "Se encontraron los siguientes servicios, ¿Cual quieres usar?")

	link, _ := base64.StdEncoding.DecodeString(serviceList[serviceTarget-1].Url)

	urls := strings.Split(string(link), "?url=")

	var url string

	if len(urls) == 2 {
		url = urls[1]

	} else {
		url = urls[0]

	}

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("monas-chinas-cli")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(url)
	w.Run()
}

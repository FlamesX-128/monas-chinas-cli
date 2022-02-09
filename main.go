/*

	MIT License

	Copyright (c) 2021 FlamesX128

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/controllers"
	"github.com/FlamesX-128/monas-chinas-cli/tools"
)

func main() {
	animeName := controllers.StringSurvey("¿Qué anime quieres ver?")
	animeList := tools.FindAnime(animeName)

	if len(animeList) == 0 {
		fmt.Println("No se encontró el anime que buscabas.")

		os.Exit(1)
	}

	for i, anime := range animeList {
		log.Printf("[%d]: %s", i+1, anime.Name)

	}

	animeTarget := controllers.NumberSurvey(len(animeList), "Se encontraron los siguientes animes, ¿Cual quieres ver?")
	episodeList := tools.FindEpisodes(animeList[animeTarget-1].Url)

	if len(episodeList) == 0 {
		fmt.Println("No se encontraron episodios.")

		os.Exit(1)
	}

	for i, episode := range episodeList {
		log.Printf("[%d]: %s", i+1, episode[strings.Index(episode, "ver/")+4:])

	}

	episodeTarget := controllers.NumberSurvey(len(episodeList), "Se encontraron los siguientes episodios, ¿Cual quieres ver?")
	serviceList := tools.FindServices(episodeList[episodeTarget-1])

	for i, service := range serviceList {
		log.Printf("[%d]: %s", i+1, service.Name)

	}

	serviceTarget := controllers.NumberSurvey(len(serviceList), "Se encontraron los siguientes servicios, ¿Cual quieres usar?")

	link, _ := base64.StdEncoding.DecodeString(serviceList[serviceTarget-1].Url)
	urls := strings.Split(string(link), "?url=")

	go controllers.RichPresence(animeList[animeTarget-1].Name, animeList[animeTarget-1].Image, episodeTarget)
	controllers.WebView(urls)
}

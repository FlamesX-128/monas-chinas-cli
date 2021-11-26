package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/FlamesX-128/monas-chinas-cli/utils"
	"github.com/manifoldco/promptui"
)

func ToSelectAnime(animeList []utils.Anime) string {
	for i, anime := range animeList {
		log.Printf("[%d]: %s", i+1, anime.Name)

	}

	prompt := promptui.Prompt{
		Label: "Enter a number to select the anime",
		Validate: func(input string) error {
			res, err := strconv.Atoi(input)

			if err != nil {
				return fmt.Errorf("enter a number")

			}

			if res < 1 || res > len(animeList) {
				return fmt.Errorf("enter a number greater than 1 or less than %d", len(animeList))

			}

			return nil
		},
	}

	input, _ := prompt.Run()
	res, _ := strconv.Atoi(input)

	return animeList[(res - 1)].URL
}

package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func ToSelectEpisode(episodeList []string) string {
	for i, episode := range episodeList {
		log.Printf("[%d]: %s", i+1, episode[strings.Index(episode, "ver/")+4:])

	}

	prompt := promptui.Prompt{
		Label: "Enter a number to select the anime",
		Validate: func(input string) error {
			res, err := strconv.Atoi(input)

			if err != nil {
				return fmt.Errorf("enter a number")

			}

			if res < 1 || res > len(episodeList) {
				return fmt.Errorf(
					"enter a number greater than 1 or less than %d", len(episodeList),
				)

			}

			return nil
		},
	}

	res, _ := prompt.Run()
	resInt, _ := strconv.Atoi(res)

	return episodeList[(resInt - 1)]
}

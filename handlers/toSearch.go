package handlers

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func ToSearch() string {
	prompt := promptui.Prompt{
		Label: "Enter the name of the anime to search",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("you need to enter a name")

			}

			return nil
		},
	}

	result, _ := prompt.Run()

	return result
}

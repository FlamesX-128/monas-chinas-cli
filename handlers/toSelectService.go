package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

func ToSelectService(servicesList []string) string {
	for i, service := range servicesList {
		log.Printf("[%d]: %s", i+1, service)

	}

	prompt := promptui.Prompt{
		Label: "Enter a number to select the service",
		Validate: func(input string) error {
			res, err := strconv.Atoi(input)

			if err != nil {
				return fmt.Errorf("enter a number")

			}

			if res < 1 || res > len(servicesList) {
				return fmt.Errorf(
					"enter a number greater than 1 or less than %d", len(servicesList),
				)

			}

			return nil
		},
	}

	res, _ := prompt.Run()
	resInt, _ := strconv.Atoi(res)

	return servicesList[(resInt - 1)]
}

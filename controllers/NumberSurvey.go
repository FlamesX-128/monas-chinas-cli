package controllers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func NumberSurvey(max int, msg string) (input int) {
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

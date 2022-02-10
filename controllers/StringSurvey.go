package controllers

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func StringSurvey(msg string) (input string) {
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

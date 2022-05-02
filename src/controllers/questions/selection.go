package questions

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func Selection(message string, options []string) (input string) {
	prompt := survey.Select{
		Message: message,
		Options: options,
	}

	if err := survey.AskOne(&prompt, &input); err != nil {
		log.Panicln(err.Error())

	}

	return
}

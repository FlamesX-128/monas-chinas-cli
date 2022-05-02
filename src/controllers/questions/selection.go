package questions

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func Selection(message string, options []string) (language string) {
	prompt := survey.Select{
		Message: message,
		Options: options,
	}

	if err := survey.AskOne(&prompt, &language); err != nil {
		log.Panicln(err.Error())

	}

	return
}

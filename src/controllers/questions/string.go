package questions

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func String(message string) (input string) {
	prompt := survey.Input{
		Message: message,
	}

	if err := survey.AskOne(&prompt, &input); err != nil {
		log.Panicln(err.Error())

	}

	return
}

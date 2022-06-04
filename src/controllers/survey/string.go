package survey

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func String(msg string) (res string) {
	prompt := survey.Input{
		Message: msg,
	}

	if err := survey.AskOne(&prompt, &res); err != nil {
		log.Panicln(err.Error())

	}

	return
}

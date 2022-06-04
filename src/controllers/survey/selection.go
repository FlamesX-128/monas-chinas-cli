package survey

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func Selection(opts []string, msg string) (res string) {
	prompt := survey.Select{
		Message: msg,
		Options: opts,
	}

	if err := survey.AskOne(&prompt, &res); err != nil {
		log.Panicln(err.Error())

	}

	return
}

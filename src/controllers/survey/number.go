package survey

import (
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func Number(m, em string, max int) (resp int) {
	var (
		ign string
		err error
	)

	prompt := []*survey.Question{
		{
			Prompt: &survey.Input{
				Message: fmt.Sprintf(m, max),
			},

			Validate: func(ans interface{}) error {
				resp, err = strconv.Atoi(ans.(string))

				if err == nil && (resp > 0 && resp <= max) {

					return nil
				}

				return fmt.Errorf(m, max)
			},
		},
	}

	if err := survey.Ask(prompt, &ign); err != nil {
		log.Panicln(err.Error())

	}

	return resp - 1
}

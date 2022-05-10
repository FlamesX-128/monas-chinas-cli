package questions

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func Number(m string, em string, max int) (input string) {
	prompt := []*survey.Question{
		{
			Prompt: &survey.Input{
				Message: fmt.Sprintf(m, max),
			},
			Validate: func(ans interface{}) error {
				number, err := strconv.Atoi(ans.(string))

				if err == nil && (number > 0 && number <= max) {

					return nil
				}

				return errors.New(em + strconv.Itoa(max))
			},
		},
	}

	if err := survey.Ask(prompt, &input); err != nil {
		log.Panicln(err.Error())

	}

	return
}

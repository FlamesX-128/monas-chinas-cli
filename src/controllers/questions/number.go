package questions

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func Number(message string, err_msg string, max int) (input int) {
	var (
		ignore string
		err    error
	)

	prompt := []*survey.Question{
		{
			Prompt: &survey.Input{
				Message: fmt.Sprintf(message, max),
			},
			Validate: func(ans interface{}) error {
				input, err = strconv.Atoi(ans.(string))

				if err == nil && (input > 0 && input <= max) {

					return nil
				}

				return errors.New(err_msg + strconv.Itoa(max))
			},
		},
	}

	if err = survey.Ask(prompt, &ignore); err != nil {
		log.Panicln(err.Error())

	}

	fmt.Println(input)

	return
}

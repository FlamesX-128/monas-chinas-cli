package services

import (
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
)

var handler = map[string]func(string) string{
	"fembed": Fembed,
	"ok.ru": func(s string) string {
		return s
	},
}

func Handler(u string) string {
	for _, k := range tools.GetKeys(handler) {
		if strings.Contains(strings.ToLower(u), k) {

			return handler[k](u)
		}

	}

	return u
}

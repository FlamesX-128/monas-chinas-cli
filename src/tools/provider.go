package tools

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/controllers"
	"github.com/FlamesX-128/monas-chinas-cli/src/controllers/providers"
)

func GetProvider(language string) controllers.Provider {
	switch language {
	case "español":
		return *providers.Monoschinos

	}

	return *providers.Gogoanime
}

package providers

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/providers/gogoanime"
	"github.com/FlamesX-128/monas-chinas-cli/src/providers/monoschinos"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
)

var Handler = map[string]types.Provider{
	"monoschinos2": types.Provider(monoschinos.Handler),
	"gogoanime":    types.Provider(gogoanime.Handler),
}

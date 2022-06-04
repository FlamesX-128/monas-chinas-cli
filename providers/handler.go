package providers

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/providers/gogoanime"
	"github.com/FlamesX-128/monas-chinas-cli/src/providers/monoschinos2"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
)

var Handler = map[string]types.Provider{
	"[español] monoschinos2": monoschinos2.Handler,
	"[english] gogoanime":    gogoanime.Handler,
}

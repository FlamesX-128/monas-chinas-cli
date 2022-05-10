package providers

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/providers/gogoanime"
	"github.com/FlamesX-128/monas-chinas-cli/src/providers/monoschinos2"
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
)

var Providers = map[string]types.Provider{
	"[english] gogoanime":    gogoanime.Handler,
	"[español] monoschinos2": monoschinos2.Handler,
}

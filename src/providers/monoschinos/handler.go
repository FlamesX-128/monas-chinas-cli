package monoschinos

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
)

const url string = "https://monoschinos2.com/"

type provider struct {
	types.Translation
}

var Handler = &provider{types.Translation{}}

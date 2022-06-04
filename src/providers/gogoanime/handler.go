package gogoanime

import (
	"github.com/FlamesX-128/monas-chinas-cli/src/types"
)

const url string = "https://gogoanime.sk/"

type provider struct {
	types.Translation
}

var Handler = &provider{types.Translation{}}

package solvers

type Translation struct {
	NumberErr string

	ToWatch string
}

var Translations = map[string]Translation{
	"english": {
		NumberErr: "You must enter a number between 1 to %v",

		ToWatch: "What anime do you want to see?",
	},
	"español": {
		ToWatch: "¿Qué anime quieres ver?",
	},
}

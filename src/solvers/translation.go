package solvers

type Translation struct {
	ToWatch string
}

var Translations = map[string]Translation{
	"english": {
		ToWatch: "What anime do you want to see?",
	},
	"español": {
		ToWatch: "¿Qué anime quieres ver?",
	},
}

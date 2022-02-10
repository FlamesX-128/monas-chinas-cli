package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/FlamesX-128/monas-chinas-cli/src/interfaces"
)

func ListenAndServe(anime interfaces.Anime) {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		json.NewEncoder(rw).Encode(anime)
	})

	http.ListenAndServe(":43078", nil)
}

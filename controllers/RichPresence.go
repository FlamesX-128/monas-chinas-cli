package controllers

import (
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func RichPresence(name string, imageURL string, episode int) {
	if err := client.Login("940982082891579482"); err != nil {
		fmt.Println("No se pudo conectar con el cliente para tener presencia en Discord.")

		return
	}

	now := time.Now()

	for {
		time.Sleep(time.Second * 1)

		err := client.SetActivity(client.Activity{
			State:      fmt.Sprintf("Watching episode %d", episode),
			Details:    name,
			LargeImage: imageURL,
			LargeText:  name,
			Timestamps: &client.Timestamps{
				Start: &now,
			},
			Buttons: []*client.Button{
				{
					Label: "GitHub",
					Url:   "https://github.com/FlamesX-128/monas-chinas-cli",
				},
			},
		})

		if err != nil {
			fmt.Println("Sucedió un error al actualizar la presencia.")

			continue
		}
	}
}

package controllers

import "fmt"

func BuildPage(page [5]string, title string, handlers []serviceHandler, service string) string {
	addTitle := page[0] + title + page[1]

	var handlerTags string

	for _, handler := range handlers {
		handlerTags += handler.Tag

	}

	addHandlers := addTitle + handlerTags + page[2]

	var servicesTag string

	for _, handler := range handlers {
		servicesTag += fmt.Sprintf(`
			<li id="play-video">
				<a class="play-video dropdown-item cap" data-player="%s">
					%s
				</a>
			</li>
		`, handler.Data, handler.ID)
	}

	addServices := addHandlers + servicesTag + page[3]
	addService := addServices + service + page[4]

	return addService
}

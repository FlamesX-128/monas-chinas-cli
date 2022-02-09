package controllers

import "github.com/webview/webview"

func WebView(urls []string) {
	var url string

	if len(urls) == 2 {
		url = urls[1]

	} else {
		url = urls[0]

	}

	window := webview.New(false)
	defer window.Destroy()

	window.SetTitle("monas-chinas-cli")
	window.SetSize(800, 600, webview.HintNone)
	window.Navigate(url)

	window.Run()
}

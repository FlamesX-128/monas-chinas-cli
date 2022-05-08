package controllers

import "github.com/webview/webview"

func WebView(url string) {
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("monas-chinas-cli")
	w.SetSize(800, 600, webview.HintNone)

	w.Navigate(url)
	w.Run()

}

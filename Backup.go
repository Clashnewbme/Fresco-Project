print("This is a backup of main.go")

package main

import (
	"github.com/webview/webview_go"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Fresco Browser")
	w.SetSize(1000, 700, webview.HintNone)

	w.Navigate("https://setup-delta.vercel.app")

	w.Run()
}

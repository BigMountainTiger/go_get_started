package main

import (
	"os"

	"github.com/webview/webview"
)

func main() {

	os.Setenv("JSC_SIGNAL_FOR_GC", "20")

	debug := false
	w := webview.New(debug)
	defer w.Destroy()

	w.SetSize(600, 700, 0)

	w.SetTitle("Minimal webview example")
	w.Navigate("https://www.google.com")
	w.Run()

	// time.Sleep(1 * time.Minute)

}

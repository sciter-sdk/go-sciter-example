package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	// create rect for window
	rect := sciter.NewRect(200, 200, 800, 600)

	// create scister window object with rect
	win, _ := window.New(sciter.SW_MAIN, rect)

	// set title for window
	// win.SetTitle("Load HTML Page as UI")

	// load index.html file in window
	win.LoadFile("./index.html")

	// Launch sciter window
	win.Show()

	// run sciter application
	win.Run()
}

package main

import (
	"syscall"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	// make rect for window
	rect := sciter.NewRect(100, 100, 800, 200)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|
		sciter.SW_ENABLE_DEBUG, rect)

	win.SetTitle("custom-screen")


	win.SetResourceArchive(resources)
	win.LoadFile("this://app/htdocs/main.html")

	win.Show()
	win.Run()
	win.CloseArchive()
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

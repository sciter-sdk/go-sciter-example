package main

import (
	"syscall"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	// make rect for window
	rect := sciter.NewRect(100, 100, 300, 500)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, rect)
	win.LoadFile("./notepad.html")
	win.SetTitle("Notepad+-")

	// registring methods
	win.DefineFunction("closeWindow", closeApplication)
	win.DefineFunction("save", save)

	win.Show()
	win.Run()
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

func save(vals ...*sciter.Value) *sciter.Value {
	
	return nil
}

package main

import (
	"fmt"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

var appWindow *window.Window
var windowErr error

func main() {

	// Creating A Reactangle of size we want
	rect := sciter.NewRect(200, 200, 400, 400)
	// Create Window Over The rect
	appWindow, windowErr = window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|
		sciter.SW_ENABLE_DEBUG, rect)
	// If we cannot create window
	// Application execution has to be stopped
	// Because app has been failed in its first most stage
	if windowErr != nil {
		fmt.Errorf("Failed to create application window due to %s ", windowErr.Error())
		return
	}
	uiLoadErr := appWindow.LoadHtml(screens(0), "/")
	if uiLoadErr != nil {
		fmt.Errorf("Failed to Load UI dur to %s ", uiLoadErr.Error())
		return
	}

	appWindow.DefineFunction("changePage", changepage)

	appWindow.SetTitle("Score")
	// Showing window on screen
	appWindow.Show()
	// Making window Running ...
	appWindow.Run()

}

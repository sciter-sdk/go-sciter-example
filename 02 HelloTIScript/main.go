package main

import (
	"fmt"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	rect := sciter.NewRect(100, 100, 300, 300)
	window, windowCreationErr := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, rect)

	if windowCreationErr != nil {
		fmt.Errorf("Could not create sciter window : %s",
			windowCreationErr.Error())
		return
	}

	uiFileLoadErr := window.LoadFile("./main.html")
	if uiFileLoadErr != nil {
		fmt.Errorf("Could not load ui file : %s",
			uiFileLoadErr.Error())
	}

	// Setting up stage for Harmony
	window.SetTitle("Simple Input")
	window.Show()
	window.Run()

}

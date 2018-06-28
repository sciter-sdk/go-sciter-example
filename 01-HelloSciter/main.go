package main

import (
	"github.com/fatih/color"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	rect := sciter.NewRect(100, 100, 600, 400)
	window, windowsGenerateionError := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, rect)

	if windowsGenerateionError != nil {
		color.RedString("Failed to generate sciter window ", windowsGenerateionError.Error())
	}

	uiLoadingError := window.LoadFile("./hello.html")
	if uiLoadingError != nil {
		color.RedString("Failed to load ui file ", uiLoadingError.Error())
	}

	window.SetTitle("Hello ")
	window.Show()
	window.Run()

}

package main

import (
	"fmt"
	"os"

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

	uiLoadingError := window.LoadFile("./main.html")
	if uiLoadingError != nil {
		color.RedString("Failed to load ui file ", uiLoadingError.Error())
	}

	rootElementBind, _ := window.GetRootElement()

	btn2Bind, _ := rootElementBind.SelectById("bt2")
	btnValue, _ := btn2Bind.GetValue()

	btn2Bind.DefineMethod("tester", func(vals ...*sciter.Value) *sciter.Value {
		var text string
		for _, scits := range vals {
			text += scits.String()
		}
		file, _ := os.OpenFile("./sample.file", os.O_WRONLY|os.O_CREATE|os.O_RDWR, os.ModePerm)
		file.Write([]byte(text))
		return btnValue
	})
	// Setting up stage for Harmony
	window.SetTitle("Harmony")
	window.Show()
	window.Run()

}

func Mux(v1, v2 *sciter.Value) (a *sciter.Value) {
	fmt.Printf("wow its clicked ")
	return a
}

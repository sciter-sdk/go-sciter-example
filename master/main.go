package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	rect := sciter.NewRect(100, 100, 300, 300)
	window, windowsGenerateionError := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, rect)

	if windowsGenerateionError != nil {
		color.RedString("Failed to generate sciter window ", windowsGenerateionError.Error())
	}

	uiLoadingError := window.LoadFile("./main.html")
	if uiLoadingError != nil {
		color.RedString("Failed to load ui file ", uiLoadingError.Error())
	}

	window.DefineFunction("Calc", Calc)

	// Setting up stage for Harmony
	window.SetTitle("Simple Input")
	window.Show()
	window.Run()

}

func Calc(vals ...*sciter.Value) *sciter.Value {
	ans := 0
	switch vals[2].String() {
	case "+":
		{
			ans = vals[1].Int() + vals[0].Int()
		}
	case "-":
		{
			ans = vals[1].Int() - vals[0].Int()
		}
	case "*":
		{
			ans = vals[1].Int() * vals[0].Int()
		}
	case "/":
		{
			ans = vals[1].Int() / vals[0].Int()
		}
	}
	fmt.Println("Ans is ", ans)
	return sciter.NewValue(ans)
}

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

	uiLoadingError := window.LoadFile("./main-update.html")
	if uiLoadingError != nil {
		color.RedString("Failed to load ui file ", uiLoadingError.Error())
	}

	window.DefineFunction("Operate", Operate)

	// Setting up stage for Harmony
	window.SetTitle("Simple Input")
	window.Show()
	window.Run()

}

func Operate(vals ...*sciter.Value) *sciter.Value {
	sumval := 0
	switch vals[2].String() {
	case "+":
		{
			sumval = vals[0].Int() + vals[1].Int()
		}
	case "-":
		{
			sumval = vals[0].Int() - vals[1].Int()
		}
	case "*":
		{
			sumval = vals[0].Int() * vals[1].Int()
		}
	case "/":
		{
			sumval = vals[0].Int() / vals[1].Int()
		}
	default:
		{
			fmt.Println("undefined opertaion", vals[2].String())
		}
	}
	fmt.Println("summation is ", sumval)
	// sumString := fmt.Sprintf("%v", sumval)

	return sciter.NewValue(sumval)
}

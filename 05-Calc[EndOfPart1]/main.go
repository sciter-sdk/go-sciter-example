package main

import (
	"fmt"
	"strconv"
	"strings"

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

// Operate function just expects one parametre
// mathamatic string containing both
// operands and operator
func Operate(val ...*sciter.Value) *sciter.Value {

	// Trim All space as we are taking
	//  now input as string
	trimmedString := val[0].String()

	// if input empty or not
	if strings.TrimSpace(trimmedString) == "" {
		fmt.Println("Invalid input ")
		return nil
	}

	opGroup := []string{"+", "-", "*", "/"}

	for _, op := range opGroup {
		if strings.Contains(trimmedString, op) {
			fmt.Println("we found ", op, " operator in string")
			inputString := strings.Split(trimmedString, op)
			op1, _ := strconv.Atoi(inputString[0])
			op2, _ := strconv.Atoi(inputString[1])
			switch op {
			case "+":
				{
					return sciter.NewValue(op1 + op2)
				}
			case "-":
				{
					return sciter.NewValue(op1 - op2)
				}
			case "/":
				{
					return sciter.NewValue(op1 / op2)
				}
			case "*":
				{
					return sciter.NewValue(op1 * op2)
				}
			default:
				{
					fmt.Println("Awesome !!! , but ... no operator found ")
				}
			}
		}

	}
	return nil
}

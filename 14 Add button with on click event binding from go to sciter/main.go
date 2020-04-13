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

	// load index.html file in window
	win.LoadFile("./index.html")

	// Add button from go...

	// 1. Select root element of window
	rootEl, _ := win.GetRootElement()

	// 1.1 select button-spot div
	buttonspotEl, _ := rootEl.SelectById("button-spot")

	// 2. Preapre button element
	buttonEl, _ := sciter.CreateElement("button", "Click me button")

	// 3. Append button to button-spot element
	buttonspotEl.Append(buttonEl)

	// 4. Add onclick event for button
	buttonEl.OnClick(func() {
		buttonEl.SetText("You clicked me")
	})

	// Launch sciter window
	win.Show()
	// run sciter application
	win.Run()
}

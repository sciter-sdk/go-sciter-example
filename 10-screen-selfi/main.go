package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"syscall"

	"github.com/satori/go.uuid"

	"github.com/vova616/screenshot"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

// All read images will
// be stored in Images array
// as base64 strings
// It will be reparsed by sciter
var rootElement *sciter.Element
var err error

func main() {
	// make rect for window
	rect := sciter.NewRect(0, 0, 800, 200)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_POPUP|sciter.SW_CONTROLS|
		sciter.SW_ENABLE_DEBUG, rect)

	win.DefineFunction("snapNow", snapCalled)
	win.SetTitle("ScreenSefli+-")

	win.SetResourceArchive(resources)
	win.LoadFile("this://app/htdocs/screen.htm")

	rootElement, err = win.GetRootElement()
	if err != nil {
		fmt.Println(" failed to load root element")
		return
	}

	win.Show()
	win.Run()
	win.CloseArchive()
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

func snapCalled(vals ...*sciter.Value) *sciter.Value {

	x1 := vals[0].Int()
	y1 := vals[1].Int()
	x2 := vals[2].Int()
	y2 := vals[3].Int()

	takeASelfi(x1, y1, x2, y2)
	return nil
}

func takeASelfi(xi, yi, xe, ye int) {

	sefliRect := image.Rect(xi, yi, xe, ye)
	sefliData, err := screenshot.CaptureRect(sefliRect)
	if err != nil {
		fmt.Println("We failed to take a sefli. sorry ....", err.Error())
		return
	}

	imageName, _ := uuid.NewV4()
	f, err := os.Create("./" + imageName.String() + ".png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, sefliData)
	if err != nil {
		panic(err)
	}
	f.Close()

}

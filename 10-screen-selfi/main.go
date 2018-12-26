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

func main() {
	// make rect for window
	rect := sciter.NewRect(0, 0, 800, 200)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|
		sciter.SW_ENABLE_DEBUG, rect)

	win.DefineFunction("snapNow", snapCalled)
	win.DefineFunction("closeApp", closeApplication)
	win.SetTitle("ScreenSefli+-")

	win.SetResourceArchive(resources)
	win.LoadFile("this://app/htdocs/main.htm")

	win.Show()
	win.Run()
	win.CloseArchive()
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

// snapCalled is binding for sciter-Go
// it calls takeASelfi function after
// successfully getting required inputs
// from sciter ...
func snapCalled(vals ...*sciter.Value) *sciter.Value {

	// If inputs are not exactly 4
	// then something is wrong ...
	if len(vals) == 4 {
		x1 := vals[0].Int()
		y1 := vals[1].Int()
		x2 := vals[2].Int()
		y2 := vals[3].Int()
		fmt.Println(x1, y2, x2, y2, " are the cordinates")
		takeASelfi(x1, y1, x2, y2)
	}

	return nil
}

// takeASefli takes two cordinates as input
// creates a rectangle from those cordinates
// and takes snaps of that rectanle and stores
// it as a png image
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

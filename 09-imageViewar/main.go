package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	// make rect for window
	rect := sciter.NewRect(100, 100, 300, 500)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_RESIZEABLE|sciter.SW_ENABLE_DEBUG, rect)

	win.SetTitle("ImageViewar+-")

	// registering methods
	win.DefineFunction("closeWindow", closeApplication)
	win.DefineFunction("testImage", TestImage)

	win.SetResourceArchive(resources)
	win.LoadFile("this://app/htdocs/notepad.htm")

	win.Show()
	win.Run()
	win.CloseArchive()
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

// Find and load every image file
// from current directory and
// store them all inside array of string
func findAndLoadImageFromCurrentDirectory() {
	// Getting working directory
	this_dir, dirErr := os.Getwd()
	if dirErr != nil {
		fmt.Println("Failed to get current directory")
		return
	}
	files, readDirErr := ioutil.ReadDir(this_dir)
	if readDirErr != nil {
		fmt.Println("failed to read current directory")
		return
	}
	for _, file := range files {

		// Just supporting jpg and png file to be loaded
		// others are on the way .. .
		if strings.Contains(file.Name(), ".jpg") || strings.Contains(file.Name(), ".png") {
			imageFile, imageFileErr := os.Open(filepath.Join(this_dir, file.Name()))
			if imageFileErr != nil {
				fmt.Println("failed to load image file")
				return
			}
			state, stateError := imageFile.Stat()
			if stateError != nil {
				fmt.Println("failed to get state of the image file ")
				return
			}
			size := state.Size()
			buf := make([]byte, size)

			// Reading image file in buffer
			fReader := bufio.NewReader(imageFile)
			fReader.Read(buf)

			// Convert file to base64
			imgStrging := base64.StdEncoding.EncodeToString(buf)

		}
	}

}

func TestImage(vals ...*sciter.Value) *sciter.Value {

	fmt.Println("test image function called")

	testFile, fileOpenError := os.Open("./9999999991_01.jpg")
	if fileOpenError != nil {
		fmt.Println("failed to open file %s", fileOpenError.Error())
		return nil
	}

	// create a new buffer base on file size
	fInfo, _ := testFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// Reading image file in buffer
	fReader := bufio.NewReader(testFile)
	fReader.Read(buf)

	// Convert file to base64
	imgStrging := base64.StdEncoding.EncodeToString(buf)
	return sciter.NewValue(imgStrging)
}

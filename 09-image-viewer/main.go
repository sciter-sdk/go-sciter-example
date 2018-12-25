package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

var Index int       // Stores current index of image
var Images []string // Images stores base64 string of images

func main() {
	// make rect for window
	rect := sciter.NewRect(0, 0, 800, 600)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|
		sciter.SW_RESIZEABLE|sciter.SW_ENABLE_DEBUG, rect)

	win.SetTitle("ImageViewar+-")

	// Scanning and loading base64 of images
	findAndLoadImageFromCurrentDirectory()

	// registering methods
	win.DefineFunction("loadFirstImage", LoadFirstImage)
	win.DefineFunction("loadNextImage", LoadNextImage)
	win.DefineFunction("loadPreviousImage", LoadPreviousImage)
	win.DefineFunction("closeWindow", closeApplication)

	// Getting data from archive
	win.SetResourceArchive(resources)
	win.LoadFile("this://app/htdocs/image-viewer.htm")

	// Running app
	win.Show()
	win.Run()
	win.CloseArchive()
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

// findAndLoadImageFromCurrentDirectory scans directory
// in which exec is for jpg / png files. Reads those files
// and store base64 string of those file in Images([]string)
func findAndLoadImageFromCurrentDirectory() {

	var waitGroup sync.WaitGroup
	// Getting working directory
	thisDir, dirErr := os.Getwd()
	if dirErr != nil {
		fmt.Println("Failed to get current directory")
		return
	}
	files, readDirErr := ioutil.ReadDir(thisDir)
	if readDirErr != nil {
		fmt.Println("failed to read current directory")
		return
	}

	if len(files) > 0 {
		imgString := getImageString(files[0], thisDir)
		if imgString != "" {
			Images = append(Images, imgString)
		}
	}

	// Loading files excpet first via goroutine
	// so we don't have to wait for every image
	// to be loaded to show up first image
	waitGroup.Add(1)
	go func() {
		for i, file := range files {
			if i == 0 {
				continue
			}
			imgString := getImageString(file, thisDir)
			if imgString != "" {
				Images = append(Images, imgString)
			}
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}

// LoadFirstImage return first
// image from Image array
// to sciter
func LoadFirstImage(vals ...*sciter.Value) *sciter.Value {
	if len(Images) > 0 {
		Index = 0
		fmt.Println("Returning first image")
		return sciter.NewValue(Images[0])
	}
	return sciter.NewValue(string("-"))
}

// LoadNextImage return image from
// next index to current index
func LoadNextImage(vals ...*sciter.Value) *sciter.Value {
	if Index < len(Images)-1 {
		Index++
		return sciter.NewValue(Images[Index])
	}
	Index = 0
	return sciter.NewValue(Images[0])
}

// LoadPreviousImage return image from
// previous index to current index
func LoadPreviousImage(vals ...*sciter.Value) *sciter.Value {
	if Index > 0 {
		Index--
		return sciter.NewValue(Images[Index])
	}
	Index = len(Images) - 1
	return sciter.NewValue(Images[0])
}

// getImageString returns base64 string
// of file provided as input
func getImageString(file os.FileInfo, thisDir string) string {

	// Just supporting jpg and png file to be loaded
	// others are on the way .. .
	if strings.Contains(file.Name(), ".jpg") || strings.Contains(file.Name(), ".png") {
		imageFile, imageFileErr := os.Open(filepath.Join(thisDir, file.Name()))
		if imageFileErr != nil {
			fmt.Println("failed to load image file")
			return ""
		}
		state, stateError := imageFile.Stat()
		if stateError != nil {
			fmt.Println("failed to get state of the image file ")
			return ""
		}
		size := state.Size()
		buf := make([]byte, size)

		// Reading image file in buffer
		fReader := bufio.NewReader(imageFile)
		fReader.Read(buf)

		// Convert file to base64
		imgStrging := base64.StdEncoding.EncodeToString(buf)
		return imgStrging
	}
	return ""
}

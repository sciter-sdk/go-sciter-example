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

// All read images will
// be stored in Images array
// as base64 strings
// It will be reparsed by sciter

var Index int
var Images []string

func main() {
	// make rect for window
	rect := sciter.NewRect(0, 0, 800, 600)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|
		sciter.SW_RESIZEABLE|sciter.SW_ENABLE_DEBUG, rect)

	win.SetTitle("ImageViewar+-")

	findAndLoadImageFromCurrentDirectory()
	// registering methods
	win.DefineFunction("closeWindow", closeApplication)
	// win.DefineFunction("testImage", TestImage)

	win.DefineFunction("loadFirstImage", LoadFirstImage)
	win.DefineFunction("loadNextImage", LoadNextImage)
	win.DefineFunction("loadPreviousImage", LoadPreviousImage)

	win.SetResourceArchive(resources)
	win.LoadFile("this://app/htdocs/image-viewer.htm")

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

	var waitGroup sync.WaitGroup
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

	if len(files) > 0 {
		imgString := getImageString(files[0], this_dir)
		if imgString != "" {
			Images = append(Images, imgString)
		}
		fmt.Println("first image loaded to array")
	}

	// Loading rest of file
	// with different go routinbe
	// to make process more fast
	// and smooth
	waitGroup.Add(1)
	go func() {
		for i, file := range files {
			if i == 0 {
				continue
			}
			imgString := getImageString(file, this_dir)
			if imgString != "" {
				Images = append(Images, imgString)
			}
			fmt.Println("new image added to array", len(Images))
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("Images been loaded", len(Images))
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
		fmt.Println("Returning image on index ", Index, "for loadNext")
		return sciter.NewValue(Images[Index])
	}
	fmt.Println("Total Imagaes ", len(Images))
	Index = 0
	fmt.Println("Returning image on index by indexin zero ", Index, "for loadNext")
	return sciter.NewValue(Images[0])
}

// LoadPrevoiusImage return image from
// previous index to current index
func LoadPreviousImage(vals ...*sciter.Value) *sciter.Value {
	if Index > 0 {
		Index--
		fmt.Println("Returning image on index ", Index, "for loadPrevious")
		return sciter.NewValue(Images[Index])
	}
	Index = len(Images) - 1
	fmt.Println("Returning image on index by indexin zero ", Index, "for loadPrevious")
	return sciter.NewValue(Images[0])
}

// getImageString accepts file and
// directory path to that file
// It will find file,open it
// and will return base64 contetn
// for that file if it is .jpg / .png
// else it will return empty string
func getImageString(file os.FileInfo, this_dir string) string {

	// Just supporting jpg and png file to be loaded
	// others are on the way .. .
	if strings.Contains(file.Name(), ".jpg") || strings.Contains(file.Name(), ".png") {
		imageFile, imageFileErr := os.Open(filepath.Join(this_dir, file.Name()))
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

// func TestImage(vals ...*sciter.Value) *sciter.Value {

// 	fmt.Println("test image function called")

// 	testFile, fileOpenError := os.Open(--somefix-image-path--)
// 	if fileOpenError != nil {
// 		fmt.Println("failed to open file %s", fileOpenError.Error())
// 		return nil
// 	}

// 	// create a new buffer base on file size
// 	fInfo, _ := testFile.Stat()
// 	var size int64 = fInfo.Size()
// 	buf := make([]byte, size)

// 	// Reading image file in buffer
// 	fReader := bufio.NewReader(testFile)
// 	fReader.Read(buf)

// 	// Convert file to base64
// 	imgStrging := base64.StdEncoding.EncodeToString(buf)
// 	return sciter.NewValue(imgStrging)
// }

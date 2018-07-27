package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	// make rect for window
	rect := sciter.NewRect(100, 100, 300, 500)

	// create a window using upper rect
	win, _ := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, rect)

	win.DefaultHandler()
	win.OpenArchive(resources)

	// csbytes := win.GetArchiveItem("style.css")
	htbytes := win.GetArchiveItem("notepad.htm")

	win.LoadHtml(string(htbytes), "")

	win.SetTitle("Notepad+-")

	// registering methods
	win.DefineFunction("closeWindow", closeApplication)
	win.DefineFunction("save", save)
	win.DefineFunction("open", open)

	win.Show()
	win.Run()
	win.CloseArchive()
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

func save(vals ...*sciter.Value) *sciter.Value {

	fmt.Println("Saving Your Document")
	path := vals[0]
	doc := vals[1]

	processedFilePath := strings.Replace(path.String(), "file://", "", 1)
	file, fileCreationError := os.OpenFile(processedFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)

	if fileCreationError != nil {
		fmt.Println("failed to create a blank file ", fileCreationError.Error())
		return nil
	}

	defer file.Close()

	charCount, writeError := file.WriteString(doc.String())
	if writeError != nil {
		fmt.Println("failed to write on file due to : ", writeError.Error())
		return nil
	}
	fmt.Println("chars written ", charCount)
	fmt.Println("we got path as ", path.String())
	return nil
}

func open(vals ...*sciter.Value) *sciter.Value {
	fmt.Println("Saving Your Document")
	path := vals[0]
	processedFilePath := strings.Replace(path.String(), "file://", "", 1)

	readBytes, readError := ioutil.ReadFile(processedFilePath)

	if readError != nil {
		fmt.Println("failed to write on file due to : ", readError.Error())
		return nil
	}
	fmt.Println("chars written ", len(readBytes))
	fmt.Println("we got path as ", path.String())
	return sciter.NewValue(string(readBytes))
}

package main

import (
	"fmt"

	"github.com/sciter-sdk/go-sciter"
)

func changepage(opt ...*sciter.Value) *sciter.Value {
	fmt.Println("yeah , we are on right path")
	appWindow.LoadHtml(screens(1), "/")
	return nil
}

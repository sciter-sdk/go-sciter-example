package main

/**
*	every mac os x app is incomplete with-out its own unity bar menus
*   .....
*   You can add it directly via sciter or golang itself. But ...
*   You can achive this via cgo [ And to be specific via Objective-c code]
*   How !.
*
*   See this .
 */

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>


int LoadMenu(void) {
    [NSAutoreleasePool new];
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
    id menubar = [[NSMenu new] autorelease];
    id appMenuItem = [[NSMenuItem new] autorelease];
    [menubar addItem:appMenuItem];
    [NSApp setMainMenu:menubar];
    id appMenu = [[NSMenu new] autorelease];
    id appName = [[NSProcessInfo processInfo] processName];
    id quitTitle = [@"Quit " stringByAppendingString:appName];
    id quitMenuItem = [[[NSMenuItem alloc] initWithTitle:quitTitle
        action:@selector(terminate:) keyEquivalent:@"q"]
			  autorelease];

	[appMenu addItem:[[[NSMenuItem alloc] initWithTitle:@"I am The menu" action:@selector(caller:) keyEquivalent:@""] autorelease] ];
	[appMenu addItem:quitMenuItem];

    [appMenuItem setSubmenu:appMenu];

return 0;
}
*/
import "C"

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	// Call the cgo function that will bind menu to unitybar
	C.LoadMenu()

	// create rect for window
	rect := sciter.NewRect(200, 200, 800, 600)

	// create scister window object with rect
	win, _ := window.New(sciter.DefaultWindowCreateFlag, rect)

	// load index.html file in window
	win.LoadFile("./index.html")

	win.Show()

	win.Run()

}

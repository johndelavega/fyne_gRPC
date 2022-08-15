package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

//
// mobile device specific height, fyne.Window.Canvas().Size()
// Pixel 4a = 780, _mainMobileDeltaHeight = 52
// Nexus 5X = 598, _mainMobileDeltaHeight = 25
//

// var _mainAppVersion = "v0.0.0"
// var _mainInternalBuildNumber = "10000"

var _mainAppType string = "Desktop" // Mobile

var _mainMobileDeltaHeight float32 = 0

func fyneApp() {

	fmt.Printf("mobile %s\n\n", _mainAppVersion)
	// fmt.Printf("single arg options: debug | geom (background graphics)\n\n")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	bDebug := false
	bGeom := false
	_ = bGeom

	args := os.Args

	if len(args) > 1 {
		if strings.ToLower(args[1]) == "debug" {
			bDebug = true
		}
		if strings.ToLower(args[1]) == "geom" {
			bGeom = true
		}
	}

	if bDebug {
		fmt.Printf("debugging... print main() start\n\n")
	}

	a := app.NewWithID("com.fynetest.mobile1")
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("Fyne App")

	w.SetMaster()

	w.Resize(fyne.NewSize(300, 450))

	w.SetContent(pageInit(w)) // init page, ok button only
	w.CenterOnScreen()

	fmt.Printf("debug: before w.Show() 1 : %v\n", w.Canvas().Size())
	// w.ShowAndRun()

	w.Show()

	fmt.Printf("debug: after  w.Show() 2 : %v\n", w.Canvas().Size())

	a.Lifecycle().SetOnStarted(func() {
		fmt.Println("debug: before  a.Run() test1 - a.Lifecycle().SetOnStarted(...)")
		fmt.Printf("debug: before  a.Run() test1 : %v\n", w.Canvas().Size())

	})

	a.Run()

	a.Lifecycle().SetOnStarted(func() {
		// not called
		fmt.Println("debug: after a.Run() test2 - a.Lifecycle().SetOnStarted(...)")
		fmt.Printf("debug: after a.Run() %v\n", w.Canvas().Size())

	})
}

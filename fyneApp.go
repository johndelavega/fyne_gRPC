package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

//
// mobile device specific height, fyne.Window.Canvas().Size()
// Pixel 4a = 780, _mobileDeltaHeight = 52
// Nexus 5X = 598, _mobileDeltaHeight = 25
//

var _appType string = "Desktop" // or Mobile

var _mobileDeltaHeight float32 = 0

func fyneApp() {

	fmt.Printf("fyneApp gRPC %s\n\n", _mainAppVersion)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	a := app.NewWithID("com.fynetest.mobile1")
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("fyneApp gRPC")

	w.SetMaster()

	w.Resize(fyne.NewSize(300, 450))
	w.SetFixedSize(true)

	w.SetContent(pageInit(w)) // init page, ok button only
	w.CenterOnScreen()

	w.Show()

	a.Lifecycle().SetOnStarted(func() {
		fmt.Println("debug: before fyne.App.Run() test1 - a.Lifecycle().SetOnStarted(...)")
		fmt.Printf("debug: before fyne.App.Run() canvas size : %v\n", w.Canvas().Size())

	})

	a.Run()

}

//
// Initial page
//

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//
// mobile device specific height, fyne.Window.Canvas().Size()
// Pixel 4a = 780, _mainMobileDeltaHeight = 52
// Nexus 5X = 598, _mainMobileDeltaHeight = 25
//

func pageInit(w fyne.Window) *fyne.Container {

	var content *fyne.Container

	// noLayout := container.NewWithoutLayout()
	center := container.NewCenter()

	if fyne.CurrentDevice().IsMobile() {
		_mainAppType = "Mobile"
	}

	// noLayout.Objects = []fyne.CanvasObject{widget.NewLabel("center")}
	centerText := fmt.Sprintf("Press OK to Start\n\n%s app %s\n\ndebug: center layout\n\n\n\n\nbuild# %s", _mainAppType, _mainAppVersion, _mainInternalBuildNumber)
	center.Objects = []fyne.CanvasObject{widget.NewLabel(centerText)}

	btnOK := widget.NewButton("OK", func() {
		fmt.Printf("debug: btnOK - OK : %v\n", w.Canvas().Size())

		w.SetContent(page0(w, content, fyne.CurrentDevice().IsMobile()))
	})
	btnOK.Resize(fyne.NewSize(100, 100))

	// WithoutLayout areas is entire above button
	right := container.NewBorder(nil, nil, nil, btnOK)
	content = container.NewBorder(nil, right, nil, nil, center) // WithoutLayout can work within a Layout ie BorderLayout

	return content

}

// func page0(w fyne.Window, c *fyne.Container, bMobile bool) *fyne.Container {

// 	var content *fyne.Container

// 	if fyne.CurrentDevice().IsMobile() {

// 		if w.Canvas().Size().Height == 780.0 {
// 			_mainMobileDeltaHeight = 46
// 		} else if w.Canvas().Size().Height == 598.0 {
// 			_mainMobileDeltaHeight = 25
// 		} else {
// 			_mainMobileDeltaHeight = 50
// 		}
// 	}

// 	content = container.NewWithoutLayout()

// 	// content.Objects = []fyne.CanvasObject{lblVer, lblDebug, lblRand, btnRandom, btnHeight, btnPage1, btnPage2}
// 	// content.Objects = []fyne.CanvasObject{lblVer, lblRand, lblDebug, vbox, btnInitPage}

// 	return content

// }

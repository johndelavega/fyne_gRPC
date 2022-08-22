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
// Pixel 4a = 780, _mobileDeltaHeight = 52
// Nexus 5X = 598, _mobileDeltaHeight = 25
//

func pageInit(w fyne.Window) *fyne.Container {

	var content *fyne.Container

	// noLayout := container.NewWithoutLayout()
	center := container.NewCenter()

	if fyne.CurrentDevice().IsMobile() {
		_appType = "Mobile"
	}

	// noLayout.Objects = []fyne.CanvasObject{widget.NewLabel("center")}
	centerText := fmt.Sprintf("Press OK to Start\n\n%s app %s\n\ndebug: center layout\n\n\n\n\nbuild# %s", _appType, _mainAppVersion, _mainInternalBuildNumber)
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

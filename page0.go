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

func page0(w fyne.Window, c *fyne.Container, bMobile bool) *fyne.Container {

	var content *fyne.Container

	if fyne.CurrentDevice().IsMobile() {

		if w.Canvas().Size().Height == 780.0 {
			_mobileDeltaHeight = 46
		} else if w.Canvas().Size().Height == 598.0 {
			_mobileDeltaHeight = 25
		} else {
			_mobileDeltaHeight = 50
		}
	}

	lblVer := widget.NewLabel(fmt.Sprintf("%s app %s", _appType, _mainAppVersion))
	lblVer.Move(fyne.Position{X: 10, Y: 10})

	lblRand := widget.NewLabel("rand int32 from server:")
	lblRand.Move(fyne.Position{X: 10, Y: 60})

	lblDebug := widget.NewLabel("lblDebug Home Page0")
	lblDebug.Move(fyne.Position{X: 10, Y: 90})

	btnInitPage := widget.NewButton(" back to Init page", func() {
		// fmt.Printf("debug: init page button, container = %v\n", c)
		w.SetContent(c)
	})
	btnInitPage.Move(fyne.Position{X: 10, Y: 140})
	btnInitPage.Resize(fyne.Size{Width: 150, Height: 37})
	btnInitPage.Hidden = true

	btnPage1 := widget.NewButton("Clear message", func() {

		// w.SetContent(page1(w, content))
		lblDebug.SetText("message cleared")
	})

	btnPage2 := widget.NewButton("Open page2", func() {
		// w.SetContent(page2(w, content))
	})
	btnPage2.Disable()
	// btnPage2.Hidden = true

	btnGrid := widget.NewButton("Open grid", func() {
		// w.SetContent(grid(w, content))
	})
	btnGrid.Disable()

	btnExit := widget.NewButton("Exit", func() {
		w.Close()
	})

	btnHeight := widget.NewButton("canvas Height", func() {

		lblDebug.SetText("debug canvas Height - " + fmt.Sprintf("%f", w.Canvas().Size().Height))

		w.SetContent(content)
	})
	btnHeight.Disable()

	btnRandom := widget.NewButton("gRPC msg rand", func() {
		m, r := SayHello()

		lblRand.SetText("rand int32 from server: " + fmt.Sprintf("%d", r))

		lblDebug.SetText(fmt.Sprintf("message from server:   %s", m))

	})

	vbox := container.NewVBox(btnRandom, btnHeight, btnPage1, btnPage2, btnGrid)
	if !bMobile {
		vbox.Objects = append(vbox.Objects, widget.NewSeparator(), btnExit)
	}

	size1 := w.Canvas().Size()

	sizeY := size1.Height - 260 - _mobileDeltaHeight
	vbox.Move(fyne.Position{X: size1.Width - 160, Y: sizeY})

	// lblDebug.SetText(fmt.Sprintf("debug: y = %v MobileDeltaHeight = %v", sizeY, _mobileDeltaHeight))
	lblDebug.SetText("message from server:")

	vbox.Resize(fyne.Size{Width: 150, Height: 400})

	content = container.NewWithoutLayout()

	content.Objects = []fyne.CanvasObject{lblVer, lblRand, lblDebug, vbox, btnInitPage}

	return content

}

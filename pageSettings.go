package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var buttonWidth float32 = 42
var canvasBottomMargin float32 = 100
var canvasRightMargin float32 = 200

func pageSettings(w fyne.Window, c *fyne.Container) *fyne.Container {

	var input *widget.Entry
	closeButton := widget.NewButton(" Close", func() {
		if *addr != input.Text {
			addr = &input.Text
			grpcClientInit()
		}
		w.SetContent(c)
	})
	closeButton.Resize(fyne.Size{Width: 150, Height: 37})
	input = widget.NewEntry()
	input.SetPlaceHolder("Enter socket address")
	input.Text = *addr
	input.Move(fyne.Position{X: 0, Y: 50})

	// looks like Entry widget does not work with WithoutLayout container
	// so we'll put it in VBox container first
	conVBoxEntry := container.NewVBox()
	conVBoxEntry.Move(fyne.Position{X: 0, Y: 70})
	conVBoxEntry.Resize(fyne.Size{Width: 200, Height: 400})
	conVBoxEntry.Objects = []fyne.CanvasObject{input}

	size1 := w.Canvas().Size()

	conVBox := container.NewVBox()
	conVBox.Move(fyne.Position{X: size1.Width - canvasRightMargin, Y: size1.Height - canvasBottomMargin*3})
	conVBox.Resize(fyne.Size{Width: 150, Height: 400})

	conVBox.Objects = []fyne.CanvasObject{closeButton}

	lblPage1 := widget.NewLabel("Settings\n\nsocket address:")

	content := container.NewWithoutLayout()

	content.Objects = []fyne.CanvasObject{lblPage1, conVBoxEntry, conVBox}

	return content

}

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func MakeCharBox(char rune) *widget.Entry {
	entry := widget.NewEntry()
	entry.SetText(string(char))
	entry.SetPlaceHolder(".")
	// Limit width to single character
	entry.Resize(fyne.NewSize(30, 30))
	// Limit input to single character
	entry.OnChanged = func(s string) {
		if len(s) == 0 {
			entry.SetText(".")
		} else {
			entry.SetText(string([]rune(s)[0]))
		}
	}
	return entry
}

func MakeTextBox(text, placeHolder string) *fyne.Container {
	entry := widget.NewEntry()
	entry.SetText(text)
	entry.SetPlaceHolder(placeHolder)
	spacer := MakeTextBoxSpacer()

	return container.NewStack(spacer, entry)
}

func MakeCharBoxSpacer() fyne.CanvasObject {
	mock := widget.NewEntry()
	mock.SetText("*")
	mock.SetPlaceHolder("*")
	// Limit width to single character
	mock.Resize(fyne.NewSize(30, 30))
	mockSize := mock.MinSize()

	spacer := canvas.NewRectangle(color.Transparent)
	spacer.Resize(mockSize)
	spacer.SetMinSize(mockSize)

	return spacer
}

func MakeTextBoxSpacer() fyne.CanvasObject {
	spacer := canvas.NewRectangle(color.Transparent)
	spacer.Resize(fyne.NewSize(400, 30))
	spacer.SetMinSize(fyne.NewSize(400, 30))

	return spacer
}

func MakeEntrySpacer() fyne.CanvasObject {
	mock := widget.NewEntry()
	mockSize := mock.MinSize()

	spacer := canvas.NewRectangle(color.Black)
	spacer.Resize(mockSize)
	spacer.SetMinSize(mockSize)

	return spacer
}

func MakeCharLine(box *fyne.Container, str string) {
	box.RemoveAll()
	for _, char := range str {
		box.Add(MakeCharBox(char))
	}
}

func GetCharLine(box *fyne.Container) string {
	ret := ""
	for _, entry := range box.Objects {
		if entry, ok := entry.(*widget.Entry); ok {
			ret += entry.Text
		}
	}
	return ret
}

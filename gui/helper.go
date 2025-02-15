package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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

func MakeButton(label string, tapped func()) *fyne.Container {
	button := widget.NewButton(label, tapped)
	spacer := MakeTextBoxSpacer()

	return container.NewStack(spacer, button)
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

func MakeCharBoxArrow(upward bool) fyne.CanvasObject {
	spacer := MakeCharBoxSpacer()

	var arrow *canvas.Text
	if upward {
		arrow = canvas.NewText("\u2B67", theme.Color(theme.ColorNameForeground))
	} else {
		arrow = canvas.NewText("\u2B68", theme.Color(theme.ColorNameForeground))
	}

	arrow.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}
	arrow.TextSize = 30
	arrow.Alignment = fyne.TextAlignCenter
	arrow.Resize(spacer.MinSize())

	return arrow
}

func MakeCharBoxSpacerRow(length int) *fyne.Container {
	row := container.NewHBox()

	for i := 0; i < length; i++ {
		box := MakeCharBoxSpacer()
		row.Add(box)
	}

	return row
}

func MakeCharBoxSpacerGrid(rows, cols int) *fyne.Container {
	row := container.NewVBox()
	for i := 0; i < rows; i++ {
		row.Add(MakeCharBoxSpacerRow(cols))
	}
	return row
}

func MakeTextBoxSpacer() fyne.CanvasObject {
	spacer := canvas.NewRectangle(color.Transparent)
	spacer.Resize(fyne.NewSize(300, 30))
	spacer.SetMinSize(fyne.NewSize(300, 30))

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

// GetEntryText retrieves the entry.Text out of stacks of (spacer, entry)
func GetEntryText(box *fyne.Container) (string, bool) {
	entry, ok := box.Objects[1].(*widget.Entry)
	if !ok {
		return "", false
	}
	return entry.Text, true
}

func ReverseBox(container *fyne.Container) *fyne.Container {
	for i, j := 0, len(container.Objects)-1; i < j; i, j = i+1, j-1 {
		container.Objects[i], container.Objects[j] = container.Objects[j], container.Objects[i]
	}
	return container
}

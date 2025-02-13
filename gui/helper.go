package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
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

package lin

import (
	"crossmatcher/gui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
	"time"
)

type View struct {
	window        fyne.Window
	model         *Model
	lengthEntry   *widget.Entry
	ruleEntry     *widget.Entry
	alphabetEntry *widget.Entry
	charBoxes     *fyne.Container
	content       *fyne.Container
}

func NewView(window fyne.Window, rule string, alphabet string, candidate string) *View {
	v := &View{}

	v.window = window
	v.model = NewModel(rule, alphabet, candidate)

	v.lengthEntry = widget.NewEntry()
	v.lengthEntry.SetText(strconv.Itoa(len(candidate)))
	v.lengthEntry.SetPlaceHolder("Length")
	v.ruleEntry = widget.NewEntry()
	v.ruleEntry.SetText(rule)
	v.ruleEntry.SetPlaceHolder("Regex Rule")
	v.alphabetEntry = widget.NewEntry()
	v.alphabetEntry.SetText(alphabet)
	v.alphabetEntry.SetPlaceHolder("Alphabet characters")
	v.charBoxes = container.NewHBox()
	gui.MakeCharLine(v.charBoxes, candidate)

	updateLengthButton := widget.NewButton("Update Length", v.onUpdateLength)
	solveButton := widget.NewButton("Solve", v.onSolve)

	v.content = container.NewVBox(
		widget.NewLabel("Length:"),
		v.lengthEntry,
		updateLengthButton,
		widget.NewLabel("Rule:"),
		v.ruleEntry,
		widget.NewLabel("Alphabet:"),
		v.alphabetEntry,
		widget.NewLabel("Solution:"),
		v.charBoxes,
		solveButton)

	return v
}

func (v *View) onSolve() {
	rule := v.ruleEntry.Text
	alphabet := v.alphabetEntry.Text
	candidate := gui.GetCharLine(v.charBoxes)

	v.model = NewModel(rule, alphabet, candidate)
	candidate = v.model.Solve()

	gui.MakeCharLine(v.charBoxes, candidate)

	time.AfterFunc(50*time.Millisecond, func() {
		v.window.Resize(fyne.NewSize(400, 300))
	})
}

func (v *View) onUpdateLength() {
	length, _ := strconv.Atoi(v.lengthEntry.Text)
	candidate := strings.Repeat(".", length)

	gui.MakeCharLine(v.charBoxes, candidate)

	time.AfterFunc(50*time.Millisecond, func() {
		v.window.Resize(fyne.NewSize(400, 300))
	})
}

func Window() {
	myApp := app.New()
	window := myApp.NewWindow("Linear Regex Crossword")

	// Create and set the content
	view := NewView(window, "a|b", "ab", "ababa")
	window.SetContent(view.content)

	// Set a reasonable starting size
	window.Resize(fyne.NewSize(400, 300))

	// Show and run the window
	window.ShowAndRun()
}

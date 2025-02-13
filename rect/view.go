package rect

import (
	"crossmatcher/gui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	window            fyne.Window
	model             *Model
	lengthEntry       *widget.Entry
	horizontalEntries *fyne.Container
	verticalEntries   *fyne.Container
	alphabetEntry     *widget.Entry
	charBoxes         *fyne.Container
	content           *fyne.Container
}

func NewView() *View {
	v := &View{}

	width := 15
	height := 4

	charBoxes := container.NewVBox()
	for x := 0; x < width+height-1; x++ {
		row := container.NewHBox()
		for y := 0; y < width+height-1; y++ {
			box := gui.MakeCharBoxSpacer()
			row.Add(box)
		}
		charBoxes.Add(row)
	}
	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			row := charBoxes.Objects[width-1-w+h]
			if rowContainer, ok := row.(*fyne.Container); ok {
				box := gui.MakeCharBox('*')
				rowContainer.Objects[w+h] = box
			}
		}
	}

	vRules := container.NewVBox()
	for w := width - 1; w >= 0; w-- {
		leftSpace := container.NewHBox()
		rightSpace := container.NewHBox()
		for h := 0; h < w; h++ {
			box := gui.MakeCharBoxSpacer()
			leftSpace.Add(box)
		}
		for h := 0; h < height+width-w; h++ {
			box := gui.MakeCharBoxSpacer()
			rightSpace.Add(box)
		}
		entry := gui.MakeTextBox("ab|.*", "rule")
		row := container.NewHBox(leftSpace, entry, rightSpace)
		vRules.Add(row)
	}

	hRules := container.NewVBox()
	for h := 0; h < height; h++ {
		leftSpace := container.NewHBox()
		rightSpace := container.NewHBox()
		for w := 0; w < h; w++ {
			box := gui.MakeCharBoxSpacer()
			leftSpace.Add(box)
		}
		for w := 0; w < height+width-h; w++ {
			box := gui.MakeCharBoxSpacer()
			rightSpace.Add(box)
		}
		entry := gui.MakeTextBox("ab|.*", "rule")
		row := container.NewHBox(leftSpace, entry, rightSpace)
		hRules.Add(row)
	}

	candidateLayer := container.NewVBox(gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		container.NewHBox(gui.MakeTextBoxSpacer(), gui.MakeCharBoxSpacer(), charBoxes),
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer())

	ruleLayer := container.NewVBox(vRules,
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		hRules)

	v.content = container.NewStack(candidateLayer, ruleLayer)

	return v
}

func Window() {
	myApp := app.New()
	window := myApp.NewWindow("Rectangular Regex Crossword")

	// Create and set the content
	linView := NewView()
	window.SetContent(linView.content)

	// Set a reasonable starting size
	window.Resize(fyne.NewSize(400, 300))

	// Show and run the window
	window.ShowAndRun()
}

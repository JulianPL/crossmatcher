package rect

import (
	"crossmatcher/gui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type View struct {
	window            fyne.Window
	model             *Model
	widthEntry        *fyne.Container
	heightEntry       *fyne.Container
	horizontalEntries *fyne.Container
	verticalEntries   *fyne.Container
	alphabetEntry     *fyne.Container
	charBoxes         *fyne.Container
	content           *fyne.Container
}

func NewView() *View {
	v := &View{}

	width := 10
	height := 6

	vRuleStrings := make([]string, width)
	for i := 0; i < width; i++ {
		vRuleStrings[i] = strconv.Itoa(i)
	}
	hRuleStrings := make([]string, height)
	for i := 0; i < height; i++ {
		hRuleStrings[i] = strconv.Itoa(i)
	}
	candidate := make([]string, height)
	for i := 0; i < height; i++ {
		candidate[i] = "0123456789"
	}

	fullSpace := container.NewHBox(gui.MakeTextBoxSpacer(), gui.MakeCharBoxSpacerRow(height+width))

	v.widthEntry = gui.MakeTextBox(strconv.Itoa(len([]rune(candidate[0]))), "Width")
	v.heightEntry = gui.MakeTextBox(strconv.Itoa(len(candidate)), "Height")
	v.alphabetEntry = gui.MakeTextBox("abc", "Alphabet")

	vRules := gui.ReverseBox(addRuleStrings(createRuleRows(width, height), vRuleStrings))
	hRules := addRuleStrings(createRuleRows(height, width), hRuleStrings)

	vArrows := gui.ReverseBox(changeToArrowLines(createRuleRows(width, height), false))
	hArrows := changeToArrowLines(createRuleRows(height, width), true)

	v.charBoxes = gui.MakeCharBoxSpacerGrid(width+height-1, width+height-1)
	v.charBoxes = PopulateCandidateSubgrid(v.charBoxes, width, height)
	v.charBoxes = AddCandidateChars(v.charBoxes, candidate)

	ruleLayer := container.NewVBox(vRules,
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		hRules)

	arrowLayer := container.NewVBox(gui.MakeCharBoxSpacer(),
		vArrows,
		gui.MakeCharBoxSpacer(),
		hArrows,
		gui.MakeCharBoxSpacer())

	candidateLayer := container.NewVBox(gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		container.NewHBox(gui.MakeTextBoxSpacer(), gui.MakeCharBoxSpacer(), v.charBoxes),
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer())

	controlLayer := container.NewVBox(container.NewHBox(fullSpace, widget.NewLabel("Width:")),
		container.NewHBox(fullSpace, v.widthEntry),
		container.NewHBox(fullSpace, widget.NewLabel("Height:")),
		container.NewHBox(fullSpace, v.heightEntry),
		container.NewHBox(fullSpace, widget.NewLabel("Alphabet:")),
		container.NewHBox(fullSpace, v.alphabetEntry))

	v.content = container.NewStack(ruleLayer, arrowLayer, candidateLayer, controlLayer)

	return v
}

func getCandidateBox(grid *fyne.Container, row, column, width int) *fyne.CanvasObject {
	gridRow := grid.Objects[width-1-row+column]
	rowContainer, ok := gridRow.(*fyne.Container)
	if !ok {
		return nil
	}
	gridColumn := row + column
	return &rowContainer.Objects[gridColumn]
}

func PopulateCandidateSubgrid(grid *fyne.Container, width, height int) *fyne.Container {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			box := getCandidateBox(grid, w, h, width)
			*box = gui.MakeCharBox('.')
		}
	}
	return grid
}

func AddCandidateChars(grid *fyne.Container, candidate []string) *fyne.Container {
	width := len([]rune(candidate[0]))
	for h, row := range candidate {
		for w, char := range row {
			box := getCandidateBox(grid, w, h, width)
			if box, ok := (*box).(*widget.Entry); ok {
				box.SetText(string(char))
			}
		}
	}
	return grid
}

func GetCandidateChars(grid *fyne.Container, width, height int) []string {
	var ret []string
	for h := 0; h < height; h++ {
		row := ""
		for w := 0; w < width; w++ {
			box := getCandidateBox(grid, w, h, width)
			if box, ok := (*box).(*widget.Entry); ok {
				row += box.Text
			}
		}
		ret = append(ret, row)
	}
	return ret
}

func createRuleRows(dim1, dim2 int) *fyne.Container {
	rules := container.NewVBox()

	for i := 0; i < dim1; i++ {
		leftSpace := gui.MakeCharBoxSpacerRow(i)
		if i == 0 {
			leftSpace.Hide()
		}
		rightSpace := gui.MakeCharBoxSpacerRow(dim1 + dim2 - i)
		entry := gui.MakeTextBox("", "rule")
		row := container.NewHBox(leftSpace, entry, rightSpace)
		rules.Add(row)
	}

	return rules
}

func changeToArrowLines(ruleLines *fyne.Container, upward bool) *fyne.Container {
	for _, row := range ruleLines.Objects {
		if row, ok := row.(*fyne.Container); ok {
			/* A row consists of vBox(spacer), textBox, vBox(spacer)*/
			row.Objects[1] = gui.MakeTextBoxSpacer()
			vbox, _ := getSecondVBoxFromRuleLine(row)
			vbox.Objects[0] = gui.MakeCharBoxArrow(upward)
		}
	}
	return ruleLines
}

func addRuleStrings(ruleLines *fyne.Container, rules []string) *fyne.Container {
	for i, row := range ruleLines.Objects {
		if entry, ok := getEntryFromRuleLine(row); ok {
			entry.SetText(rules[i])
		}
	}
	return ruleLines
}

func getEntryFromRuleLine(line fyne.CanvasObject) (*widget.Entry, bool) {
	row, ok := line.(*fyne.Container)
	if !ok {
		return nil, false
	}
	/* A row consists of vBox(spacer), textBox, vBox(spacer)*/
	textBox, ok := row.Objects[1].(*fyne.Container)
	if !ok {
		return nil, false
	}
	/* A textbox consists of spacer, textBox */
	entry, ok := textBox.Objects[1].(*widget.Entry)
	return entry, ok
}

func getSecondVBoxFromRuleLine(line fyne.CanvasObject) (*fyne.Container, bool) {
	row, ok := line.(*fyne.Container)
	if !ok {
		return nil, false
	}
	/* A row consists of vBox(spacer), textBox, vBox(spacer)*/
	vBox, ok := row.Objects[2].(*fyne.Container)
	if !ok {
		return nil, false
	}
	return vBox, ok
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

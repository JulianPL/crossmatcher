package rect

import (
	"crossmatcher/gui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"slices"
	"strconv"
	"strings"
)

type View struct {
	window        fyne.Window
	model         *Model
	widthEntry    *fyne.Container
	heightEntry   *fyne.Container
	alphabetEntry *fyne.Container
	fullSpace     *fyne.Container
	hRules        *fyne.Container
	vRules        *fyne.Container
	hArrows       *fyne.Container
	vArrows       *fyne.Container
	charBoxes     *fyne.Container
	content       *fyne.Container
}

func NewView(window fyne.Window, vRules, hRules []string, alphabet string, candidate []string) *View {
	v := &View{}
	v.window = window

	return v.updateView(vRules, hRules, alphabet, candidate)
}

func (v *View) updateView(vRules, hRules []string, alphabet string, candidate []string) *View {
	width := len(vRules)
	height := len(hRules)

	v.fullSpace = container.NewHBox(gui.MakeTextBoxSpacer(), gui.MakeCharBoxSpacerRow(height+width+1))

	v.widthEntry = gui.MakeTextBox(strconv.Itoa(len([]rune(candidate[0]))), "Width")
	v.heightEntry = gui.MakeTextBox(strconv.Itoa(len(candidate)), "Height")
	v.alphabetEntry = gui.MakeTextBox(alphabet, "Alphabet")

	v.vRules = gui.ReverseBox(addRuleStrings(createRuleRows(width, height), vRules))
	v.hRules = addRuleStrings(createRuleRows(height, width), hRules)

	v.vArrows = gui.ReverseBox(changeToArrowLines(createRuleRows(width, height), false))
	v.hArrows = changeToArrowLines(createRuleRows(height, width), true)

	v.charBoxes = gui.MakeCharBoxSpacerGrid(width+height-1, width+height-1)
	v.charBoxes = PopulateCandidateSubgrid(v.charBoxes, width, height)
	v.charBoxes = AddCandidateChars(v.charBoxes, candidate)

	importExportButton := gui.MakeButton("Import/Export", v.onImportExport)
	updateLengthButton := gui.MakeButton("Reset Crossword and Update Length", v.onUpdateLength)
	emptyCandidateButton := gui.MakeButton("Empty Candidate", v.onEmptyCandidate)
	solveButton := gui.MakeButton("Solve", v.onSolve)

	ruleLayer := container.NewVBox(v.vRules,
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		v.hRules)

	arrowLayer := container.NewVBox(gui.MakeCharBoxSpacer(),
		v.vArrows,
		gui.MakeCharBoxSpacer(),
		v.hArrows,
		gui.MakeCharBoxSpacer())

	candidateLayer := container.NewVBox(gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer(),
		container.NewHBox(gui.MakeTextBoxSpacer(), gui.MakeCharBoxSpacer(), v.charBoxes),
		gui.MakeCharBoxSpacer(),
		gui.MakeCharBoxSpacer())

	controlLayer := container.NewVBox(container.NewHBox(v.fullSpace, importExportButton),
		container.NewHBox(v.fullSpace),
		container.NewHBox(v.fullSpace, widget.NewLabel("Width:")),
		container.NewHBox(v.fullSpace, v.widthEntry),
		container.NewHBox(v.fullSpace, widget.NewLabel("Height:")),
		container.NewHBox(v.fullSpace, v.heightEntry),
		container.NewHBox(v.fullSpace, widget.NewLabel("Alphabet:")),
		container.NewHBox(v.fullSpace, v.alphabetEntry),
		container.NewHBox(v.fullSpace),
		container.NewHBox(v.fullSpace, updateLengthButton),
		container.NewHBox(v.fullSpace),
		container.NewHBox(v.fullSpace, emptyCandidateButton),
		container.NewHBox(v.fullSpace),
		container.NewHBox(v.fullSpace, solveButton))

	v.content = container.NewStack(ruleLayer, arrowLayer, candidateLayer, controlLayer)

	return v
}

func (v *View) onImportExport() {
	textArea := widget.NewMultiLineEntry()

	width := len(v.vRules.Objects)
	height := len(v.hRules.Objects)
	alphabet, _ := gui.GetEntryText(v.alphabetEntry)
	vRules := readRuleRows(v.vRules)
	slices.Reverse(vRules)
	hRules := readRuleRows(v.hRules)
	candidate := GetCandidateChars(v.charBoxes, width, height)

	text := alphabet + "\n\n" +
		strings.Join(hRules, "\n") + "\n\n" +
		strings.Join(vRules, "\n") + "\n\n" +
		strings.Join(candidate, "\n")

	textArea.SetText(text)
	textArea.Resize(fyne.NewSize(300, 400))

	importFunc := func(importButton bool) {
		if importButton {
			v.onImport(textArea.Text)
		}
	}

	dialogWindow := dialog.NewCustomConfirm("Import/Export", "Import", "Abort", textArea, importFunc, v.window)
	dialogWindow.Resize(fyne.NewSize(300, 400))
	dialogWindow.Show()
}

func (v *View) onImport(textbox string) {
	textboxSplit := strings.Split(textbox, "\n\n")
	alphabet := textboxSplit[0]
	vRules := strings.Split(textboxSplit[2], "\n")
	hRules := strings.Split(textboxSplit[1], "\n")
	candidate := strings.Split(textboxSplit[3], "\n")

	width := len(vRules)
	height := len(hRules)

	if len(candidate) != height || len([]rune(candidate[0])) != width {
		candidate = make([]string, height)
		for i := 0; i < height; i++ {
			candidate[i] = strings.Repeat(".", width)
		}
	}

	v.updateView(vRules, hRules, alphabet, candidate)

	v.window.SetContent(v.content)
	v.window.Resize(fyne.NewSize(400, 300))
}

func (v *View) onUpdateLength() {
	widthString, _ := gui.GetEntryText(v.widthEntry)
	heightString, _ := gui.GetEntryText(v.heightEntry)
	alphabetString, _ := gui.GetEntryText(v.alphabetEntry)
	width, err := strconv.Atoi(widthString)
	if err != nil {
		return
	}
	height, err := strconv.Atoi(heightString)
	if err != nil {
		return
	}

	vRuleStrings := make([]string, width)
	for i := 0; i < width; i++ {
		vRuleStrings[i] = ""
	}
	hRuleStrings := make([]string, height)
	for i := 0; i < height; i++ {
		hRuleStrings[i] = ""
	}
	candidate := make([]string, height)
	for i := 0; i < height; i++ {
		candidate[i] = strings.Repeat(".", width)
	}

	v.updateView(vRuleStrings, hRuleStrings, alphabetString, candidate)

	v.window.SetContent(v.content)
	v.window.Resize(fyne.NewSize(400, 300))
}

func (v *View) onEmptyCandidate() {
	width := len(v.vRules.Objects)
	height := len(v.hRules.Objects)

	candidate := make([]string, height)
	for i := 0; i < height; i++ {
		candidate[i] = strings.Repeat(".", width)
	}

	AddCandidateChars(v.charBoxes, candidate)
}

func (v *View) onSolve() {
	width := len(v.vRules.Objects)
	height := len(v.hRules.Objects)
	vRules := readRuleRows(v.vRules)
	slices.Reverse(vRules)
	hRules := readRuleRows(v.hRules)
	alphabet, _ := gui.GetEntryText(v.alphabetEntry)
	candidate := GetCandidateChars(v.charBoxes, width, height)
	v.model = NewModel(vRules, hRules, alphabet, candidate)
	candidate = v.model.Solve()
	AddCandidateChars(v.charBoxes, candidate)

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

func readRuleRows(rows *fyne.Container) []string {
	ret := make([]string, len(rows.Objects))
	for i, row := range rows.Objects {
		entry, _ := getEntryFromRuleLine(row)
		ret[i] = entry.Text
	}
	return ret
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

	width := 6
	height := 6

	vRuleStrings := make([]string, width)
	for i := 0; i < width; i++ {
		vRuleStrings[i] = "00.*0"
	}
	hRuleStrings := make([]string, height)
	for i := 0; i < height; i++ {
		hRuleStrings[i] = "00.*0"
	}
	candidate := make([]string, height)
	for i := 0; i < height; i++ {
		candidate[i] = strings.Repeat(".", width)
	}

	// Create and set the content
	linView := NewView(window, vRuleStrings, hRuleStrings, "01", candidate)
	window.SetContent(linView.content)

	// Set a reasonable starting size
	window.Resize(fyne.NewSize(400, 300))

	// Show and run the window
	window.ShowAndRun()
}

// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"strings"

	"github.com/mark-summerfield/accelhint"
	"github.com/pwiecz/go-fltk"
	"golang.org/x/exp/slices"
)

func makeAccelHintsTab(app *App, x, y, width, height int) {
	group := fltk.NewFlex(x, y, width, height, "Accel &Hints")
	group.SetSpacing(pad)
	vbox := fltk.NewFlex(x, y, width, height)
	vbox.SetSpacing(pad)
	unhintedLabel, hbox := makeLabelRow(x, y, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	makeDataRow(app, x, y+buttonHeight, width, height-(3*buttonHeight))
	alphabetLabel, hbox := makeAlphabetRow(app, x, height-(2*buttonHeight),
		width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	hbox = makeStatusRow(app, x, height-buttonHeight, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	addCallbacks(app, unhintedLabel, alphabetLabel)
	app.accelTextEditor.TakeFocus()
	fltk.AddTimeout(0.0, func() { updateAccels(app) })
}

func makeLabelRow(x, y, width, height int) (*fltk.Button, *fltk.Flex) {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	unhintedLabel := makeAccelLabel(0, 0, width/2, height, "Unh&inted")
	align := fltk.ALIGN_CENTER | fltk.ALIGN_BOTTOM | fltk.ALIGN_INSIDE
	unhintedLabel.SetAlign(align)
	hintedLabel := fltk.NewBox(fltk.NO_BOX, 0, 0, width/2, height, "Hinted")
	hintedLabel.SetAlign(align)
	hbox.End()
	return unhintedLabel, hbox
}

func makeDataRow(app *App, x, y, width, height int) {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	app.accelTextEditor, app.accelTextBuffer = makeTextEditor(x, y, width/2,
		height)
	app.accelTextBuffer.SetText(app.config.LastUnhinted)
	app.accelTextEditor.SelectAll()
	app.accelView = fltk.NewHelpView(x, y, width/2, height)
	app.accelView.TextFont(fltk.HELVETICA)
	app.accelView.TextSize(app.config.ViewFontSize)
	hbox.End()
}

func makeAlphabetRow(app *App, x, y, width, height int) (*fltk.Button,
	*fltk.Flex) {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	width = (labelWidth * 5) / 4
	alphabetLabel := makeAccelLabel(0, 0, width, height, "A&lphabet")
	app.accelAlphabetInput = fltk.NewInput(0, 0, labelWidth, height)
	app.accelAlphabetInput.SetValue(defaultAlphabet)
	hbox.Fixed(alphabetLabel, width)
	hbox.End()
	return alphabetLabel, hbox
}

func makeStatusRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	app.accelStatusOutput = fltk.NewOutput(x, y, width, buttonHeight)
	const checkWidth = 2 * labelWidth
	x = width - (2 * checkWidth)
	app.accelShowLettersCheckButton = fltk.NewCheckButton(x, y,
		checkWidth, buttonHeight, "Sho&w Letters")
	app.accelShowLettersCheckButton.SetValue(app.config.AccelShowLetters)
	app.accelShowIndexesCheckButton = fltk.NewCheckButton(x, y,
		checkWidth, buttonHeight, "Show &Indexes")
	app.accelShowIndexesCheckButton.SetValue(app.config.AccelShowIndexes)
	hbox.Fixed(app.accelShowLettersCheckButton, checkWidth)
	hbox.Fixed(app.accelShowIndexesCheckButton, checkWidth)
	hbox.End()
	return hbox
}

func addCallbacks(app *App, unhintedLabel, alphabetLabel *fltk.Button) {
	unhintedLabel.SetCallback(func() { app.accelTextEditor.TakeFocus() })
	alphabetLabel.SetCallback(func() { app.accelAlphabetInput.TakeFocus() })
	app.accelTextEditor.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	app.accelTextEditor.SetCallback(func() { updateAccels(app) })
	app.accelAlphabetInput.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	app.accelAlphabetInput.SetCallback(func() { updateAccels(app) })
	app.accelShowLettersCheckButton.SetCallback(
		func() {
			app.config.AccelShowLetters =
				app.accelShowLettersCheckButton.Value()
			updateAccels(app)
		})
	app.accelShowIndexesCheckButton.SetCallback(
		func() {
			app.config.AccelShowIndexes =
				app.accelShowIndexesCheckButton.Value()
			updateAccels(app)
		})
}

func updateAccels(app *App) {
	alphabet := getAlphabet(app.accelAlphabetInput)
	unhinted := []string{}
	for _, item := range strings.Split(app.accelTextBuffer.Text(), "\n") {
		item = strings.TrimSpace(item)
		if item != "" {
			unhinted = append(unhinted, item)
		}
	}
	hinted, n, err := accelhint.HintedX(unhinted, '&', alphabet)
	if err != nil {
		app.accelView.SetValue("")
		app.accelStatusOutput.SetColor(fltk.RED)
		app.accelStatusOutput.SetValue(fmt.Sprintf(
			"Failed to set accelerators: %s", err))
	} else if len(unhinted) == 0 {
		app.accelStatusOutput.SetColor(fltk.CYAN)
		app.accelStatusOutput.SetValue("Enter items…")
	} else {
		showLetters := app.accelShowLettersCheckButton.Value()
		showIndexes := app.accelShowIndexesCheckButton.Value()
		app.accelView.SetValue(getHintHtml(unhinted, hinted, showLetters,
			showIndexes))
		app.accelStatusOutput.SetColor(fltk.YELLOW)
		app.accelStatusOutput.SetValue(fmt.Sprintf(
			"Hinted — %d/%d — %.0f%%", n, len(unhinted),
			(float64(n) / float64(len(unhinted)) * 100.0)))
	}
	app.Window.Redraw()
}

func getAlphabet(accelAlphabetInput *fltk.Input) string {
	alphabet := strings.ToUpper(strings.TrimSpace(
		accelAlphabetInput.Value()))
	if len(alphabet) < 5 {
		alphabet = defaultAlphabet
		accelAlphabetInput.SetValue(alphabet)
	}
	return alphabet
}

func getHintHtml(unhinted, hinted []string, showLetters,
	showIndexes bool) string {
	var text strings.Builder
	for i := 0; i < len(hinted); i++ {
		chars := []rune(strings.ReplaceAll(hinted[i], escMarker,
			placeholder))
		j := slices.Index(chars, rune('&'))
		if j > -1 && j+1 < len(chars) {
			left := strings.ReplaceAll(string(chars[:j]), placeholder,
				escMarker)
			accel := string(chars[j+1])
			right := ""
			if j+2 < len(chars) {
				right = strings.ReplaceAll(string(chars[j+2:]),
					placeholder, escMarker)
			}
			if showLetters {
				text.WriteString(fmt.Sprintf(
					"<font face=courier color=green><b>%s</b></font>"+
						"&nbsp;&nbsp;&nbsp;", strings.ToUpper(accel)))
			}
			if showIndexes {
				text.WriteString(fmt.Sprintf(
					"<font color=gray>%d</font>&nbsp;&nbsp;&nbsp;", j))
			}
			text.WriteString(left)
			text.WriteString(fmt.Sprintf(
				"<font color=blue><u>%s</u></font>", accel))
			text.WriteString(right)
		} else {
			if showLetters {
				text.WriteString("<font face=courier><b>&nbsp;</b></font>" +
					"&nbsp;&nbsp;&nbsp;")
			}
			if showIndexes {
				text.WriteString("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;")
			}
			text.WriteString(strings.ReplaceAll(string(chars), placeholder,
				escMarker))

		}
		text.WriteString("<br>")
	}
	return text.String()
}

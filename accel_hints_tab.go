// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeAccelHintsTab(app *App, x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "Accel &Hints")
	vbox := fltk.NewFlex(x, y, width, height)
	unhintedLabel, hbox := makeLabelRow(x, y, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	makeDataRow(app, x, y+buttonHeight, width, height-(3*buttonHeight))
	alphabetLabel, hbox := makeAlphabetRow(app, x, height-(2*buttonHeight),
		width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	app.accelStatusOutput = fltk.NewOutput(x, height-buttonHeight, width,
		buttonHeight)
	app.accelStatusOutput.SetValue("Ready")
	app.accelStatusOutput.SetBox(fltk.DOWN_FRAME)
	vbox.Fixed(app.accelStatusOutput, buttonHeight)
	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	unhintedLabel.SetCallback(func() { app.accelTextEditor.TakeFocus() })
	alphabetLabel.SetCallback(func() { app.accelAlphabetInput.TakeFocus() })
	app.accelTextEditor.TakeFocus()
}

func makeLabelRow(x, y, width, height int) (*fltk.Button, *fltk.Flex) {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	unhintedLabel := makeAccelLabel(0, 0, width/2, height, "&Unhinted")
	unhintedLabel.SetAlign(fltk.ALIGN_CENTER)
	fltk.NewBox(fltk.NO_BOX, 0, 0, width/2, height, "Hinted")
	hbox.End()
	return unhintedLabel, hbox
}

func makeDataRow(app *App, x, y, width, height int) {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	app.accelTextEditor, app.accelTextBuffer = makeTextEditor(x, y, width/2,
		height)
	app.accelView = fltk.NewHelpView(x, y, width/2, height)
	app.accelView.TextFont(fltk.HELVETICA)
	app.accelView.TextSize(app.config.ViewFontSize)
	hbox.End()
}

func makeAlphabetRow(app *App, x, y, width, height int) (*fltk.Button,
	*fltk.Flex) {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	width = (labelWidth * 5) / 4
	alphabetLabel := makeAccelLabel(0, 0, width, height, "A&lphabet")
	app.accelAlphabetInput = fltk.NewInput(0, 0, labelWidth, height)
	app.accelAlphabetInput.SetValue("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	hbox.Fixed(alphabetLabel, width)
	hbox.End()
	return alphabetLabel, hbox
}

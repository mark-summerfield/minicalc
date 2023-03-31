// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"strings"

	"github.com/pwiecz/go-fltk"
)

const optionsLabelWidth = labelWidth * 2

func makeOptionsTab(app *App, x, y, width, height int) {
	yoffset := 2 * buttonHeight
	group := fltk.NewGroup(x, y, width, height, "&Options")
	vbox := fltk.NewFlex(x, y, width, height)
	hbox := makeScaleRow(app, x, yoffset, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	hbox = makeThemeRow(app, x, y, yoffset, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	hbox = makeFontSizeRow(app, x, y, yoffset, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	hbox = makeHelpCheckboxRow(app, x, y, yoffset, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	hbox = makeCustomTitleRow(app, x, y, yoffset, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	button := makeCustomTextRows(app, x, y, yoffset, buttonHeight)
	vbox.Fixed(button, buttonHeight)
	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	app.scaleSpinner.TakeFocus()
}

func makeScaleRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	scaleLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight, "&Scale")
	app.scaleSpinner = fltk.NewSpinner(0, 0, labelWidth, buttonHeight)
	app.scaleSpinner.SetType(fltk.SPINNER_FLOAT_INPUT)
	app.scaleSpinner.SetMinimum(0.5)
	app.scaleSpinner.SetMaximum(3.5)
	app.scaleSpinner.SetStep(0.1)
	app.scaleSpinner.SetValue(float64(app.config.Scale))
	app.scaleSpinner.SetCallback(func() {
		fltk.SetScreenScale(0, float32(app.scaleSpinner.Value()))
	})
	scaleLabel.SetCallback(func() { app.scaleSpinner.TakeFocus() })
	hbox.Fixed(scaleLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeThemeRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	themeLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight, "The&me")
	app.themeChoice = fltk.NewChoice(0, 0, optionsLabelWidth, buttonHeight)
	for i, theme := range themes {
		theme := theme
		if theme == app.config.Theme {
			app.themeChoice.SetValue(i)
		}
		app.themeChoice.AddEx(theme, 0, func() {
			app.config.Theme = theme
			fltk.SetScheme(theme)
		}, 0)
	}
	themeLabel.SetCallback(func() { app.themeChoice.TakeFocus() })
	hbox.Fixed(themeLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeFontSizeRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	sizeLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight,
		"View &Font Size")
	app.sizeSpinner = fltk.NewSpinner(0, 0, labelWidth, buttonHeight)
	app.sizeSpinner.SetType(fltk.SPINNER_INT_INPUT)
	app.sizeSpinner.SetMinimum(10)
	app.sizeSpinner.SetMaximum(20)
	app.sizeSpinner.SetValue(float64(app.config.ViewFontSize))
	app.sizeSpinner.SetCallback(func() {
		size := int(app.sizeSpinner.Value())
		app.config.ViewFontSize = size
		fmt.Println(size)
		for _, widget := range []*fltk.HelpView{app.evalView, app.regexView,
			app.accelView, app.unicodeView, app.customView, app.aboutView} {
			if widget != nil {
				widget.TextSize(size)
			}
		}
	})
	sizeLabel.SetCallback(func() { app.sizeSpinner.TakeFocus() })
	hbox.Fixed(sizeLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeHelpCheckboxRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	label := fltk.NewBox(fltk.NO_BOX, 0, 0, labelWidth, buttonHeight)
	app.showInitialHelpCheckButton = fltk.NewCheckButton(x, y, width,
		buttonHeight, "Show &Initial Help Text")
	app.showInitialHelpCheckButton.SetValue(app.config.ShowIntialHelpText)
	hbox.Fixed(label, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeCustomTitleRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	titleLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight,
		"&Custom Title")
	app.customTitleInput = fltk.NewInput(0, 0, labelWidth, buttonHeight)
	app.customTitleInput.SetValue(app.config.CustomTitle)
	app.customTitleInput.SetCallback(func() {
		app.customGroup.SetLabel(app.customTitleInput.Value())
	})
	titleLabel.SetCallback(func() { app.customTitleInput.TakeFocus() })
	hbox.Fixed(titleLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeCustomTextRows(app *App, x, y, width, height int) *fltk.Button {
	textLabel := makeAccelLabel(x, y, labelWidth, buttonHeight,
		"Custom &Text:")
	app.customTextEditor, app.customTextBuffer = makeTextEditor(x, y, width,
		height*5)
	app.customTextBuffer.SetText(app.config.CustomHtml)
	app.customTextEditor.SetCallback(func() {
		app.customView.SetValue(strings.ReplaceAll(
			app.customTextBuffer.Text(), "\"\"\"", "&quot;&quot;&quot;"))
	})
	textLabel.SetCallback(func() { app.customTextEditor.TakeFocus() })
	return textLabel
}

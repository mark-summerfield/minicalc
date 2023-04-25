// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"strings"

	"github.com/pwiecz/go-fltk"
)

func makeConversionTab(app *App, x, y, width, height int) {
	yoffset := 2 * buttonHeight
	group := fltk.NewFlex(x, y, width, height, "&Conversion")
	group.SetSpacing(pad)
	vbox := fltk.NewFlex(x, y, width, height)
	vbox.SetSpacing(pad)
	hbox := makeFromUnitRow(app, x, yoffset, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	hbox = makeToUnitRow(app, x, yoffset, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	hbox = makeAmountRow(app, x, yoffset, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	yoffset += buttonHeight
	hbox = makeResultRow(app, x, y, yoffset, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	addUnits(app)
	app.convFromChoice.TakeFocus()
}

func makeFromUnitRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	fromLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight, "&From")
	app.convFromChoice = fltk.NewChoice(0, 0, optionsLabelWidth,
		buttonHeight)
	fromLabel.SetCallback(func() { app.convFromChoice.TakeFocus() })
	hbox.Fixed(fromLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeToUnitRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	toLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight, "&To")
	app.convToChoice = fltk.NewChoice(0, 0, optionsLabelWidth,
		buttonHeight)
	toLabel.SetCallback(func() { app.convToChoice.TakeFocus() })
	hbox.Fixed(toLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeAmountRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	amountLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight, "&Amount")
	app.convAmountSpinner = fltk.NewSpinner(0, 0, labelWidth, buttonHeight)
	app.convAmountSpinner.SetType(fltk.SPINNER_FLOAT_INPUT)
	app.convAmountSpinner.SetMinimum(0)
	app.convAmountSpinner.SetMaximum(1e9)
	app.convAmountSpinner.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	app.convAmountSpinner.SetCallback(func() { onConvert(app) })
	amountLabel.SetCallback(func() { app.convAmountSpinner.TakeFocus() })
	hbox.Fixed(amountLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func makeResultRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	align := fltk.ALIGN_LEFT | fltk.ALIGN_INSIDE
	resultLabelLabel := fltk.NewBox(fltk.NO_BOX, 0, 0, labelWidth,
		buttonHeight, "Result")
	resultLabelLabel.SetAlign(align)
	app.convResultOutput = fltk.NewOutput(labelWidth, 0, labelWidth,
		buttonHeight)
	app.convResultOutput.SetAlign(align)
	hbox.Fixed(resultLabelLabel, optionsLabelWidth)
	hbox.End()
	return hbox
}

func addUnits(app *App) {
	callback := func() { onConvert(app) }
	for _, name := range []string{"&Feet", "&Inches", "&Meters",
		"Mi&llimeters", "&Points"} {
		app.convFromChoice.Add(name, callback)
		app.convToChoice.Add(name, callback)
	}
	app.convFromChoice.SetValue(app.convFromChoice.FindIndex("&Feet"))
	app.convToChoice.SetValue(app.convFromChoice.FindIndex("&Meters"))
	onConvert(app)
}

func onConvert(app *App) {
	fromUnit := strings.ReplaceAll(app.convFromChoice.SelectedText(), "&",
		"")
	toUnit := strings.ReplaceAll(app.convToChoice.SelectedText(), "&", "")
	ok := true
	factor := 0.0
	switch fromUnit {
	case "Feet":
		switch toUnit {
		case "Feet":
			factor = 1
		case "Inches":
			factor = 12
		case "Meters":
			factor = 0.3048
		case "Millimeters":
			factor = 304.8
		case "Points":
			factor = 12 * 72
		default:
			ok = false
		}
	case "Inches":
		switch toUnit {
		case "Feet":
			factor = 1.0 / 12.0
		case "Inches":
			factor = 1
		case "Meters":
			factor = 0.0254
		case "Millimeters":
			factor = 25.4
		case "Points":
			factor = 72
		default:
			ok = false
		}
	case "Meters":
		switch toUnit {
		case "Feet":
			factor = 3.28084
		case "Inches":
			factor = 39.3701
		case "Meters":
			factor = 1
		case "Millimeters":
			factor = 1000
		case "Points":
			factor = 39.3701 * 72
		default:
			ok = false
		}
	case "Millimeters":
		switch toUnit {
		case "Feet":
			factor = 0.00328084
		case "Inches":
			factor = 0.0393701
		case "Meters":
			factor = 1.0 / 1000.0
		case "Millimeters":
			factor = 1
		case "Points":
			factor = 0.0393701 * 72
		default:
			ok = false
		}
	case "Points":
		switch toUnit {
		case "Feet":
			factor = 1.0 / (12.0 * 72.0)
		case "Inches":
			factor = 1.0 / 72.0
		case "Meters":
			factor = 1.0 / (39.3701 * 72.0)
		case "Millimeters":
			factor = 1.0 / (39.3701 * 72.0 / 1000.0)
		case "Points":
			factor = 1
		default:
			ok = false
		}
	}
	if ok {
		result := app.convAmountSpinner.Value() * factor
		app.convResultOutput.SetColor(fltk.YELLOW)
		app.convResultOutput.SetValue(fmt.Sprintf("%v", result))
	} else {
		app.convResultOutput.SetColor(fltk.RED)
		app.convResultOutput.SetValue(fmt.Sprintf(
			"Can't convert from %s to %s", fromUnit, toUnit))
	}
}

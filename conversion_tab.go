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
	factor, ok := factorForUnits[Units{fromUnit, toUnit}]
	if ok {
		result := factor * app.convAmountSpinner.Value()
		app.convResultOutput.SetColor(fltk.YELLOW)
		app.convResultOutput.SetValue(fmt.Sprintf("%v", result))
	} else {
		app.convResultOutput.SetColor(fltk.RED)
		app.convResultOutput.SetValue(fmt.Sprintf(
			"Can't convert from %s to %s", fromUnit, toUnit))
	}
}

type Units struct {
	from string
	to   string
}

var factorForUnits = map[Units]float64{
	{feet, feet}:               1,
	{feet, inches}:             12,
	{feet, meters}:             0.3048,
	{feet, millimeters}:        304.8,
	{feet, points}:             12 * 72,
	{inches, feet}:             1.0 / 12.0,
	{inches, meters}:           0.0254,
	{inches, millimeters}:      25.4,
	{inches, points}:           27,
	{meters, feet}:             3.28084,
	{meters, inches}:           39.3701,
	{meters, meters}:           1,
	{meters, millimeters}:      1000,
	{meters, points}:           39.3701 * 72,
	{millimeters, feet}:        0.00328084,
	{millimeters, inches}:      0.0393701,
	{millimeters, meters}:      1.0 / 1000.0,
	{millimeters, millimeters}: 1,
	{millimeters, points}:      0.0393701 * 72,
	{points, feet}:             1.0 / (12.0 * 72.0),
	{points, inches}:           1.0 / 72.0,
	{points, meters}:           1.0 / (39.3701 * 72.0),
	{points, millimeters}:      1.0 / (39.3701 * 72.0 / 1000.0),
	{points, points}:           1,
}

const (
	feet        = "Feet"
	inches      = "Inches"
	meters      = "Meters"
	millimeters = "Millimeters"
	points      = "Points"
)

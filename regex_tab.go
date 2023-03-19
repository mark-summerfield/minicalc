// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeRegexTab(app *App, x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "&2 Regex")
	vbox := fltk.NewPack(x, y, width, height)
	hoffset := 2 * BUTTON_HEIGHT
	regexView := fltk.NewHelpView(x, y, width, height-hoffset)
	regexView.SetValue(REGEX_HELP_HTML)

	hbox := fltk.NewPack(x, height-hoffset, width, BUTTON_HEIGHT)
	hbox.SetType(fltk.HORIZONTAL)
	regexLabel := makeAccelLabel(0, 0, LABEL_WIDTH, BUTTON_HEIGHT, "&Regex")
	app.regexInput = fltk.NewInput(0, BUTTON_HEIGHT, width-LABEL_WIDTH,
		BUTTON_HEIGHT)
	regexLabel.SetCallback(func() { app.regexInput.TakeFocus() })
	app.regexInput.SetCallbackCondition(fltk.WhenEnterKey)
	// TODO SetCallback
	hbox.End()

	hbox = fltk.NewPack(x, height-BUTTON_HEIGHT, width, BUTTON_HEIGHT)
	hbox.SetType(fltk.HORIZONTAL)
	textLabel := makeAccelLabel(0, 0, LABEL_WIDTH, BUTTON_HEIGHT, "&Text")
	textInput := fltk.NewInput(0, BUTTON_HEIGHT, width-LABEL_WIDTH,
		BUTTON_HEIGHT)
	textLabel.SetCallback(func() { textInput.TakeFocus() })
	textInput.SetCallbackCondition(fltk.WhenEnterKey)
	// TODO SetCallback
	hbox.End()

	vbox.End()
	vbox.Resizable(regexView) // TODO Doesn't work: need Flex
	group.End()
	group.Resizable(vbox)
	group.End()
	app.regexInput.TakeFocus()
}

const REGEX_HELP_HTML = `<p><font face=sans size=4>Type a regular
expression and some text to test it on and press Enter.</font></p>` // TODO complete

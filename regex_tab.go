// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pwiecz/go-fltk"
)

func makeRegexTab(app *App, x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "&2 Regex")
	vbox := fltk.NewPack(x, y, width, height)
	hoffset := 2 * BUTTON_HEIGHT
	regexView := fltk.NewHelpView(x, y, width, height-hoffset)
	if app.config.ShowIntialHelpText {
		regexView.SetValue(REGEX_HELP_HTML)
	}

	hbox := fltk.NewPack(x, height-hoffset, width, BUTTON_HEIGHT)
	hbox.SetType(fltk.HORIZONTAL)
	regexLabel := makeAccelLabel(0, 0, LABEL_WIDTH, BUTTON_HEIGHT, "&Regex")
	app.regexInput = fltk.NewInput(0, BUTTON_HEIGHT, width-LABEL_WIDTH,
		BUTTON_HEIGHT)
	app.regexInput.SetValue(`^\s*(\S+)\s*[=:]\s*(.*)$`)
	regexLabel.SetCallback(func() { app.regexInput.TakeFocus() })
	app.regexInput.SetCallbackCondition(fltk.WhenEnterKey)
	hbox.End()

	hbox = fltk.NewPack(x, height-BUTTON_HEIGHT, width, BUTTON_HEIGHT)
	hbox.SetType(fltk.HORIZONTAL)
	textLabel := makeAccelLabel(0, 0, LABEL_WIDTH, BUTTON_HEIGHT, "&Text")
	textInput := fltk.NewInput(0, BUTTON_HEIGHT, width-LABEL_WIDTH,
		BUTTON_HEIGHT)
	textInput.SetValue("scale: 1.15")
	textLabel.SetCallback(func() { textInput.TakeFocus() })
	textInput.SetCallbackCondition(fltk.WhenEnterKey)
	hbox.End()

	vbox.End()
	vbox.Resizable(regexView) // TODO Doesn't work: need Flex
	group.End()
	group.Resizable(vbox)
	group.End()
	callback := func() { onRegex(app.regexInput, textInput, regexView) }
	app.regexInput.SetCallback(callback)
	textInput.SetCallback(callback)
	onRegex(app.regexInput, textInput, regexView)
	app.regexInput.TakeFocus()
}

func onRegex(regexInput, textInput *fltk.Input, regexView *fltk.HelpView) {
	regex := regexInput.Value()
	text := textInput.Value()
	if regex != "" {
		rx, err := regexp.Compile(regex)
		if err != nil {
			regexView.SetValue(fmt.Sprintf(errTemplate, err))
		} else {
			empty := true
			var textBuilder strings.Builder
			textBuilder.WriteString("<font face=sans size=4>")
			if rx.MatchString(text) {
				textBuilder.WriteString(
					"<tt>MatchString(text)</tt> → true<br>")
				empty = false
			}
			header := false
			for i, match := range rx.FindAllString(text, -1) {
				if !header {
					textBuilder.WriteString(
						"<tt>FindAllString(text, -1)</tt><br>")
					header = true
					empty = false
				}
				textBuilder.WriteString(fmt.Sprintf(
					"&nbsp;&nbsp;Found #%d: <tt>%q</tt><br>", i, match))
			}
			header = false
			for i, matches := range rx.FindAllStringSubmatch(text, -1) {
				if !header {
					textBuilder.WriteString(
						"<tt>FindAllStringSubmatch(text, -1)</tt><br>")
					header = true
					empty = false
				}
				textBuilder.WriteString(fmt.Sprintf(
					"&nbsp;&nbsp;Submatch #%d<br>", i))
				for j, match := range matches {
					textBuilder.WriteString(fmt.Sprintf(
						`&nbsp;&nbsp;&nbsp;&nbsp;Found #%d:
						<tt>%q</tt><br>`, j, match))
				}
			}
			textBuilder.WriteString("</font>")
			output := textBuilder.String()
			if empty {
				output = `<font face=sans size=4>Valid regex does
				<i>not</i> match the text.</font>`
			}
			regexView.SetValue(output)
		}
	}
}

const REGEX_HELP_HTML = `<p><font face=sans size=4>Type a regular
expression and some text to test it on and press Enter.</font></p>` // TODO complete

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
	group := fltk.NewFlex(x, y, width, height, "Rege&x")
	vbox := fltk.NewFlex(x, y, width, height)
	hoffset := 2 * BUTTON_HEIGHT
	app.regexView = fltk.NewHelpView(x, y, width, height-hoffset)
	if app.config.ShowIntialHelpText {
		app.regexView.SetValue(regexHelpHtml)
	}

	hbox := fltk.NewFlex(x, height-hoffset, width, BUTTON_HEIGHT)
	hbox.SetType(fltk.ROW)
	regexLabel := makeAccelLabel(0, 0, LABEL_WIDTH, BUTTON_HEIGHT, "&Regex")
	app.regexInput = fltk.NewInput(0, BUTTON_HEIGHT, width-LABEL_WIDTH,
		BUTTON_HEIGHT)
	app.regexInput.SetValue(`\s*(\S+)\s*[=:]\s*(\S+)`)
	regexLabel.SetCallback(func() { app.regexInput.TakeFocus() })
	app.regexInput.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	hbox.End()
	hbox.Fixed(regexLabel, LABEL_WIDTH)
	vbox.Fixed(hbox, BUTTON_HEIGHT)

	hbox = fltk.NewFlex(x, height-BUTTON_HEIGHT, width, BUTTON_HEIGHT)
	hbox.SetType(fltk.ROW)
	textLabel := makeAccelLabel(0, 0, LABEL_WIDTH, BUTTON_HEIGHT, "&Text")
	textInput := fltk.NewInput(0, BUTTON_HEIGHT, width-LABEL_WIDTH,
		BUTTON_HEIGHT)
	textInput.SetValue("scale: 1.15 width=24.5")
	textLabel.SetCallback(func() { textInput.TakeFocus() })
	textInput.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	hbox.End()
	hbox.Fixed(textLabel, LABEL_WIDTH)
	vbox.Fixed(hbox, BUTTON_HEIGHT)

	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	callback := func() { onRegex(app, textInput) }
	app.regexInput.SetCallback(callback)
	textInput.SetCallback(callback)
	if !app.config.ShowIntialHelpText {
		onRegex(app, textInput)
	}
	app.regexInput.TakeFocus()
}

func onRegex(app *App, textInput *fltk.Input) {
	regex := app.regexInput.Value()
	text := textInput.Value()
	if regex != "" {
		rx, err := regexp.Compile(regex)
		if err != nil {
			app.regexView.SetValue(fmt.Sprintf(errTemplate, err))
		} else {
			empty := true
			var textBuilder strings.Builder
			textBuilder.WriteString("<font color=green face=sans size=4>")
			if rx.MatchString(text) {
				textBuilder.WriteString(`<tt>MatchString(text)</tt> → <font
					color=blue>true</font><br>`)
				empty = false
			}
			match := rx.FindString(text)
			if match != "" {
				textBuilder.WriteString(fmt.Sprintf(
					`<tt>FindString(text) → <font color=blue>%q
					</font></tt><br>`, match))
				empty = false
			}
			header := false
			for i, match := range rx.FindAllString(text, -1) {
				if !header {
					textBuilder.WriteString(`<tt>FindAllString(text,
					-1)</tt> → <tt>[]string</tt><br>`)
					header = true
					empty = false
				}
				textBuilder.WriteString(fmt.Sprintf(
					`&nbsp;&nbsp;&nbsp;&nbsp;[%d] =
					<font color=blue><tt>%q</tt></font><br>`, i, match))
			}
			header = false
			for i, matches := range rx.FindAllStringSubmatch(text, -1) {
				if !header {
					textBuilder.WriteString(`<tt>FindAllStringSubmatch(text,
					-1)</tt> → <tt>[][]string</tt><br>`)
					header = true
					empty = false
				}
				for j, match := range matches {
					textBuilder.WriteString(fmt.Sprintf(
						`&nbsp;&nbsp;&nbsp;&nbsp;[%d][%d] =
						<font color=blue><tt>%q</tt></font><br>`, i, j,
						match))
				}
			}
			textBuilder.WriteString("</font>")
			output := textBuilder.String()
			if empty {
				output = `<font color=navy face=sans size=4>Valid regex does
				<i>not</i> match the text.</font>`
			}
			app.regexView.SetValue(output)
		}
	}
}

const regexHelpHtml = `<p><font face=sans size=4>Type a regular
expression and some text to test it on and press Enter.</font></p>`

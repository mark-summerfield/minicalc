// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"regexp"
	"strings"

	"github.com/pwiecz/go-fltk"
)

func makeRegexTab(app *App, x, y, width, height int) {
	group := fltk.NewFlex(x, y, width, height, "Rege&x")
	group.SetSpacing(pad)
	vbox := fltk.NewFlex(x, y, width, height)
	vbox.SetSpacing(pad)
	const yoffset = 2 * buttonHeight
	app.regexView = fltk.NewHelpView(x, y, width, height-yoffset)
	app.regexView.TextFont(fltk.COURIER)
	app.regexView.TextSize(app.config.ViewFontSize)
	if app.config.ShowIntialHelpText {
		app.regexView.SetValue(regexHelpHtml)
	}
	hbox := makeRegexRow(app, x, y, width, height, yoffset)
	vbox.Fixed(hbox, buttonHeight)
	hbox = makeTextRow(app, x, y, width, height, yoffset)
	vbox.Fixed(hbox, buttonHeight)
	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	app.regexInput.SetCallback(func() {
		updateInputChoice(app.regexInput)
		onRegex(app)
	})
	app.regexTextInput.SetCallback(func() {
		updateInputChoice(app.regexTextInput)
		onRegex(app)
	})
	if !app.config.ShowIntialHelpText {
		onRegex(app)
	}
	app.regexInput.TakeFocus()
}

func makeRegexRow(app *App, x, y, width, height, yoffset int) *fltk.Flex {
	hbox := fltk.NewFlex(x, height-yoffset, width, buttonHeight)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	regexLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight, "&Regex")
	app.regexInput = fltk.NewInputChoice(0, buttonHeight,
		width-labelWidth, buttonHeight)
	app.regexInput.Input().SetValue(
		`\s*(?P<key>\S+)\s*[=:]\s*(?P<value>\S+)`)
	regexLabel.SetCallback(func() { app.regexInput.TakeFocus() })
	app.regexInput.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	hbox.End()
	hbox.Fixed(regexLabel, labelWidth)
	return hbox
}

func makeTextRow(app *App, x, y, width, height, yoffset int) *fltk.Flex {
	hbox := fltk.NewFlex(x, height-buttonHeight, width, buttonHeight)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	textLabel := makeAccelLabel(0, 0, labelWidth, buttonHeight, "&Text")
	app.regexTextInput = fltk.NewInputChoice(0, buttonHeight,
		width-labelWidth, buttonHeight)
	text := "scale: 1.15 width=24.5"
	app.regexTextInput.SetValue(text)
	app.regexTextInput.MenuButton().Add(text,
		func() { app.regexTextInput.Input().SetValue(text) })
	textLabel.SetCallback(func() { app.regexTextInput.TakeFocus() })
	app.regexTextInput.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	hbox.End()
	hbox.Fixed(textLabel, labelWidth)
	return hbox
}

func onRegex(app *App) {
	regex := app.regexInput.Value()
	text := app.regexTextInput.Value()
	if regex != "" {
		rx, err := regexp.Compile(regex)
		if err != nil {
			app.regexView.SetValue(fmt.Sprintf(errTemplate, err))
		} else {
			empty := true
			var textBuilder strings.Builder
			textBuilder.WriteString("<font color=green>")
			if rx.MatchString(text) {
				textBuilder.WriteString(`<tt>MatchString(text)</tt> → <font
					color=blue>true</font><br>`)
				empty = false
			}
			match := rx.FindString(text)
			if match != "" {
				textBuilder.WriteString(fmt.Sprintf(
					`<tt>FindString(text) → <font color=blue>%q
					</font></tt><br>`, html.EscapeString(match)))
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
					<font color=blue><tt>%q</tt></font><br>`, i,
					html.EscapeString(match)))
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
						html.EscapeString(match)))
				}
			}
			for i, name := range rx.SubexpNames() {
				if i == 0 || name == "" {
					continue
				}
				match := rx.FindStringSubmatch(text)
				if len(match) == 0 {
					continue
				}
				textBuilder.WriteString(fmt.Sprintf(`<tt>SubexpNames()[%d]
				</tt> → <font color=blue><tt>%q</tt></font>
				<br>&nbsp;&nbsp;&nbsp;&nbsp;
				<tt>FindStringSubmatch(text)[%d]</tt> →
				<font color=blue><tt>%q</tt></font><br>`, i, name, i,
					html.EscapeString(match[i])))
				empty = false
			}
			textBuilder.WriteString("</font>")
			output := textBuilder.String()
			if empty {
				output = `<font color=navy>Valid regex does <i>not</i> 
				match the text.</font>`
			}
			app.regexView.SetValue(output)
		}
	}
}

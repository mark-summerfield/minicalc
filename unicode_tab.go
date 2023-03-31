// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/pwiecz/go-fltk"
)

func makeAsciiTab(app *App, x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "&Unicode")
	vbox := fltk.NewFlex(x, y, width, height)
	hbox := makeTopRow(x, y, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	y += buttonHeight
	hbox = makeChoiceRow(x, y, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	y += buttonHeight
	height -= (2 * buttonHeight)
	app.unicodeView = fltk.NewHelpView(x, y, width, height)
	app.unicodeView.TextFont(fltk.COURIER)
	app.unicodeView.TextSize(app.config.ViewFontSize)
	app.unicodeView.SetValue(asciiHtml())
	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	app.unicodeView.TakeFocus() // TODO change + in app.go
}

func makeTopRow(x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	colWidth := (labelWidth * 3) / 2
	cpLabel := makeAccelLabel(0, 0, colWidth, buttonHeight, "Code &Point")
	cpInput := fltk.NewInput(colWidth, 0, colWidth, buttonHeight)
	cpLabel.SetCallback(func() { cpInput.TakeFocus() })
	cpInput.SetCallbackCondition(fltk.WhenEnterKeyChanged)
	cpOutput := fltk.NewOutput(2*colWidth, 0, colWidth, buttonHeight)
	cpOutput.ClearVisibleFocus()
	cpInput.SetCallback(func() {
		text, color := cpHtml(cpInput.Value())
		cpOutput.SetColor(color)
		cpOutput.SetValue(text)
	})
	cpCopyButton := fltk.NewButton(4*colWidth, 0, labelWidth, buttonHeight,
		"&Copy")
	cpCopyButton.SetCallback(func() {
		for _, c := range cpOutput.Value() {
			fltk.CopyToClipboard(string(c))
			break
		}
	})
	hbox.End()
	hbox.Fixed(cpLabel, colWidth)
	hbox.Fixed(cpInput, colWidth)
	hbox.Fixed(cpCopyButton, labelWidth)
	return hbox
}

func makeChoiceRow(x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	// TODO
	// &Category [Choice v] // ASCII (<32) | ASCII (>=32) | unicode cats...
	hbox.End()
	//hbox.Fixed(regexLabel, labelWidth)
	return hbox
}

func cpHtml(s string) (string, fltk.Color) {
	s = strings.TrimSpace(s)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i, err = strconv.ParseInt(s, 16, 64)
		if err != nil {
			return "integer required", fltk.RED
		}
	}
	return fmt.Sprintf("%c  %U  %d", i, i, i), fltk.YELLOW
}

func asciiHtml() string {
	var text strings.Builder
	populateAsciiHigh(&text)
	populateAsciiLow(&text)
	return text.String()
}

func populateAsciiHigh(text *strings.Builder) {
	text.WriteString("<p>")
	const (
		start  = 33
		end    = 127
		step   = 5
		stride = (end - start + 1) / step
	)
	for i := 0; i < stride; i++ {
		text.WriteString("&nbsp;")
		j := start + i
		for k := 0; k < 5; k++ {
			populateOne(text, rune(j+(k*stride)), k == 4)
		}
	}
	text.WriteString("</p>")
}

func populateOne(text *strings.Builder, i rune, isEnd bool) {
	if i == 127 {
		text.WriteString("7F backspace<br>")
		return
	}
	var end string
	if isEnd {
		end = "<br>"
	} else {
		end = " <font color=#aaa>|</font> "
	}
	text.WriteString(fmt.Sprintf("%02X <font color=navy>%s</font>%s", i,
		html.EscapeString(string(i)), end))
}

func populateAsciiLow(text *strings.Builder) {
	text.WriteString("<p>")
	descForChar := getDescForChar()
	for i := 0; i < len(descForChar); i++ {
		charDesc := descForChar[i]
		c := charDesc.char
		if c == "�" {
			c += "&nbsp;"
		}
		istr := fmt.Sprintf("%d", i)
		if len(istr) == 1 {
			istr = "&nbsp;&nbsp;" + istr
		} else if len(istr) == 2 {
			istr = "&nbsp;" + istr
		}
		text.WriteString(fmt.Sprintf(
			`&nbsp;%02X&nbsp;<font color=navy>%s</font>
			&nbsp;%s&nbsp;%s<br>`, i, c, istr, charDesc.desc))
	}
	text.WriteString("</p>")
}

type CharDesc struct {
	char string
	desc string
}

type DescForCharMap map[int]CharDesc

func getDescForChar() DescForCharMap {
	descForChar := make(DescForCharMap)
	for i, charDesc := range []CharDesc{
		{"�", "NUL&nbsp;Null"},
		{"�", "SOH&nbsp;Start of Header"},
		{"�", "STX&nbsp;Start of Text"},
		{"�", "ETX&nbsp;End of Text"},
		{"�", "EOT&nbsp;End of Transmission"},
		{"�", "ENQ&nbsp;Enquiry"},
		{"�", "ACK&nbsp;Acknowledge"},
		{"\\a", "BEL&nbsp;Bell"},
		{"\\b", "BS&nbsp;&nbsp;Backspace"},
		{"\\t", "HT&nbsp;&nbsp;Horizontal Tab"},
		{"\\n", "LF&nbsp;&nbsp;Line Feed"},
		{"\\t", "VT&nbsp;&nbsp;Vertical Tab"},
		{"\\f", "FF&nbsp;&nbsp;Form Feed"},
		{"\\r", "CR&nbsp;&nbsp;Carriage Return"},
		{"�", "SO&nbsp;&nbsp;Shift Out"},
		{"�", "SI&nbsp;&nbsp;Shift In"},
		{"�", "DLE&nbsp;Data Link Escape"},
		{"�", "DC1&nbsp;Device Control 1"},
		{"�", "DC2&nbsp;Device Control 2"},
		{"�", "DC3&nbsp;Device Control 3"},
		{"�", "DC4&nbsp;Device Control 4"},
		{"�", "NAK&nbsp;Negative Acknowledge"},
		{"�", "SYN&nbsp;Synchronize"},
		{"�", "ETB&nbsp;End of Transmission Block"},
		{"�", "CAN&nbsp;Cancel"},
		{"�", "EM&nbsp;&nbsp;End of Medium"},
		{"�", "SUB&nbsp;Substitute"},
		{"�", "ESC&nbsp;Escape"},
		{"�", "FS&nbsp;&nbsp;File Separator"},
		{"�", "GS&nbsp;&nbsp;Group Separator"},
		{"�", "RS&nbsp;&nbsp;Record Separator"},
		{"�", "US&nbsp;&nbsp;Unit Separator"},
		{"&nbsp;&nbsp;", "space"}} {
		descForChar[i] = charDesc
	}
	return descForChar
}

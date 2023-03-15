// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"strings"

	"github.com/pwiecz/go-fltk"
)

func makeAsciiTab(x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "&ASCII")
	view := fltk.NewHelpView(x, y, width, height)
	view.SetValue(asciiHtml())
	group.End()
}

func asciiHtml() string {
	var text strings.Builder
	populateAsciiHigh(&text)
	populateAsciiLow(&text)
	return text.String()
}

func populateAsciiHigh(text *strings.Builder) {
	text.WriteString("<p><font face=courier size=4>")
	const start = 33
	const end = 127
	const step = 5
	const stride = (end - start + 1) / step
	for i := 0; i < stride; i++ {
		text.WriteString("&nbsp;")
		j := start + i
		populateOne(text, rune(j+(0*stride)), false)
		populateOne(text, rune(j+(1*stride)), false)
		populateOne(text, rune(j+(2*stride)), false)
		populateOne(text, rune(j+(3*stride)), false)
		populateOne(text, rune(j+(4*stride)), true)
	}
	text.WriteString("</font></p>")
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
	text.WriteString("<p><font face=courier size=4>")
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
	text.WriteString("</font></p>")
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

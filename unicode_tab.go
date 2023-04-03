// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"strconv"
	"strings"
	"unicode"

	"github.com/pwiecz/go-fltk"
)

func makeAsciiTab(app *App, x, y, width, height int) {
	group := fltk.NewFlex(x, y, width, height, "&Unicode")
	group.SetSpacing(pad)
	vbox := fltk.NewFlex(x, y, width, height)
	vbox.SetSpacing(pad)
	hbox := makeTopRow(x, y, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	y += buttonHeight
	hbox = makeChoiceRow(app, x, y, width, buttonHeight)
	vbox.Fixed(hbox, buttonHeight)
	y += buttonHeight
	height -= (2 * buttonHeight)
	app.unicodeView = fltk.NewHelpView(x, y, width, height)
	app.unicodeView.TextFont(fltk.COURIER)
	app.unicodeView.TextSize(app.config.ViewFontSize)
	vbox.End()
	group.End()
	group.Resizable(vbox)
	group.End()
	addCategories(app.categoryChoice, app.unicodeView)
	app.unicodeView.SetValue(getAsciiHigh())
	app.categoryChoice.SetValue(1)
	app.categoryChoice.TakeFocus()
}

func makeTopRow(x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	colWidth := (labelWidth * 3) / 2
	cpLabel := makeAccelLabel(0, 0, colWidth, buttonHeight, "Code &Point")
	cpInput := fltk.NewInput(colWidth, 0, labelWidth, buttonHeight)
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
	hbox.Fixed(cpInput, labelWidth)
	hbox.Fixed(cpCopyButton, labelWidth)
	return hbox
}

func makeChoiceRow(app *App, x, y, width, height int) *fltk.Flex {
	hbox := fltk.NewFlex(x, y, width, height)
	hbox.SetType(fltk.ROW)
	hbox.SetSpacing(pad)
	colWidth := (labelWidth * 3) / 2
	categoryLabel := makeAccelLabel(0, 0, colWidth, buttonHeight,
		"C&ategory")
	app.categoryChoice = fltk.NewChoice(0, 0, width-colWidth, buttonHeight)
	categoryLabel.SetCallback(func() { app.categoryChoice.TakeFocus() })
	hbox.End()
	hbox.Fixed(categoryLabel, colWidth)
	return hbox
}

func addCategories(choice *fltk.Choice, view *fltk.HelpView) {
	choice.Add("ASCII (0-20)", func() {
		view.SetValue(getAsciiLow())
	})
	choice.Add("ASCII (21-7F)", func() {
		view.SetValue(getAsciiHigh())
	})
	choice.Add("NATO Alphabet", func() {
		view.SetValue(getNato())
	})
	choice.Add("Greek", func() {
		view.SetValue(getGreek())
	})
	choice.Add("Symbols", func() {
		runes := getFullRange(5, 0x2013, 0x204A, 0xD7, 0xF7)
		view.SetValue(getSymbols(runes))
	})
	// TODO Unicode categories
	/*
		// 0x00D7-0x00F7
		// 0x2013-0x204A
		// 0x2012-0x2027
		// 0x2030-0x205E
		// 0x20A0-0x20BF
		// 0x2100-0x214F
		// 0x2150-0x218B
		// 0x2190-0x21FF
		// 0x2200-0x22FF
		// 0x2300-0x23FF
		// 0x2460-0x24FF
		// 0x2500-0x257F
		// 0x2580-0x259F
		// 0x25A0-0x25FF
		// 0x2600-0x26FF
		// 0x2700-0x27BF
		// 0xFB00-0xFB06
	*/

}

func cpHtml(s string) (string, fltk.Color) {
	s = strings.TrimSpace(s)
	d, err := strconv.Atoi(s)
	if err != nil {
		d = -1
	}
	if !unicode.IsGraphic(rune(d)) {
		d = -1
	}
	u, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		u = -1
	}
	if !unicode.IsGraphic(rune(u)) {
		u = -1
	}
	if d == -1 && u == -1 {
		return "hex or dec required", fltk.RED
	}
	if d == -1 {
		return fmt.Sprintf("%c  %U  %d", u, u, u), fltk.YELLOW
	} else if u == -1 {
		return fmt.Sprintf("%c  %U  %d", d, d, d), fltk.YELLOW
	} else {
		return fmt.Sprintf("%c  %U | %q %d", d, d, u, u), fltk.YELLOW
	}
}

func getAsciiHigh() string {
	var text strings.Builder
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
			populateOne(&text, rune(j+(k*stride)), k == 4)
		}
	}
	text.WriteString("</p>")
	return text.String()
}

func populateOne(text *strings.Builder, i rune, isEnd bool) {
	if i == 127 {
		text.WriteString("&nbsp;&nbsp;7F backspace<br>")
		return
	}
	var end string
	if isEnd {
		end = "<br>"
	} else {
		end = " <font color=#aaa>|</font> "
	}
	text.WriteString(fmt.Sprintf("<font color=navy>%s</font> %02X%s",
		html.EscapeString(string(i)), i, end))
}

func getAsciiLow() string {
	var text strings.Builder
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
			`&nbsp;<font color=navy>%s</font>&nbsp;&nbsp;%02X&nbsp;<font
			color=navy>%s</font>%s&nbsp;<font color=green>%s</font><br>`,
			c, i, charDesc.shortDesc, istr, charDesc.longDesc))
	}
	text.WriteString("</p>")
	return text.String()
}

type CharDesc struct {
	char      string
	shortDesc string
	longDesc  string
}

type DescForCharMap map[int]CharDesc

func getDescForChar() DescForCharMap {
	descForChar := make(DescForCharMap)
	for i, charDesc := range []CharDesc{
		{"�", "NUL", "Null"},
		{"�", "SOH", "Start of Header"},
		{"�", "STX", "Start of Text"},
		{"�", "ETX", "End of Text"},
		{"�", "EOT", "End of Transmission"},
		{"�", "ENQ", "Enquiry"},
		{"�", "ACK", "Acknowledge"},
		{"\\a", "BEL", "Bell"},
		{"\\b", "BS&nbsp;", "Backspace"},
		{"\\t", "HT&nbsp;", "Horizontal Tab"},
		{"\\n", "LF&nbsp;", "Line Feed"},
		{"\\t", "VT&nbsp;", "Vertical Tab"},
		{"\\f", "FF&nbsp;", "Form Feed"},
		{"\\r", "CR&nbsp;", "Carriage Return"},
		{"�", "SO&nbsp;", "Shift Out"},
		{"�", "SI&nbsp;", "Shift In"},
		{"�", "DLE", "Data Link Escape"},
		{"�", "DC1", "Device Control 1"},
		{"�", "DC2", "Device Control 2"},
		{"�", "DC3", "Device Control 3"},
		{"�", "DC4", "Device Control 4"},
		{"�", "NAK", "Negative Acknowledge"},
		{"�", "SYN", "Synchronize"},
		{"�", "ETB", "End of Transmission Block"},
		{"�", "CAN", "Cancel"},
		{"�", "EM&nbsp;", "End of Medium"},
		{"�", "SUB", "Substitute"},
		{"�", "ESC", "Escape"},
		{"�", "FS&nbsp;", "File Separator"},
		{"�", "GS&nbsp;", "Group Separator"},
		{"�", "RS&nbsp;", "Record Separator"},
		{"�", "US&nbsp;", "Unit Separator"},
		{"&nbsp;&nbsp;", "SPC", "Space"}} {
		descForChar[i] = charDesc
	}
	return descForChar
}

func getGreek() string {
	var text strings.Builder
	text.WriteString("<p>")
	for _, greek := range []struct {
		lower, upper rune
		desc         string
	}{
		{'Α', 'α', "Alpha"},
		{'Β', 'β', "Beta"},
		{'Γ', 'γ', "Gamma"},
		{'Δ', 'δ', "Delta"},
		{'Ε', 'ε', "Epsilon"},
		{'Ζ', 'ζ', "Zeta"},
		{'Η', 'η', "Eta"},
		{'Θ', 'θ', "Theta"},
		{'Ι', 'ι', "Iota"},
		{'Κ', 'κ', "Kappa"},
		{'Λ', 'λ', "Lambda"},
		{'Μ', 'μ', "Mu"},
		{'Ν', 'ν', "Nu"},
		{'Ξ', 'ξ', "Xi"},
		{'Ο', 'ο', "Omicron"},
		{'Π', 'π', "Pi"},
		{'Ρ', 'ρ', "Rho"},
		{'Σ', 'σ', "Σ sigma"},
		{'Τ', 'τ', "Tau"},
		{'Υ', 'υ', "Upsilon"},
		{'Φ', 'φ', "Phi"},
		{'Χ', 'χ', "Chi"},
		{'Ψ', 'ψ', "Psi"},
		{'Ω', 'ω', "Omega"},
	} {
		text.WriteString(fmt.Sprintf(`&nbsp;<font
		color=navy>%c</font>&nbsp;%03X&nbsp;<font
		color=navy>%c</font>&nbsp;%03X&nbsp;%s<br>`, greek.lower,
			greek.lower, greek.upper, greek.upper, greek.desc))
	}
	text.WriteString("</p>")
	return text.String()
}

func getNato() string {
	return "<font color=navy>Alpha Bravo Charlie Delta Echo Foxtrot Golf " +
		"Hotel India Juliet Kilo Lima Mike November Oscar Papa Quebec " +
		"Romeo Sierra Tango Uniform Victor Whiskey X-ray Yankee Zulu</font>"
}

func getSymbols(runes []rune) string {
	var text strings.Builder
	text.WriteString("<p>")
	i := 1
	for _, r := range runes {
		if r == ' ' {
			break
		} else if r >= 0x202A && r <= 0x202F {
			continue
		} else {
			text.WriteString(fmt.Sprintf(
				"&nbsp;<font color=navy>%s</font> %04X",
				html.EscapeString(string(r)), r))
		}
		if i%5 == 0 {
			text.WriteString("<br>")
		} else {
			text.WriteString("&nbsp;<font color=#aaa>|</font>")
		}
		i++
	}
	text.WriteString("</p>")
	return text.String()
}

func getFullRange(step, start, end int, initials ...rune) []rune {
	runes := make([]rune, 0)
	runes = append(runes, initials...)
	for i := start; i <= end; i++ {
		runes = append(runes, rune(i))
	}
	for len(runes)%step != 0 {
		runes = append(runes, ' ')
	}
	return runes
}

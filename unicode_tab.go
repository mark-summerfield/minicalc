// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"strconv"
	"strings"
	"unicode"

	"github.com/mark-summerfield/gset"
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
	addCategories(app)
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

func addCategories(app *App) {
	items := []struct {
		title    string
		callback func()
	}{
		{"AS&CII (00-20)", func() {
			app.unicodeView.SetValue(getAsciiLow())
		}},
		{"ASC&II (21-7E)", func() {
			app.unicodeView.SetValue(getAsciiHigh())
		}},
		{"&NATO Alphabet", func() {
			app.unicodeView.SetValue(getNato())
		}},
		{"&Greek", func() {
			app.unicodeView.SetValue(getGreek())
		}},
		{"&Symbols", func() {
			runes := make([]rune, 0)
			runes = append(runes, 0xD7, 0xF7)
			runes = getFullRange(runes, 0x2012, 0x205E)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"&Arrows", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2190, 0x21FF)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"&Boxes && Blocks", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2500, 0x257F)
			runes = getFullRange(runes, 0x2580, 0x259F)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"&Dingbats", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2700, 0x27BF)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"G&eometric Shapes", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x25A0, 0x25FF)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"&Letterlike", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2100, 0x214F)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"Math &Op.", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2200, 0x22FF)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"&Misc.", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2600, 0x26FF)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"Misc. &Tech.", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2300, 0x23FF)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
		{"Number &Forms", func() {
			runes := make([]rune, 0)
			runes = getFullRange(runes, 0x2150, 0x218B)
			app.unicodeView.SetValue(getSymbols(runes))
		}},
	}
	index := app.config.LastCategory
	for i, item := range items {
		app.categoryChoice.Add(item.title, item.callback)
		if i == index {
			item.callback()
		}
	}
	app.categoryChoice.SetValue(index)
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
	text.WriteString("<table border=1>")
	const (
		start  = 33
		end    = 128
		step   = 8
		stride = (end - start + 1) / step
	)
	for i := 0; i < stride; i++ {
		text.WriteString("<tr>")
		j := start + i
		for k := 0; k < step; k++ {
			r := rune(j + (k * stride))
			if r < 0x7F {
				text.WriteString(fmt.Sprintf(
					"<td><font color=navy>%s</font> %02X</td>",
					html.EscapeString(string(r)), r))
			}
		}
		text.WriteString("</tr>")
	}
	text.WriteString("</table>")
	return text.String()
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
		{"\\b", "BS&nbsp;", "Backspace (also 7F)"},
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
	return "<font color=navy>Alpha<br>Bravo<br>Charlie<br>Delta<br>Echo" +
		"<br>Foxtrot<br>Golf<br>Hotel<br>India<br>Juliet<br>Kilo<br>Lima" +
		"<br>Mike<br>November<br>Oscar<br>Papa<br>Quebec<br>Romeo" +
		"<br>Sierra<br>Tango<br>Uniform<br>Victor<br>Whiskey<br>X-ray" +
		"<br>Yankee<br>Zulu</font>"
}

func getSymbols(runes []rune) string {
	var text strings.Builder
	text.WriteString("<table border=1>")
	const step = 6
	stride := len(runes) / step
	for i := 0; i < stride; i++ {
		text.WriteString("<tr>")
		for k := 0; k < step; k++ {
			index := i + (k * stride)
			if index >= len(runes) {
				break
			}
			r := runes[index]
			text.WriteString(fmt.Sprintf(
				"<td><font color=navy>%s</font> %02X</td>",
				html.EscapeString(string(r)), r))
		}
		text.WriteString("</tr>")
	}
	text.WriteString("</table>")
	return text.String()
}

var badRunes gset.Set[int]

func getFullRange(runes []rune, start, end int) []rune {
	if len(badRunes) == 0 {
		badRunes = gset.New(0x2029, 0x202A, 0x202B, 0x202C, 0x202D, 0x202E,
			0x202F, 0x203F, 0x2040, 0x2041, 0x2043, 0x2044)
	}
	for i := start; i <= end; i++ {
		if badRunes.Contains(i) {
			continue
		}
		runes = append(runes, rune(i))
	}
	return runes
}

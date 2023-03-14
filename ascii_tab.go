// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
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
	text.WriteString("<font face=courier size=5>")
	descForChar := getDescForChar()
	for i := 0; i < len(descForChar); i++ {
		if i == 33 {
			i = 127
		}
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
		text.WriteString(fmt.Sprintf(`%02X&nbsp;%s&nbsp;%s&nbsp;%s<br>`,
			i, c, istr, charDesc.desc))
	}
	text.WriteString("</font>")

	/*
		min := 33
		max := 127
		items := make([]string, max - min + 1)
		step := 5
		stride := len(items) / step
		for i :=0;i<stride;i++{
			item := fmt.Sprintf("%02X",
		}
	*/

	return text.String()
}

type CharDesc struct {
	char string
	desc string
}

func newCharDesc(char, desc string) CharDesc {
	return CharDesc{char, desc}
}

type DescForCharMap map[int]CharDesc

func getDescForChar() DescForCharMap {
	descForChar := make(DescForCharMap)
	descForChar[0] = newCharDesc("�", "NUL&nbsp;(Null)")
	descForChar[1] = newCharDesc("�", "SOH&nbsp;(Start of Header)")
	descForChar[2] = newCharDesc("�", "STX&nbsp;(Start of Text)")
	descForChar[3] = newCharDesc("�", "ETX&nbsp;(End of Text)")
	descForChar[4] = newCharDesc("�", "EOT&nbsp;(End of Transmission)")
	descForChar[5] = newCharDesc("�", "ENQ&nbsp;(Enquiry)")
	descForChar[6] = newCharDesc("�", "ACK&nbsp;(Acknowledge)")
	descForChar[7] = newCharDesc("\\a", "BEL&nbsp;(Bell)")
	descForChar[8] = newCharDesc("\\b", "BS&nbsp;&nbsp;(Backspace)")
	descForChar[9] = newCharDesc("\\t", "HT&nbsp;&nbsp;(Horizontal Tab)")
	descForChar[10] = newCharDesc("\\n", "LF&nbsp;&nbsp;(Line Feed)")
	descForChar[11] = newCharDesc("\\t", "VT&nbsp;&nbsp;(Vertical Tab)")
	descForChar[12] = newCharDesc("\\f", "FF&nbsp;&nbsp;(Form Feed)")
	descForChar[13] = newCharDesc("\\r", "CR&nbsp;&nbsp;(Carriage Return)")
	descForChar[14] = newCharDesc("�", "SO&nbsp;&nbsp;(Shift Out)")
	descForChar[15] = newCharDesc("�", "SI&nbsp;&nbsp;(Shift In)")
	descForChar[16] = newCharDesc("�", "DLE&nbsp;(Data Link Escape)")
	descForChar[17] = newCharDesc("�", "DC1&nbsp;(Device Control 1)")
	descForChar[18] = newCharDesc("�", "DC2&nbsp;(Device Control 2)")
	descForChar[19] = newCharDesc("�", "DC3&nbsp;(Device Control 3)")
	descForChar[20] = newCharDesc("�", "DC4&nbsp;(Device Control 4)")
	descForChar[21] = newCharDesc("�", "NAK&nbsp;(Negative Acknowledge)")
	descForChar[22] = newCharDesc("�", "SYN&nbsp;(Synchronize)")
	descForChar[23] = newCharDesc("�",
		"ETB&nbsp;(End of Transmission Block)")
	descForChar[24] = newCharDesc("�", "CAN&nbsp;(Cancel)")
	descForChar[25] = newCharDesc("�", "EM&nbsp;&nbsp;(End of Medium)")
	descForChar[26] = newCharDesc("�", "SUB&nbsp;(Substitute)")
	descForChar[27] = newCharDesc("�", "ESC&nbsp;(Escape)")
	descForChar[28] = newCharDesc("�", "FS&nbsp;&nbsp;(File Separator)")
	descForChar[29] = newCharDesc("�", "GS&nbsp;&nbsp;(Group Separator)")
	descForChar[30] = newCharDesc("�", "RS&nbsp;&nbsp;(Record Separator)")
	descForChar[31] = newCharDesc("�", "US&nbsp;&nbsp;(Unit Separator)")
	descForChar[32] = newCharDesc("&nbsp;&nbsp;", "space")
	descForChar[127] = newCharDesc("␈", "backspace")
	return descForChar
}

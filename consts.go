// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

const (
	APPNAME       = "MiniCalc"
	BUTTON_HEIGHT = 32

	CALCULATOR_TAB = 0
	REGEX_TAB      = 1
	CPU_RAM_TAB    = 2
	ASCII_TAB      = 3
	CUSTOM_TAB     = 4

	KEY_ESC = 65307

	customPlaceHolderTemplate = `<font color=navy size=4>
To add custom content add</font>
<p><font size=4
face=courier>CustomHtml="""&lt;html&gt;<br>
&lt;body&gt;<br>
...<br>
&lt;/body&gt;<br>
&lt;/html&gt;<br>
"""</font></p>
<p><font color=navy size=4>
to</font> <font size=4 face=courier>%s</font></p>`
)

//go:embed images/icon.svg
var iconSvg string

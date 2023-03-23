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
	LABEL_WIDTH   = 60
	PAD           = 3

	CALCULATOR_TAB = 0
	REGEX_TAB      = 1
	ASCII_TAB      = 2
	CUSTOM_TAB     = 3
	ABOUT_TAB      = 4

	errTemplate               = "<font face=sans color=red size=4>Error: %s</font>"
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

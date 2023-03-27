// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

const (
	appName      = "MiniCalc"
	buttonHeight = 32
	labelWidth   = 60
	pad          = 3

	evaluatorTabIndex = 0
	regexTabIndex     = 1
	asciiTabIndex     = 2
	customTabIndex    = 3
	aboutTabIndex     = 4

	errTemplate               = "<font face=sans color=red size=4>Error: %s</font>"
	customPlaceHolderTemplate = `<font color=navy size=4>
To add custom content add</font>
<p><font size=4
face=courier>CustomHtml=&quot;&quot;&quot;&lt;html&gt;<br>
&lt;body&gt;<br>
...<br>
&lt;/body&gt;<br>
&lt;/html&gt;<br>
&quot;&quot;&quot;</font></p>
<p><font color=navy size=4>
to</font> <font size=4 face=courier>%s</font></p>`

	maxCopyResults = 9
	maxMenuTexts   = 9
	evalShortHelp  = `<p><font color=#888 face=sans size=4>Type an
	expression then press Enter or press F1 for help.</font></p>`
	evalHelpHtml = `<p><font face=sans size=4>Type an expression and press
Enter, e.g., <tt>5 + sqrt(pi)</tt>.</font></p>
<p><font face=sans size=4>Results are automatically assigned to successive
variables, <tt>a</tt>, <tt>b</tt>, ..., unless explicitly assigned with
<tt>=</tt>, e.g., <tt>x = -19 + pow(2, 2/3)</tt></font></p>
<p><font face=sans size=4>To delete a variable use <tt><i>varname</i>=</tt>
and press Enter—or just reassign to it.</font></p>
<p><font face=sans size=4>Supported operators: <tt>+ - * / %</tt>.
</font></p>
<p><font face=sans size=4>Predefined variables: <tt>pi</tt>.
</font></p>
<p><font face=sans size=4>Functions:
<tt>pow(<i>x</i>, <i>y</i>)</tt>,
<tt>rand()</tt>, <tt>randint(<i>x</i>)</tt>,
<tt>sin(<i>n</i>)</tt>,
<tt>sqrt(<i>n</i>)</tt>.
</font></p>
<p><font face=sans size=4>
Click the menu button or press <b>F2</b> to choose a previous expression.
</font></p>
<p><font face=sans size=4>
Click the X close button or press <b>Ctrl+Q</b> to quit.
</font></p>
</font>`
	regexHelpHtml = `<p><font face=sans size=4>Type a regular
expression and some text to test it on and press Enter.<br>
Press F2 to choose a previous regex and Shift+F2 to choose a previous
text</font></p>
<p><font face=sans size=4>
See also the Go regex
<a href="https://pkg.go.dev/regexp/syntax">syntax</a> and
<a href="https://pkg.go.dev/regexp">API</a>.
</font></p>
<p><font face=sans size=4>
Click the X close button or press <b>Ctrl+Q</b> to quit.
</font></p>`
)

//go:embed images/icon.svg
var iconSvg string

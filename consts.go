// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

var themes = []string{"Base", "Gleam", "Gtk", "Oxy", "Plastic"}

const (
	domain       = "qtrac.eu"
	appName      = "MiniCalc"
	buttonHeight = 32
	labelWidth   = 60
	pad          = 5

	evaluatorTabIndex  = 0
	regexTabIndex      = 1
	conversionTabIndex = 2
	accelHintsTabIndex = 3
	unicodeTabIndex    = 4
	customTabIndex     = 5
	optionsTabIndex    = 6
	aboutTabIndex      = 7

	defaultThemeIndex = 3

	defaultAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	escMarker       = "&&"
	placeholder     = "||"
	maxCopyResults  = 9
	maxMenuTexts    = 9

	errTemplate      = "<font color=red size=4>Error: %s</font>"
	defaultRegex     = `\s*(?P<key>\S+)\s*[=:]\s*(?P<value>\S+)`
	defaultRegexText = "scale: 1.15 width=24.5"
	defaultUnhinted  = `Undo
Redo
Copy
Cu&t
Paste
Find
Find Again
Find && Replace`

	customPlaceHolderText = `<p>
<font color=red><i>Default custom text: go to the Options tab to
customize this text (and this tab's title).</i></font></p>
<p>
<b>Aliases</b><br>
<font color=blue>gca</font> — git commit -a<br>
<font color=blue>gco</font> — git checkout <i>branchname</i><br>
<font color=blue>gbc</font> — git checkout -b <i>newbranchname</i><br>
<font color=blue>gbm</font> — git merge <i>frombranchname</i><br>
<font color=blue>gbl</font> — git branch
--list <font color=green><i># list branches</i></font><br>
<font color=blue>gblu</font> — git branch
--no-merged <font color=green><i># unmerged</i></font><br>
<font color=blue>gblm</font> — git branch
--merged <font color=green><i># deletable</i></font><br>
<font color=blue>gst</font> — git status<br>
<font color=blue>git branch -d <i>branchname</i></font>
<font color=green><i># delete branch</i></font><br>
<font color=blue>gor</font> — go run .<br>
<font color=blue>gom</font> — go mod tidy<br>
<font color=blue>gogu</font> — go get -u ./...<br>
</p>
</body>
</html>`
	evalShortHelp = `<p><font color=#888>Type an expression then press
	Enter or press F1 for help.</font></p>`
	evalHelpHtml = `<p>Type an expression and press Enter, e.g.,
	<tt>5 + sqrt(pi)</tt>.</p>
<p>Results are automatically assigned to successive uppercase
variables, <tt>A</tt>, <tt>B</tt>, ..., <tt>Z</tt>, and once <tt>Z</tt> is
reached these variables are reused starting again from <tt>A</tt>. You can
assign to your own variables, e.g., <tt>a = pow(pi, 2)</tt>, <tt>total = A
+ B + x + y</tt>, etc.; these are always preserved and never reused.</p>
<p>To delete a variable use <tt><i>varname</i>=</tt>
and press Enter—or just reassign to it.</p>
<p>Operators: <tt>+ - * / %</tt>.<br>
Predefined variables: <tt>pi</tt>.<br>
Functions:
<tt>len(<i>v</i>)</tt>,
<tt>pow(<i>x</i>, <i>y</i>)</tt>,
<tt>rand()</tt>, <tt>randint(<i>x</i>)</tt>,
<tt>sin(<i>n</i>)</tt>,
<tt>sqrt(<i>n</i>)</tt>, and
<tt>clear()</tt> — delete all automatic variables.
</p>
<p>
For the complete expression syntax see
<a href="https://github.com/maja42/goval#Documentation">goval</a>. Note
that unlike pure goval, MiniCalc supports assignment.
</p>
<p>Click the menu button or press <b>F2</b> to choose a previous expression.
</p>
<p>Click the X close button or press <b>Ctrl+Q</b> to quit.</p>`
	regexHelpHtml = `<p>Type a regular
expression and some text to test it on and press Enter.<br>
Press F2 to choose a previous regex and Shift+F2 to choose a previous
text.</p>
<p>
See also the Go regex
<a href="https://pkg.go.dev/regexp/syntax">syntax</a> and
<a href="https://pkg.go.dev/regexp">API</a>.
</p>
<p>Click the X close button or press <b>Ctrl+Q</b> to quit.</p>`
)

//go:embed images/icon.svg
var iconSvg string

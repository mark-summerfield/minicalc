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
	appName      = "MiniCalc"
	buttonHeight = 32
	labelWidth   = 60
	pad          = 3

	evaluatorTabIndex  = 0
	regexTabIndex      = 1
	asciiTabIndex      = 2
	accelHintsTabIndex = 3
	customTabIndex     = 4
	optionsTabIndex    = 5
	aboutTabIndex      = 6

	defaultThemeIndex = 3

	maxCopyResults = 9
	maxMenuTexts   = 9

	errTemplate           = "<font face=sans color=red size=4>Error: %s</font>"
	customPlaceHolderText = `<p>
<font size=4 face=sans color=red><i>Default custom text: go to the
Options tab to customize this text (and this tab's title).</i></font></p>
<p>
<font size=4 face=sans><b>Aliases</b></font><br>
<font face=courier size=4>
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
</font>
</p>
<p>
<font size=4 face=sans><b>NATO Phonetic Alphabet</b></font><br>
<font size=4 color=navy>Alpha, Bravo, Charlie, Delta, Echo, Foxtrot, 
Golf, Hotel, India, Juliet, Kilo, Lima, Mike, November, Oscar, Papa, 
Quebec, Romeo, Sierra, Tango, Uniform, Victor, Whiskey, X-ray, 
Yankee, Zulu.</font></p>
<p>
<font size=4 face=sans><b>Greek Alphabet</b></font><br>
<font size=4 color=green>Α&nbsp;α&nbsp;alpha, 
Β&nbsp;β&nbsp;beta, Γ&nbsp;γ&nbsp;gamma,  
Δ&nbsp;δ&nbsp;delta, Ε&nbsp;ε&nbsp;epsilon, 
Ζ&nbsp;ζ&nbsp;zeta, Η&nbsp;η&nbsp;eta, 
Θ&nbsp;θ&nbsp;theta, Ι&nbsp;ι&nbsp;iota, 
Κ&nbsp;κ&nbsp;kappa, Λ&nbsp;λ&nbsp;lambda, 
Μ&nbsp;μ&nbsp;mu, Ν&nbsp;ν&nbsp;nu, Ξ&nbsp;ξ&nbsp;xi, 
Ο&nbsp;ο&nbsp;omicron, Π&nbsp;π&nbsp;pi, 
Ρ&nbsp;ρ&nbsp;rho, Σ&nbsp;σ&nbsp;ς&nbsp;sigma, 
Τ&nbsp;τ&nbsp;tau, Υ&nbsp;υ&nbsp;upsilon, 
Φ&nbsp;φ&nbsp;phi, Χ&nbsp;χ&nbsp;chi, Ψ&nbsp;ψ&nbsp;psi, 
Ω&nbsp;ω&nbsp;omega.</font>
</p>
</body>
</html>`
	evalShortHelp = `<p><font color=#888 face=sans size=4>Type an
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

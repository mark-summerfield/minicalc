// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func makeCustomTab(x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "&5 Custom")
	view := fltk.NewHelpView(x, y, width, height)

	// TODO if config.Custom != "" use that else below
	// TODO replace ".ini" with config.Filename
	view.SetValue(fmt.Sprintf(placeHolderTemplate, ".ini"))

	group.End()
}

const placeHolderTemplate = `<font color=navy size=4>
To add custom content add</font>
<p><font size=4
face=courier>custom_html="""&lt;html&gt;<br>
&lt;body&gt;<br>
...<br>
&lt;/body&gt;<br>
&lt;/html&gt;<br>
"""</font></p>
<p><font color=navy size=4>
to</font> <font size=4 face=courier>%s</font></p>`

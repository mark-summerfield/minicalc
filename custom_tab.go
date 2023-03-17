// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeCustomTab(customHtml string, x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "&5 Custom")
	view := fltk.NewHelpView(x, y, width, height)
	view.SetValue(customHtml)
	group.End()
}

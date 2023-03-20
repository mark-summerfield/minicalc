// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeCustomTab(config *Config, x, y, width, height int) *fltk.HelpView {
	group := fltk.NewGroup(x, y, width, height, config.CustomTitle)
	view := fltk.NewHelpView(x, y, width, height)
	view.SetValue(config.CustomHtml)
	group.End()
	view.TakeFocus()
	return view
}

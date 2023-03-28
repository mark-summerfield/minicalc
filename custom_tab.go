// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeCustomTab(app *App, x, y, width, height int) {
	app.customGroup = fltk.NewGroup(x, y, width, height,
		app.config.CustomTitle)
	app.customView = fltk.NewHelpView(x, y, width, height)
	app.customView.SetValue(app.config.CustomHtml)
	app.customGroup.End()
	app.customView.TakeFocus()
}

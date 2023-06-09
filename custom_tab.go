// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeCustomTab(app *App, x, y, width, height int) {
	app.customGroup = fltk.NewFlex(x, y, width, height,
		app.config.CustomTitle)
	app.customGroup.SetSpacing(pad)
	app.customView = fltk.NewHelpView(x, y, width, height)
	app.customView.TextFont(fltk.HELVETICA)
	app.customView.TextSize(app.config.ViewFontSize)
	app.customView.SetValue(app.config.CustomHtml)
	app.customGroup.End()
	app.customView.TakeFocus()
}

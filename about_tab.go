// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeAboutTab(app *App, x, y, width, height int) {
	group := fltk.NewFlex(x, y, width, height, "A&bout")
	group.SetSpacing(pad)
	app.aboutView = fltk.NewHelpView(x, y, width, height)
	app.aboutView.TextFont(fltk.HELVETICA)
	app.aboutView.TextSize(app.config.ViewFontSize)
	app.aboutView.SetValue(aboutHtml())
	group.End()
}

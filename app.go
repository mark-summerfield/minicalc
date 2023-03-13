// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
}

func newApp() *App {
	app := &App{Window: nil}
	app.Window = fltk.NewWindow(512, 480)
	app.Window.SetLabel("MiniCalc")
	addIcons(app.Window, [][]byte{icon16data, icon32data, icon64data})
	addTabs(app)
	app.Window.End()
	return app
}

func addTabs(app *App) {
	width := app.Window.W()
	height := app.Window.H()
	tabs := fltk.NewTabs(0, 0, width, height)
	tabs.SetAlign(fltk.ALIGN_TOP)
	height -= BUTTON_HEIGHT // Allow room for tab
	makeAsciiTab(0, BUTTON_HEIGHT, width, height)
	makeCalculatorTab(0, BUTTON_HEIGHT, width, height)
	makeGreekTab(0, BUTTON_HEIGHT, width, height)
	makeNatoTab(0, BUTTON_HEIGHT, width, height)
	makeOptionsTab(0, BUTTON_HEIGHT, width, height)
	makeUnicodeTab(0, BUTTON_HEIGHT, width, height)
	tabs.End()
}

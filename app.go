// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	calcInput                 *fltk.Input
	calcCopyResultCheckbutton *fltk.CheckButton
	regexInput                *fltk.Input
}

func newApp() *App {
	app := &App{Window: nil}
	app.Window = fltk.NewWindow(512, 480)
	app.Window.Resizable(app.Window)
	app.Window.SetLabel(APPNAME)
	addIcons(app.Window, iconSvg)
	addTabs(app)
	app.Window.End()
	return app
}

func addTabs(app *App) {
	width := app.Window.W()
	height := app.Window.H()
	tabs := fltk.NewTabs(0, 0, width, height)
	tabs.SetAlign(fltk.ALIGN_TOP)
	tabs.SetCallbackCondition(fltk.WhenChanged)
	tabs.SetCallback(func() { onTab(app, tabs) })
	height -= BUTTON_HEIGHT // Allow room for tab
	makeCalculatorTab(app, 0, BUTTON_HEIGHT, width, height)
	makeRegexTab(app, 0, BUTTON_HEIGHT, width, height)
	makeCpuRamTab(0, BUTTON_HEIGHT, width, height)
	makeAsciiTab(0, BUTTON_HEIGHT, width, height)
	makeGreekTab(0, BUTTON_HEIGHT, width, height)
	makeNatoTab(0, BUTTON_HEIGHT, width, height)
	aboutGroup := makeAboutTab(0, BUTTON_HEIGHT, width, height)
	tabs.End()
	tabs.Resizable(aboutGroup)
}

func onTab(app *App, tabs *fltk.Tabs) {
	switch tabs.Value() {
	case CALCULATOR_TAB:
		app.calcInput.TakeFocus()
	case REGEX_TAB:
		app.regexInput.TakeFocus()
		//case CPU_RAM_TAB    :
		//case ASCII_TAB      :
		//case GREEK_TAB      :
		//case NATO_TAB       :
		//case ABOUT_TAB      :
	}
}

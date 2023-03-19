// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	config         *Config
	tabs           *fltk.Tabs
	calcInput      *fltk.Input
	calcResult     float64
	calcPrevResult float64
	regexInput     *fltk.Input
}

func (me *App) onEvent(event fltk.Event) bool {
	if fltk.EventType() == fltk.CLOSE ||
		(fltk.EventType() == fltk.KEY && fltk.EventKey() == fltk.ESCAPE) {
		me.onQuit()
	}
	return false
}

func (me *App) onQuit() {
	me.config.X = me.Window.X()
	me.config.Y = me.Window.Y()
	me.config.Width = me.Window.W()
	me.config.Height = me.Window.H()
	me.config.LastTab = me.tabs.Value()
	me.config.Scale = fltk.ScreenScale(0)
	me.config.save()
}

func newApp(config *Config) *App {
	app := &App{Window: nil, config: config}
	app.Window = fltk.NewWindow(config.Width, config.Height)
	if config.X > -1 && config.Y > -1 {
		app.Window.SetPosition(config.X, config.Y)
	}
	app.Window.Resizable(app.Window)
	app.Window.SetEventHandler(app.onEvent)
	app.Window.SetLabel(APPNAME)
	addIcons(app.Window, iconSvg)
	addTabs(app)
	app.Window.End()
	fltk.AddTimeout(0.1, func() { onTab(app) })
	return app
}

func addTabs(app *App) {
	width := app.Window.W()
	height := app.Window.H()
	app.tabs = fltk.NewTabs(0, 0, width, height)
	app.tabs.SetAlign(fltk.ALIGN_TOP)
	app.tabs.SetCallbackCondition(fltk.WhenChanged)
	app.tabs.SetCallback(func() { onTab(app) })
	height -= BUTTON_HEIGHT // Allow room for tab
	makeCalculatorTab(app, 0, BUTTON_HEIGHT, width, height)
	makeRegexTab(app, 0, BUTTON_HEIGHT, width, height)
	makeCpuRamTab(0, BUTTON_HEIGHT, width, height)
	makeAsciiTab(0, BUTTON_HEIGHT, width, height)
	makeCustomTab(app.config, 0, BUTTON_HEIGHT, width, height)
	aboutGroup := makeAboutTab(app.config.filename, 0, BUTTON_HEIGHT, width,
		height)
	app.tabs.End()
	app.tabs.Resizable(aboutGroup)
	app.tabs.SetValue(app.config.LastTab)
}

func onTab(app *App) {
	switch app.tabs.Value() {
	case CALCULATOR_TAB:
		app.calcInput.TakeFocus()
	case REGEX_TAB:
		app.regexInput.TakeFocus()
	}
}

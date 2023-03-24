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
	evalView       *fltk.HelpView
	evalInput      *fltk.InputChoice
	evalResults    []EvalResult
	evalCopyButton *fltk.MenuButton
	regexView      *fltk.HelpView
	regexInput     *fltk.InputChoice
	regexTextInput *fltk.InputChoice
	asciiView      *fltk.HelpView
	customView     *fltk.HelpView
}

func (me *App) onEvent(event fltk.Event) bool {
	key := fltk.EventKey()
	switch fltk.EventType() {
	case fltk.SHORTCUT:
		if key == fltk.ESCAPE {
			return true
		}
	case fltk.KEY:
		switch key {
		case fltk.HELP, fltk.F1:
			switch me.tabs.Value() {
			case EVALUATOR_TAB:
				me.evalView.SetValue(evalHelpHtml)
				return true
			case REGEX_TAB:
				me.regexView.SetValue(regexHelpHtml)
				return true
			}
		case fltk.F2:
			switch me.tabs.Value() {
			case EVALUATOR_TAB:
				menu := me.evalInput.MenuButton()
				if menu != nil && menu.Size() > 0 {
					menu.Popup()
				}
			case REGEX_TAB:
				var menu *fltk.MenuButton
				if fltk.EventState()&fltk.SHIFT == 0 {
					menu = me.regexInput.MenuButton()
				} else {
					menu = me.regexTextInput.MenuButton()
				}
				if menu != nil && menu.Size() > 0 {
					menu.Popup()
				}
			}
		case 'q', 'Q':
			if fltk.EventState()&fltk.CTRL != 0 {
				me.onQuit()
			}
		}
	case fltk.CLOSE:
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
	me.Window.Destroy()
}

func newApp(config *Config) *App {
	app := &App{Window: nil, config: config,
		evalResults: make([]EvalResult, 0)}
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
	makeEvaluatorTab(app, 0, BUTTON_HEIGHT, width, height)
	makeRegexTab(app, 0, BUTTON_HEIGHT, width, height)
	app.asciiView = makeAsciiTab(0, BUTTON_HEIGHT, width, height)
	app.customView = makeCustomTab(app.config, 0, BUTTON_HEIGHT, width,
		height)
	aboutGroup := makeAboutTab(app.config.filename, 0, BUTTON_HEIGHT, width,
		height)
	app.tabs.End()
	app.tabs.Resizable(aboutGroup)
	app.tabs.SetValue(app.config.LastTab)
}

func onTab(app *App) {
	switch app.tabs.Value() {
	case EVALUATOR_TAB:
		app.evalInput.TakeFocus()
	case REGEX_TAB:
		app.regexInput.TakeFocus()
	case ASCII_TAB:
		app.asciiView.TakeFocus()
	case CUSTOM_TAB:
		app.customView.TakeFocus()
	}
}

// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

type App struct {
	*fltk.Window
	config                      *Config
	tabs                        *fltk.Tabs
	evalView                    *fltk.HelpView
	evalInput                   *fltk.InputChoice
	evalResults                 []EvalResult
	evalCopyButton              *fltk.MenuButton
	regexView                   *fltk.HelpView
	regexInput                  *fltk.InputChoice
	regexTextInput              *fltk.InputChoice
	accelTextEditor             *fltk.TextEditor
	accelTextBuffer             *fltk.TextBuffer
	accelAlphabetInput          *fltk.Input
	accelStatusOutput           *fltk.Output
	accelView                   *fltk.HelpView
	accelShowLettersCheckButton *fltk.CheckButton
	accelShowIndexesCheckButton *fltk.CheckButton
	categoryChoice              *fltk.Choice
	unicodeView                 *fltk.HelpView
	scaleSpinner                *fltk.Spinner
	themeChoice                 *fltk.Choice
	sizeSpinner                 *fltk.Spinner
	showInitialHelpCheckButton  *fltk.CheckButton
	customTitleInput            *fltk.Input
	customTextEditor            *fltk.TextEditor
	customTextBuffer            *fltk.TextBuffer
	customGroup                 *fltk.Flex
	customView                  *fltk.HelpView
	aboutView                   *fltk.HelpView
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
			case evaluatorTabIndex:
				me.evalView.SetValue(evalHelpHtml)
				me.evalView.TakeFocus()
				return true
			case regexTabIndex:
				me.regexView.SetValue(regexHelpHtml)
				return true
			}
		case fltk.F2:
			switch me.tabs.Value() {
			case evaluatorTabIndex:
				menu := me.evalInput.MenuButton()
				if menu != nil && menu.Size() > 0 {
					menu.Popup()
				}
			case regexTabIndex:
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
	// config.Theme is set in callback
	me.config.ShowIntialHelpText = me.showInitialHelpCheckButton.Value()
	me.config.CustomTitle = me.customTitleInput.Value()
	me.config.CustomHtml = me.customTextBuffer.Text()
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
	app.Window.SetLabel(appName)
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
	app.tabs.SetOverflow(fltk.OverflowPulldown)
	app.tabs.SetSelectionColor(fltk.BACKGROUND2_COLOR)
	app.tabs.SetAlign(fltk.ALIGN_TOP)
	app.tabs.SetCallbackCondition(fltk.WhenChanged)
	app.tabs.SetCallback(func() { onTab(app) })
	height -= buttonHeight // Allow room for tab
	makeEvaluatorTab(app, 0, buttonHeight, width, height)
	makeRegexTab(app, 0, buttonHeight, width, height)
	makeAccelHintsTab(app, 0, buttonHeight, width, height)
	makeAsciiTab(app, 0, buttonHeight, width, height)
	makeCustomTab(app, 0, buttonHeight, width, height)
	makeOptionsTab(app, 0, buttonHeight, width, height)
	makeAboutTab(app, 0, buttonHeight, width, height)
	app.tabs.End()
	app.tabs.Resizable(app.Window)
	app.tabs.SetValue(app.config.LastTab)
}

func onTab(app *App) {
	switch app.tabs.Value() {
	case evaluatorTabIndex:
		app.evalInput.TakeFocus()
	case regexTabIndex:
		app.regexInput.TakeFocus()
	case accelHintsTabIndex:
		app.accelTextEditor.TakeFocus()
	case unicodeTabIndex:
		app.categoryChoice.TakeFocus()
	case customTabIndex:
		app.customView.TakeFocus()
	case optionsTabIndex:
		app.scaleSpinner.TakeFocus()
	}
}

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
	evalShowHexCheckButton      *fltk.CheckButton
	evalShowUnicodeCheckButton  *fltk.CheckButton
	evalResults                 []EvalResult
	evalCopyButton              *fltk.MenuButton
	regexView                   *fltk.HelpView
	regexInput                  *fltk.InputChoice
	regexTextInput              *fltk.InputChoice
	convFromChoice              *fltk.Choice
	convToChoice                *fltk.Choice
	convAmountSpinner           *fltk.Spinner
	convResultOutput            *fltk.Output
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
	showTooltipsCheckButton     *fltk.CheckButton
	showInitialHelpCheckButton  *fltk.CheckButton
	customTitleInput            *fltk.Input
	customTextEditor            *fltk.TextEditor
	customTextBuffer            *fltk.TextBuffer
	customTextHighlightBuffer   *fltk.TextBuffer
	customTextStyles            []fltk.StyleTableEntry
	customGroup                 *fltk.Flex
	customView                  *fltk.HelpView
	aboutView                   *fltk.HelpView
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
	app.addTabs()
	app.Window.End()
	fltk.AddTimeout(0.1, func() {
		app.onConfigTooltips()
		app.onTab()
	})
	return app
}

func (me *App) onConfigTooltips() {
	if me.showTooltipsCheckButton.Value() {
		fltk.EnableTooltips()
	} else {
		fltk.DisableTooltips()
	}
}

func (me *App) addTabs() {
	width := me.Window.W()
	height := me.Window.H()
	me.tabs = fltk.NewTabs(0, 0, width, height)
	me.tabs.SetOverflow(fltk.OverflowPulldown)
	me.tabs.SetSelectionColor(fltk.BACKGROUND2_COLOR)
	me.tabs.SetAlign(fltk.ALIGN_TOP)
	me.tabs.SetCallbackCondition(fltk.WhenChanged)
	me.tabs.SetCallback(func() { me.onTab() })
	height -= buttonHeight // Allow room for tab
	makeEvaluatorTab(me, 0, buttonHeight, width, height)
	makeRegexTab(me, 0, buttonHeight, width, height)
	makeConversionTab(me, 0, buttonHeight, width, height)
	makeAccelHintsTab(me, 0, buttonHeight, width, height)
	makeAsciiTab(me, 0, buttonHeight, width, height)
	makeCustomTab(me, 0, buttonHeight, width, height)
	makeOptionsTab(me, 0, buttonHeight, width, height)
	makeAboutTab(me, 0, buttonHeight, width, height)
	me.tabs.End()
	me.tabs.Resizable(me.Window)
	me.tabs.SetValue(me.config.LastTab)
}

func (me *App) onTab() {
	switch me.tabs.Value() {
	case evaluatorTabIndex:
		me.evalInput.TakeFocus()
	case regexTabIndex:
		me.regexInput.TakeFocus()
	case conversionTabIndex:
		me.convFromChoice.TakeFocus()
	case accelHintsTabIndex:
		me.accelTextEditor.TakeFocus()
	case unicodeTabIndex:
		me.categoryChoice.TakeFocus()
	case customTabIndex:
		me.customView.TakeFocus()
	case optionsTabIndex:
		me.scaleSpinner.TakeFocus()
	}
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
	me.config.LastCategory = me.categoryChoice.Value()
	me.config.LastRegex = me.regexInput.Value()
	me.config.LastRegexText = me.regexTextInput.Value()
	me.config.LastUnhinted = me.accelTextBuffer.Text()
	me.config.LastFromIndex = me.convFromChoice.Value()
	me.config.LastToIndex = me.convToChoice.Value()
	me.config.LastAmount = me.convAmountSpinner.Value()
	me.config.Scale = fltk.ScreenScale(0)
	// config.Theme is set in callback
	me.config.ShowTooltips = me.showTooltipsCheckButton.Value()
	me.config.ShowInitialHelpText = me.showInitialHelpCheckButton.Value()
	me.config.CustomTitle = me.customTitleInput.Value()
	me.config.CustomHtml = me.customTextBuffer.Text()
	me.config.save()
	me.Window.Destroy()
}

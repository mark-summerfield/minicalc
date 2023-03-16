// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

const (
	APPNAME       = "MiniCalc"
	BUTTON_HEIGHT = 32

	CALCULATOR_TAB = 0
	REGEX_TAB      = 1
	CPU_RAM_TAB    = 2
	ASCII_TAB      = 3
	GREEK_TAB      = 4
	NATO_TAB       = 5
	ABOUT_TAB      = 6
)

//go:embed images/icon.svg
var iconSvg string

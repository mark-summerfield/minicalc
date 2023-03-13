// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	fltk.SetScheme("oxy")
	// TODO save/restore window size/pos
	app := newApp()
	app.Show()
	fltk.Run()
}

// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	fltk.SetScheme("oxy")
	fltk.SetScreenScale(0, 1.1)
	// TODO save/restore window size/pos & scale & last tab
	app := newApp()
	app.Show()
	fltk.Run()
}

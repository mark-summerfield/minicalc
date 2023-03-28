// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	config := newConfig()
	fltk.SetScheme(config.Theme)
	fltk.SetScreenScale(0, config.Scale)
	app := newApp(config)
	app.Show()
	fltk.Run()
}

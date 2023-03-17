// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func makeCustomTab(config *Config, x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height,
		fmt.Sprintf("&5 %s", config.CustomTitle))
	view := fltk.NewHelpView(x, y, width, height)
	view.SetValue(config.CustomHtml)
	group.End()
}

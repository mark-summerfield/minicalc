// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func makeAsciiTab(x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "&ASCII")
	// TODO
	group.End()
}

// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	window := makeWindow()
	window.Show()
	fltk.Run()
}

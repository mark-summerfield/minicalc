// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "image/png"

	"github.com/pwiecz/go-fltk"
)

func addIcons(window *fltk.Window, svgText string) {
	if svg, err := fltk.NewSvgImageFromString(svgText); err == nil {
		icon := fltk.NewRgbImageFromSvg(svg)
		window.SetIcons([]*fltk.RgbImage{icon})
	}
}

func makeAccelLabel(x, y, w, h int, label string) *fltk.Button {
	button := fltk.NewButton(x, y, w, h, label)
	button.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	button.SetBox(fltk.NO_BOX)
	button.ClearVisibleFocus()
	return button
}

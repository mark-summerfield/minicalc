// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"bytes"
	"image"
	_ "image/png"
	"runtime"

	"github.com/pwiecz/go-fltk"
)

func makeWindow() *fltk.Window {
	// TODO save/restore size/pos
	window := fltk.NewWindow(512, 480)
	window.SetLabel("MiniCalc")
	icons := makeIcons([][]byte{icon16data, icon32data, icon64data})
	if len(icons) > 0 {
		window.SetIcons(icons)
	}
	// TODO
	window.End()
	return window
}

// NOTE until we get SVG icon support
func makeIcons(iconData [][]byte) []*fltk.RgbImage {
	runtime.LockOSThread()
	icons := make([]*fltk.RgbImage, 0, 3)
	for _, datum := range iconData {
		if img, _, err := image.Decode(bytes.NewReader(datum)); err == nil {
			if icon, err := fltk.NewRgbImageFromImage(img); err == nil {
				icons = append(icons, icon)
			}
		}
	}
	return icons
}

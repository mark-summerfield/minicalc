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

// NOTE only until we get SVG icon support
func addIcons(window *fltk.Window, iconData [][]byte) {
	runtime.LockOSThread()
	icons := make([]*fltk.RgbImage, 0, 3)
	for _, datum := range iconData {
		if img, _, err := image.Decode(bytes.NewReader(datum)); err == nil {
			if icon, err := fltk.NewRgbImageFromImage(img); err == nil {
				icons = append(icons, icon)
			}
		}
	}
	if len(icons) > 0 {
		window.SetIcons(icons)
	}
}

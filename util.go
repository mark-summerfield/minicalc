// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"strings"

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

func updateInputChoice(choice *fltk.InputChoice) {
	current := choice.Value()
	menu := choice.MenuButton()
	texts := []string{current}
	for i := 0; i < menu.Size(); i++ {
		text := menu.Text(i)
		if !(current == text || strings.Contains(current, text) ||
			strings.Contains(text, current)) {
			texts = append(texts, text)
		}
	}
	for i := menu.Size() - 1; i >= 0; i-- {
		menu.Remove(i)
	}
	for _, text := range texts {
		text := text
		escText := strings.ReplaceAll(
			strings.ReplaceAll(text, "\\", "\\\\"), "/", `\/`)
		menu.Add(escText, func() { choice.Input().SetValue(text) })
	}
}

func makeTextEditor(x, y, width, height int) (*fltk.TextEditor,
	*fltk.TextBuffer) {
	buffer := fltk.NewTextBuffer()
	editor := fltk.NewTextEditor(x, y, width, height)
	editor.SetBuffer(buffer)
	return editor, buffer
}

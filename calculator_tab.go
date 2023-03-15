// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"math"
	"strings"

	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/minicalc/eval"
	"github.com/pwiecz/go-fltk"
)

func makeCalculatorTab(x, y, width, height int) {
	allPrevious := make([]string, 0)
	nextVarName := "a"
	calcEnv := eval.Env{"pi": math.Pi}
	group := fltk.NewGroup(x, y, width, height, "&Calculator")
	vbox := fltk.NewPack(x, y, width, height)
	hoffset := 2 * BUTTON_HEIGHT
	calcView := fltk.NewHelpView(x, y, width, height-hoffset)
	calcInput := fltk.NewInput(x, y+height-hoffset, width, BUTTON_HEIGHT)
	autoCopy := fltk.NewCheckButton(x, y+height-BUTTON_HEIGHT, width,
		BUTTON_HEIGHT, "A&uto Copy Result to Clipboard")
	autoCopy.SetValue(true)
	calcInput.SetCallbackCondition(fltk.WhenEnterKey)
	calcInput.SetCallback(func() {
		const maxPrevious = 5
		const errTemplate = "<font color=red>Error: %s</font>"
		var text strings.Builder
		text.WriteString("<font size=4>")
		limit := 0
		if len(allPrevious) > maxPrevious {
			limit = len(allPrevious) - maxPrevious
		}
		for _, previous := range allPrevious[limit:] {
			text.WriteString(previous)
		}
		keys := gong.SortedMapKeys(calcEnv)
		for _, key := range keys {
			text.WriteString(fmt.Sprintf(
				"<font color=blue>%s = %g</font><br>", key, calcEnv[key]))
		}
		expression := calcInput.Value()
		expr, err := eval.Parse(expression)
		if err != nil {
			text.WriteString(fmt.Sprintf(errTemplate, html.EscapeString(
				err.Error())))
		} else {
			err := expr.Check(map[eval.Var]bool{})
			if err != nil {
				text.WriteString(fmt.Sprintf(errTemplate, html.EscapeString(
					err.Error())))
			} else {
				value := expr.Eval(calcEnv)
				text.WriteString(fmt.Sprintf(
					"<font color=green>%s = %s → <b>%g</b></font>",
					nextVarName,
					expression, value))
				calcEnv[eval.Var(nextVarName)] = value
				nextVarName = getNextVarName(nextVarName)
				allPrevious = append(allPrevious, fmt.Sprintf("%s → %g<br>",
					expression, value))
				if autoCopy.Value() {
					fltk.CopyToClipboard(fmt.Sprintf("%g", value))
				}
			}
		}
		text.WriteString("</font>")
		calcView.SetValue(text.String())
	})
	vbox.End()
	vbox.Resizable(calcView) // TODO Doesn't work: need Flex
	group.End()
	group.Resizable(vbox)
	// TODO calcInput.TakeFocus()
}

func getNextVarName(name string) string {
	first := rune(name[0])
	if len(name) == 1 {
		if first < 'z' {
			return string(first + 1)
		}
		return "aa"
	}
	second := rune(name[1])
	if second == 'z' {
		first++
	} else {
		second++
	}
	return string(first) + string(second)
}

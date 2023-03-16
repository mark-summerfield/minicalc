// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"math"
	"regexp"
	"strings"

	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/minicalc/eval"
	"github.com/pwiecz/go-fltk"
)

func makeCalculatorTab(app *App, x, y, width, height int) {
	allPrevious := make([]string, 0)
	nextVarName := "a"
	calcEnv := eval.Env{"pi": math.Pi}
	group := fltk.NewGroup(x, y, width, height, "&1 Calculator")
	vbox := fltk.NewPack(x, y, width, height)
	hoffset := 2 * BUTTON_HEIGHT
	calcView := fltk.NewHelpView(x, y, width, height-hoffset)
	calcView.SetValue(CALC_HELP_HTML)
	app.calcInput = fltk.NewInput(x, y+height-hoffset, width,
		BUTTON_HEIGHT)
	app.calcCopyResultCheckbutton = fltk.NewCheckButton(x,
		y+height-BUTTON_HEIGHT, width, BUTTON_HEIGHT,
		"&Copy Each Result to the Clipboard")
	app.calcCopyResultCheckbutton.SetValue(true)
	app.calcInput.SetCallbackCondition(fltk.WhenEnterKey)
	app.calcInput.SetCallback(func() {
		allPrevious, nextVarName = onCalc(allPrevious, calcEnv, calcView,
			app.calcInput, app.calcCopyResultCheckbutton, nextVarName)
	})
	vbox.End()
	vbox.Resizable(calcView) // TODO Doesn't work: need Flex
	group.End()
	group.Resizable(vbox)
	app.calcInput.TakeFocus()
}

func onCalc(allPrevious []string, calcEnv eval.Env, calcView *fltk.HelpView,
	calcInput *fltk.Input, calcCopyResultCheckbutton *fltk.CheckButton,
	nextVarName string) ([]string, string) {
	const errTemplate = "<font color=red>Error: %s</font>"
	autoVar := true
	deletion := false
	var text string
	var err error
	varName, expression, err := getVarNameAndExpression(calcInput.Value())
	if err != nil {
		text = fmt.Sprintf(errTemplate, html.EscapeString(err.Error()))
	} else if varName != "" {
		autoVar = false
		if expression == "" {
			deletion = true
			delete(calcEnv, eval.Var(varName))
			text = fmt.Sprintf(
				"<font color=purple>deleted <b>%s</b></font>", varName)
		}
	}
	if err == nil && !deletion {
		expr, err := eval.Parse(expression)
		if err != nil {
			text = fmt.Sprintf(errTemplate, html.EscapeString(err.Error()))
		} else {
			err := expr.Check(map[eval.Var]bool{})
			if err != nil {
				text = fmt.Sprintf(errTemplate, html.EscapeString(
					err.Error()))
			} else {
				value := expr.Eval(calcEnv)
				if autoVar {
					varName = nextVarName
					calcEnv[eval.Var(varName)] = value
					nextVarName = getNextVarName(calcEnv)
				} else {
					calcEnv[eval.Var(varName)] = value
				}
				text = fmt.Sprintf(
					"<font color=green>%s = %s → <b>%g</b></font>",
					varName, expression, value)
				allPrevious = append(allPrevious, fmt.Sprintf("%s → %g<br>",
					expression, value))
				if calcCopyResultCheckbutton.Value() {
					fltk.CopyToClipboard(fmt.Sprintf("%g", value))
				}
			}
		}
	}
	var textBuilder strings.Builder
	textBuilder.WriteString("<font size=4>")
	keys := gong.SortedMapKeys(calcEnv)
	for _, key := range keys {
		if string(key) != varName {
			textBuilder.WriteString(fmt.Sprintf(
				"<font color=blue>%s = %g</font><br>", key, calcEnv[key]))
		}
	}
	textBuilder.WriteString(text)
	textBuilder.WriteString("</font>")
	calcView.SetValue(textBuilder.String())
	return allPrevious, nextVarName
}

func getVarNameAndExpression(expression string) (string, string, error) {
	identifierRx := regexp.MustCompile(`\pL[\pL\d_]*`)
	parts := strings.SplitN(expression, "=", 2)
	if len(parts) == 1 {
		return "", strings.TrimSpace(expression), nil
	}
	varName := strings.TrimSpace(parts[0])
	expression = strings.TrimSpace(parts[1])
	if identifierRx.MatchString(varName) {
		return varName, expression, nil
	}
	return "", "", fmt.Errorf("%q is not a valid identifier", varName)
}

func getNextVarName(calcEnv eval.Env) string {
	for i := 'a'; i <= 'z'; i++ {
		varName := string(i)
		if _, found := calcEnv[eval.Var(varName)]; !found {
			return varName
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			varName := string(i) + string(j)
			if _, found := calcEnv[eval.Var(varName)]; !found {
				return varName
			}
		}
	}
	panic("more than 700 variables!")
}

const CALC_HELP_HTML = `<p><font size=4>Type an expression and press
Enter.</font></p>
<p><font size=4>Results are automatically assigned to successive variables,
<tt>a</tt>, <tt>b</tt>, ... unless explicitly assigned with <tt>=</tt>.
</font></p>
<p><font size=4>To delete a variable use <tt><i>varname</i>=</tt> and press
Enter.</font></p>
<p><font size=4>Supported operators: <tt>+ - * / %</tt>.
</font></p>
<p><font size=4>Predefined variables: <tt>pi</tt>.
</font></p>
<p><font size=4>Functions:
<tt>pow(<i>x</i>, <i>y</i>)</tt>,
<tt>sin(<i>n</i>)</tt>,
<tt>sqrt(<i>n</i>)</tt>.
</font></p>
</font>`

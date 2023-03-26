// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"html"
	"math"
	"regexp"
	"strings"
	"unicode"

	"github.com/mark-summerfield/accelhint"
	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/gset"
	"github.com/mark-summerfield/minicalc/eval"
	"github.com/pwiecz/go-fltk"
)

func makeEvaluatorTab(app *App, x, y, width, height int) {
	nextVarName := "a"
	evalEnv := eval.Env{"pi": math.Pi}
	group := fltk.NewFlex(x, y, width, height, "E&valuator")
	vbox := fltk.NewFlex(x, y, width, height)
	app.evalView = fltk.NewHelpView(x, y, width, height-buttonHeight)
	if app.config.ShowIntialHelpText {
		app.evalView.SetValue(evalHelpHtml)
	} else {
		app.evalView.SetValue(evalShortHelp)
	}
	hbox := makeBottomRow(app, x, y, width, height, nextVarName, evalEnv)
	vbox.End()
	vbox.Fixed(hbox, buttonHeight)
	group.End()
	group.Resizable(vbox)
	app.evalInput.TakeFocus()
}

func makeBottomRow(app *App, x, y, width, height int,
	nextVarName string, evalEnv eval.Env) *fltk.Flex {
	const BUTTON_WIDTH = labelWidth + (2 * pad)
	userVarNames := gset.New[string]()
	hbox := fltk.NewFlex(x, y+height-buttonHeight, width, buttonHeight)
	hbox.SetType(fltk.ROW)
	app.evalInput = fltk.NewInputChoice(x, y+height-buttonHeight,
		width-BUTTON_WIDTH, buttonHeight)
	app.evalCopyButton = fltk.NewMenuButton(x, y+height-buttonHeight,
		BUTTON_WIDTH, buttonHeight, "&Copy")
	app.evalCopyButton.ClearVisibleFocus()
	app.evalCopyButton.Deactivate()
	app.evalInput.Input().SetCallbackCondition(fltk.WhenEnterKeyAlways)
	i := new(int)
	app.evalInput.Input().SetCallback(func() {
		*i++
		text := fmt.Sprintf("Enter #%d", *i)
		fmt.Println(text)
		updateInputChoice(app.evalInput)
		nextVarName = onEval(app, userVarNames, evalEnv, nextVarName)
		if strings.TrimSpace(app.evalInput.Value()) == "" {
			app.evalView.SetValue(fmt.Sprintf(
				"<font face=sans size=4 color=red><b>%s</b></font>", text))
		}
	})
	hbox.End()
	hbox.Fixed(app.evalCopyButton, BUTTON_WIDTH)
	return hbox
}

func onEval(app *App, userVarNames gset.Set[string], evalEnv eval.Env,
	nextVarName string) string {
	input := strings.TrimSpace(app.evalInput.Value())
	autoVar := true
	deletion := false
	var text string
	var err error
	varName, expression, err := getVarNameAndExpression(userVarNames, input)
	if err != nil {
		text = fmt.Sprintf(errTemplate, html.EscapeString(err.Error()))
	} else if varName != "" {
		autoVar = false
		if expression == "" { // varName=
			deletion = true
			delete(userVarNames, varName)
			delete(evalEnv, eval.Var(varName))
			text = fmt.Sprintf(
				"<font face=sans color=purple>deleted <b>%s</b></font>",
				varName)
		}
	}
	if err == nil && !deletion { // varName=expr _or_ expr
		text, varName, nextVarName = evaluate(app, varName, nextVarName,
			expression, autoVar, evalEnv, userVarNames)
	}
	populateView(varName, text, evalEnv, app.evalView)
	return nextVarName
}

func getVarNameAndExpression(userVarNames gset.Set[string],
	expression string) (string, string, error) {
	identifierRx := regexp.MustCompile(`^\pL[\pL\d_]*$`)
	parts := strings.SplitN(expression, "=", 2)
	if len(parts) == 1 {
		return "", strings.TrimSpace(expression), nil
	}
	varName := strings.TrimSpace(parts[0])
	expression = strings.TrimSpace(parts[1])
	if identifierRx.MatchString(varName) {
		userVarNames.Add(varName)
		return varName, expression, nil
	}
	return "", "", fmt.Errorf("%q is not a valid identifier", varName)
}

func evaluate(app *App, varName, nextVarName, expression string,
	autoVar bool, evalEnv eval.Env, userVarNames gset.Set[string]) (string,
	string, string) {
	var text string
	expr, err := eval.Parse(expression)
	if err != nil {
		text = fmt.Sprintf(errTemplate, html.EscapeString(err.Error()))
	} else {
		err := expr.Check(map[eval.Var]bool{})
		if err != nil {
			text = fmt.Sprintf(errTemplate, html.EscapeString(
				err.Error()))
		} else {
			value := expr.Eval(evalEnv)
			if autoVar {
				nextVarName = getNextVarName(evalEnv, userVarNames)
				varName = nextVarName
			}
			evalEnv[eval.Var(varName)] = value
			app.evalResults = append(app.evalResults,
				EvalResult{varName, value})
			updateEvalCopyButton(app)
			text = fmt.Sprintf(
				`<font face=sans color=green>%s = %s → </font><font
				face=sans color=blue><b>%g</b>%s</font>`, varName,
				expression, value, getResultDetails(value))
		}
	}
	return text, varName, nextVarName
}

func getResultDetails(value float64) string {
	var text string
	if value > 0 && math.Trunc(value) == value {
		v := int64(value)
		text += fmt.Sprintf(" • 0x%X", v)
		if v > 32 && v <= unicode.MaxRune {
			text += fmt.Sprintf(" • '%c'", v)
		}
	}
	return text
}

func getNextVarName(evalEnv eval.Env,
	userVarNames gset.Set[string]) string {
	for i := 'a'; i <= 'z'; i++ {
		varName := string(i)
		if userVarNames.Contains(varName) {
			continue
		}
		if _, found := evalEnv[eval.Var(varName)]; !found {
			return varName
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		for j := 'a'; j <= 'z'; j++ {
			varName := string(i) + string(j)
			if userVarNames.Contains(varName) {
				continue
			}
			if _, found := evalEnv[eval.Var(varName)]; !found {
				return varName
			}
		}
	}
	panic("can't cope with more than 700 variables")
}

func populateView(varName, text string, evalEnv eval.Env,
	evalView *fltk.HelpView) {
	var textBuilder strings.Builder
	textBuilder.WriteString("<font face=sans size=4>")
	keys := gong.SortedMapKeys(evalEnv)
	for _, key := range keys {
		bs, be := "", ""
		if string(key) == varName {
			bs, be = "<b>", "</b>"
		}
		value := evalEnv[key]
		textBuilder.WriteString(fmt.Sprintf(
			`<font face=sans color=green>%s%s%s = <font face=sans
			color=blue>%g</font><font face=sans color=#444>%s</font><br>`,
			bs, key, be, value, getResultDetails(value)))
	}
	textBuilder.WriteString(text)
	textBuilder.WriteString("</font>")
	evalView.SetValue(textBuilder.String())
	// Scroll to end
	evalView.SetTopLine(999999)
	evalView.SetTopLine(evalView.TopLine() - evalView.H())
}

func updateEvalCopyButton(app *App) {
	seen := gset.New[float64]()
	filtered := make([]EvalResult, 0, len(app.evalResults))
	for _, evalResult := range app.evalResults {
		if !seen.Contains(evalResult.value) {
			seen.Add(evalResult.value)
			filtered = append(filtered, evalResult)
		}
	}
	app.evalResults = filtered
	if len(app.evalResults) > maxCopyResults {
		app.evalResults = app.evalResults[len(app.evalResults)-
			maxCopyResults:]
	}
	for i := app.evalCopyButton.Size() - 1; i >= 0; i-- {
		app.evalCopyButton.Remove(i)
	}
	varNames := make([]string, 0, len(app.evalResults))
	for _, evalResult := range app.evalResults {
		varNames = append(varNames, evalResult.varName)
	}
	hinted, _, err := accelhint.Hinted(varNames)
	for i, evalResult := range app.evalResults {
		varName := evalResult.varName
		if err == nil {
			varName = hinted[i]
		}
		value := evalResult.value
		app.evalCopyButton.AddEx(fmt.Sprintf(
			"%s = %g", varName, value), 0,
			func() { fltk.CopyToClipboard(fmt.Sprintf("%g", value)) }, 0)
	}
	if app.evalCopyButton.Size() > 0 {
		app.evalCopyButton.Activate()
	} else {
		app.evalCopyButton.Deactivate()
	}
}

type EvalResult struct {
	varName string
	value   float64
}

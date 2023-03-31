// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"errors"
	"fmt"
	"html"
	"math"
	"math/rand"
	"regexp"
	"strings"

	"github.com/maja42/goval"
	"github.com/mark-summerfield/accelhint"
	"github.com/mark-summerfield/gong"
	"github.com/mark-summerfield/gset"
	"github.com/pwiecz/go-fltk"
)

type VarMap map[string]any

func makeEvaluatorTab(app *App, x, y, width, height int) {
	evalEnv := newEvalEnv()
	group := fltk.NewFlex(x, y, width, height, "E&valuator")
	vbox := fltk.NewFlex(x, y, width, height)
	app.evalView = fltk.NewHelpView(x, y, width, height-buttonHeight)
	app.evalView.TextFont(fltk.COURIER)
	app.evalView.TextSize(app.config.ViewFontSize)
	if app.config.ShowIntialHelpText {
		app.evalView.SetValue(evalHelpHtml)
	} else {
		app.evalView.SetValue(evalShortHelp)
	}
	hbox := makeBottomRow(app, x, y, width, height, evalEnv)
	vbox.End()
	vbox.Fixed(hbox, buttonHeight)
	group.End()
	group.Resizable(vbox)
	app.evalInput.TakeFocus()
}

func makeBottomRow(app *App, x, y, width, height int,
	evalEnv *EvalEnv) *fltk.Flex {
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
	app.evalInput.Input().SetCallback(func() {
		updateInputChoice(app.evalInput)
		onEval(app, userVarNames, evalEnv)
	})
	hbox.End()
	hbox.Fixed(app.evalCopyButton, BUTTON_WIDTH)
	return hbox
}

func onEval(app *App, userVarNames gset.Set[string], evalEnv *EvalEnv) {
	input := strings.TrimSpace(app.evalInput.Value())
	if input == "clear()" {
		clear(app.evalView, evalEnv.variables)
	} else {
		autoVar := true
		deletion := false
		var text string
		var varName string
		var err error
		varName, expression, err := getVarNameAndExpression(userVarNames,
			input)
		if err != nil {
			text = fmt.Sprintf(errTemplate, html.EscapeString(err.Error()))
		} else if varName != "" {
			autoVar = false
			if expression == "" { // varName=
				deletion = true
				delete(userVarNames, varName)
				delete(evalEnv.variables, varName)
				text = fmt.Sprintf(
					"<font color=purple>deleted <b>%s</b></font>", varName)
			}
		}
		if err == nil && !deletion { // varName=expr _or_ expr
			text, varName = evaluate(app, varName, expression, autoVar,
				evalEnv, userVarNames)
		}
		populateView(varName, text, evalEnv.variables, app.evalView)
	}
}

func clear(evalView *fltk.HelpView, variables VarMap) {
	n := 0
	for name := range variables {
		if len(name) == 1 && name[0] >= 'A' && name[0] <= 'Z' {
			delete(variables, name)
			n++
		}
	}
	text := fmt.Sprintf(
		"<font color=purple>deleted %s automatic variables</font>",
		gong.Commas(n))
	populateView("", text, variables, evalView)
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

func evaluate(app *App, varName, expression string, autoVar bool,
	evalEnv *EvalEnv, userVarNames gset.Set[string]) (string, string) {
	var text string
	evaluator := goval.NewEvaluator()
	value, err := evaluator.Evaluate(expression, evalEnv.variables,
		evalEnv.functions)
	if err != nil {
		text = fmt.Sprintf(errTemplate, html.EscapeString(err.Error()))
	} else {
		if autoVar {
			getNextVarName(evalEnv)
			varName = evalEnv.nextVarName
		}
		evalEnv.variables[varName] = value
		app.evalResults = append(app.evalResults,
			EvalResult{varName, value})
		updateEvalCopyButton(app)
		text = fmt.Sprintf(`<font color=green>%s = %s → </font><font
				color=blue><b>%v</b></font>`, varName, expression, value)
	}
	return text, varName
}

func getNextVarName(evalEnv *EvalEnv) {
	for i := 'A'; i <= 'Z'; i++ {
		varName := string(i)
		if _, found := evalEnv.variables[varName]; !found {
			evalEnv.nextVarName = varName
			return
		}
	}
	// All variables have been used, so start to reuse.
	evalEnv.nextVarName = string(evalEnv.nextReuseRune)
	evalEnv.nextReuseRune++
	if evalEnv.nextReuseRune > 'Z' {
		evalEnv.nextReuseRune = 'A'
	}
}

func populateView(varName, text string, variables VarMap,
	evalView *fltk.HelpView) {
	var textBuilder strings.Builder
	keys := gong.SortedMapKeys(variables)
	for _, key := range keys {
		bs, be := "", ""
		if string(key) == varName {
			bs, be = "<b>", "</b>"
		}
		value := variables[key]
		textBuilder.WriteString(fmt.Sprintf(
			`<font color=green>%s%s%s = <font color=blue>%v</font><br>`,
			bs, key, be, value))
	}
	textBuilder.WriteString(text)
	evalView.SetValue(textBuilder.String())
	// Scroll to end
	evalView.SetTopLine(999999)
	evalView.SetTopLine(evalView.TopLine() - evalView.H())
}

func updateEvalCopyButton(app *App) {
	seen := gset.New[string]()
	filtered := make([]EvalResult, 0, len(app.evalResults))
	for _, evalResult := range app.evalResults {
		value := fmt.Sprintf("%v", evalResult.value)
		if !seen.Contains(value) {
			seen.Add(value)
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
		value := fmt.Sprintf("%v", evalResult.value)
		app.evalCopyButton.Add(fmt.Sprintf("%s = %s", varName, value),
			func() { fltk.CopyToClipboard(value) })
	}
	if app.evalCopyButton.Size() > 0 {
		app.evalCopyButton.Activate()
	} else {
		app.evalCopyButton.Deactivate()
	}
}

type EvalResult struct {
	varName string
	value   any
}

type EvalEnv struct {
	nextVarName   string
	nextReuseRune rune
	variables     VarMap
	functions     map[string]goval.ExpressionFunction
}

func newEvalEnv() *EvalEnv {
	variables := make(VarMap)
	variables["pi"] = math.Pi
	functions := make(map[string]goval.ExpressionFunction)
	functions["len"] = func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("len(v): needs one argument")
		}
		s, ok := args[0].(string)
		if ok {
			return len(s), nil
		}
		array, ok := args[0].([]any)
		if ok {
			return len(array), nil
		}
		obj, ok := args[0].(map[string]any)
		if ok {
			return len(obj), nil
		}
		return nil, fmt.Errorf("len(v): needs a string, array or object")
	}
	functions["pow"] = func(args ...any) (any, error) {
		if len(args) != 2 {
			return nil, errors.New("pow(x, y): needs two arguments")
		}
		x, err := getReal(args[0])
		if err != nil {
			return nil, fmt.Errorf("pow(x, y):x: %w", err)
		}
		y, err := getReal(args[1])
		if err != nil {
			return nil, fmt.Errorf("pow(x, y):y: %w", err)
		}
		return math.Pow(x, y), nil
	}
	functions["rand"] = func(...any) (any, error) {
		return rand.Float64(), nil
	}
	functions["randint"] = func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, errors.New("randint(n): needs one argument")
		}
		n, err := getInt(args[0])
		if err != nil {
			return nil, fmt.Errorf("randint(n) %w", err)
		}
		return rand.Intn(n), nil
	}
	functions["sin"] = func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, errors.New("sin(n): needs one argument")
		}
		x, err := getReal(args[0])
		if err != nil {
			return nil, fmt.Errorf("sin(x): %w", err)
		}
		return math.Sin(x), nil
	}
	functions["sqrt"] = func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, errors.New("sqrt(n): needs one argument")
		}
		x, err := getReal(args[0])
		if err != nil {
			return nil, fmt.Errorf("sqrt(x): %w", err)
		}
		if x < 0.0 {
			return nil, errors.New("sqrt(x): x must be positive")
		}
		return math.Sqrt(x), nil
	}
	nextVarName := "A"
	nextReuseRune := 'A'
	return &EvalEnv{nextVarName, nextReuseRune, variables, functions}
}

func getReal(x any) (float64, error) {
	switch v := x.(type) {
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	}
	return 0.0, fmt.Errorf("real expected, got %v", x)
}

func getInt(x any) (int, error) {
	switch v := x.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	}
	return 0.0, fmt.Errorf("int expected, got %v", x)
}

// Modifications Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

type Env map[Var]float64

func (me Var) Eval(env Env) float64 {
	return env[me]
}

func (me literal) Eval(_ Env) float64 {
	return float64(me)
}

func (me unary) Eval(env Env) float64 {
	switch me.op {
	case '+':
		return +me.x.Eval(env)
	case '-':
		return -me.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", me.op))
}

func (me binary) Eval(env Env) float64 {
	switch me.op {
	case '+':
		return me.x.Eval(env) + me.y.Eval(env)
	case '-':
		return me.x.Eval(env) - me.y.Eval(env)
	case '*':
		return me.x.Eval(env) * me.y.Eval(env)
	case '/':
		return me.x.Eval(env) / me.y.Eval(env)
	case '%':
		return math.Mod(me.x.Eval(env), me.y.Eval(env))
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", me.op))
}

func (me call) Eval(env Env) float64 {
	switch me.fn {
	case "pow":
		return math.Pow(me.args[0].Eval(env), me.args[1].Eval(env))
	case "sin":
		return math.Sin(me.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(me.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", me.fn))
}

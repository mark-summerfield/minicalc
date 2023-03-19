// Modifications Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strings"
)

func (me Var) Check(vars map[Var]bool) error {
	vars[me] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (me unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", me.op) {
		return fmt.Errorf("unexpected unary op %q", me.op)
	}
	return me.x.Check(vars)
}

func (me binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/%", me.op) {
		return fmt.Errorf("unexpected binary op %q", me.op)
	}
	if err := me.x.Check(vars); err != nil {
		return err
	}
	return me.y.Check(vars)
}

func (me call) Check(vars map[Var]bool) error {
	arity, ok := arityForFunc[me.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", me.fn)
	}
	if len(me.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			me.fn, len(me.args), arity)
	}
	for _, arg := range me.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var arityForFunc = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

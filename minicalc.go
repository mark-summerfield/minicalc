// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			message := fmt.Sprintf("Unrecoverable error: %s", r)
			fltk.MessageBox(fmt.Sprintf("Error — %s", appName), message)
			fmt.Println(message)
		}
	}()
	config := newConfig()
	fltk.SetScheme(config.Theme)
	fltk.SetScreenScale(0, config.Scale)
	app := newApp(config)
	app.Show()
	fltk.Run()
}

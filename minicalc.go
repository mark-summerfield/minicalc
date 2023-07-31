// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"

	"github.com/pwiecz/go-fltk"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] != "--debug" {
		defer func() {
			if r := recover(); r != nil {
				message := fmt.Sprintf("Unrecoverable error: %s", r)
				fltk.MessageBox(fmt.Sprintf("Error — %s", appName), message)
				fmt.Println(message)
			}
		}()
	}
	config := newConfig()
	fltk.SetScheme(config.Theme)
	fltk.SetScreenScale(0, config.Scale)
	app := newApp(config)
	app.Show()
	fltk.Run()
}

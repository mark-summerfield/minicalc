// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/mark-summerfield/gong"
)

type Config struct {
	filename           string
	X                  int
	Y                  int
	Width              int
	Height             int
	Scale              float32
	LastTab            int
	ShowIntialHelpText bool
	CustomTitle        string
	CustomHtml         string
}

func newConfig() *Config {
	filename, found := gong.GetIniFilename(appName)
	if found {
		cfg, err := ini.Load(filename)
		if err != nil {
			fmt.Println("newConfig #1", filename, err)
		} else {
			config := &Config{filename: filename, ShowIntialHelpText: true}
			err = cfg.MapTo(config)
			if err != nil {
				fmt.Println("newConfig #2", filename, err)
			} else {
				if config.CustomTitle == "" {
					config.CustomTitle = "Custom"
				}
				if config.CustomHtml == "" {
					config.CustomHtml = customPlaceHolderText
				}
				if config.Width < 100 || config.Width > 768 {
					config.Width = 512
				}
				if config.Height < 100 || config.Height > 768 {
					config.Height = 480
				}
				if config.Scale < 0.5 || config.Scale > 5 {
					config.Scale = 1
				}
				if config.LastTab < 0 || config.LastTab > aboutTabIndex {
					config.LastTab = 0
				}
				return config
			}

		}
	}
	config := &Config{filename: filename, X: -1, Width: 512, Height: 480,
		Scale: 1.0, ShowIntialHelpText: true, CustomTitle: "&Custom",
		CustomHtml: customPlaceHolderText}
	return config
}

func (me *Config) save() {
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, me)
	if err != nil {
		fmt.Println("save #1", me.filename, err)
	} else {
		err := cfg.SaveTo(me.filename)
		if err != nil {
			fmt.Println("save #2", me.filename, err)
		}
	}
}

// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/mark-summerfield/gong"
)

type Config struct {
	Filename   string
	X          int
	Y          int
	Width      int
	Height     int
	Scale      float32
	LastTab    int
	CustomHtml string
}

type config struct {
	X          int
	Y          int
	Width      int
	Height     int
	Scale      float32
	LastTab    int
	CustomHtml string
}

func newConfig() *Config {
	filename, found := gong.GetIniFilename(APPNAME)
	if found {
		cfg, err := ini.Load(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			config := new(config)
			err = cfg.MapTo(config)
			if err != nil {
				fmt.Println(err)
			} else {
				if config.CustomHtml == "" {
					config.CustomHtml = fmt.Sprintf(
						customPlaceHolderTemplate, filename)
				}
				return &Config{Filename: filename, X: config.X, Y: config.Y,
					Width: config.Width, Height: config.Height,
					Scale: config.Scale, LastTab: config.LastTab,
					CustomHtml: config.CustomHtml}
			}

		}
	}
	config := &Config{Filename: filename, X: -1, Width: 512, Height: 480,
		Scale:      1.0,
		CustomHtml: fmt.Sprintf(customPlaceHolderTemplate, filename)}
	return config
}

func (me *Config) save() {
	config := &config{X: me.X, Y: me.Y, Width: me.Width, Height: me.Height,
		Scale: me.Scale, LastTab: me.LastTab, CustomHtml: me.CustomHtml}
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, config)
	if err != nil {
		fmt.Println(err)
	} else {
		err := cfg.SaveTo(me.Filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}

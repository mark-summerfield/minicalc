// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

//go:embed images/icon16.png
var icon16data []byte

//go:embed images/icon32.png
var icon32data []byte

//go:embed images/icon64.png
var icon64data []byte

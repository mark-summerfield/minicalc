// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
    "fmt"
    _ "embed"
    )

//go:embed Version.dat
var Version string

func main() {
    fmt.Printf("Hello minicalc v%s\n", Version)
}

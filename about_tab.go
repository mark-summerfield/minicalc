// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/pwiecz/go-fltk"
)

func makeAboutTab(filename string, x, y, width, height int) *fltk.Group {
	group := fltk.NewGroup(x, y, width, height, "A&bout")
	view := fltk.NewHelpView(x, y, width, height)
	view.SetValue(aboutHtml(filename))
	group.End()
	return group
}

func aboutHtml(filename string) string {
	var year string
	y := time.Now().Year()
	if y == 2023 {
		year = fmt.Sprintf("%d", y)
	} else {
		year = fmt.Sprintf("2023-%d", y-2000)
	}
	distro := ""
	if runtime.GOOS == "linux" {
		if out, err := exec.Command("lsb_release",
			"-ds").Output(); err == nil {
			distro = strings.TrimSpace(string(out))
		}
	}
	return fmt.Sprintf(
		`<p><center><font face=sans size=6 color=navy><b>%s</b> v%s</font>
</center></p>
<p><center><font face=sans color=navy size=4>A little GUI
tool</font></center></p>
<p><center><font face=sans size=4>
<a href=\"https://github.com/mark-summerfield/minicalc\">https://github.com/mark-summerfield/minicalc</a>
</font></center></p>
<p><center>
<font face=sans size=4 color=green>
Copyright © %s Mark Summerfield.<br>
All rights reserved.<br>
License: GPLv3.</font>
</center></p>
<p>
<center>
<font face=sans size=4>
Configuration file: <tt>%s</tt><br>
(edit to change the Scale, the Custom tab, etc.)
</font>
</center>
</p>
<p><center><font face=sans size=4 color=#222>%s • %s/%s</font></center><br>
<center><font face=sans size=4 color=#222>go-fltk %s • FLTK
%s</font></center><br>
<center><font face=sans size=4 color=#222>%s</font></center></p>`,
		appName, Version, year, filename, runtime.Version(), runtime.GOOS,
		runtime.GOARCH, fltk.GoVersion(), fltk.Version(), distro)
}

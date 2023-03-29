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

func makeAboutTab(app *App, x, y, width, height int) {
	group := fltk.NewGroup(x, y, width, height, "A&bout")
	app.aboutView = fltk.NewHelpView(x, y, width, height)
	app.aboutView.TextFont(fltk.HELVETICA)
	app.aboutView.TextSize(app.config.ViewFontSize)
	app.aboutView.SetValue(aboutHtml(app.config.filename))
	group.End()
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
		`<center><h3><font color=navy>%s v%s</font></h3></center>
<p><center><font color=navy>A little GUI tool</font></center></p>
<p><center>
<a href=\"https://github.com/mark-summerfield/minicalc\">https://github.com/mark-summerfield/minicalc</a>
</center></p>
<p><center>
<font color=green>Copyright © %s Mark Summerfield.<br>
All rights reserved.<br>
License: GPLv3.
</center></p>
<p><center>Configuration file:<br><tt>%s</tt></center></p>
<p><center><font color=#222>%s • %s/%s</font></center><br>
<center><font color=#222>go-fltk %s • FLTK
%s</font></center><br>
<center><font color=#222>%s</font></center></p>`,
		appName, Version, year, filename, runtime.Version(), runtime.GOOS,
		runtime.GOARCH, fltk.GoVersion(), fltk.Version(), distro)
}

// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

//go:build windows

package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gonutz/w32/v2"
	"github.com/pwiecz/go-fltk"
)

func aboutHtml() string {
	var year string
	y := time.Now().Year()
	if y == 2023 {
		year = fmt.Sprintf("%d", y)
	} else {
		year = fmt.Sprintf("2023-%d", y-2000)
	}
	info := w32.RtlGetVersion()
	distro := fmt.Sprintf("Windows %d.%d", info.MajorVersion,
		info.MinorVersion)
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
<p><center><font color=#222>%s • %s/%s</font></center><br>
<center><font color=#222>go-fltk %s • FLTK
%s</font></center><br>
<center><font color=#222>%s</font></center></p>`,
		appName, Version, year, runtime.Version(), runtime.GOOS,
		runtime.GOARCH, fltk.GoVersion(), fltk.Version(), distro)
}

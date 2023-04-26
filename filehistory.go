// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"

	"github.com/pwiecz/go-fltk"
)

//go:embed Version.dat
var Version string

func main() {
	const (
		width      = 200
		height     = 80
		lineHeight = height / 2
	)
	fltk.SetScheme("oxy")
	window := fltk.NewWindow(width, height)
	window.SetLabel("File History " + Version[:len(Version)-1])
	vbox := fltk.NewFlex(0, 0, width, height)

	vbox.End()
	window.Show()
	fltk.Run()
}

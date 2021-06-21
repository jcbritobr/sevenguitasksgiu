package main

import (
	"github.com/AllenDang/giu"
)

var (
	counter int32
)

const (
	winHeight = 70
	winWidth  = 230
)

func loop() {
	giu.SingleWindow("Counter").Layout(
		giu.Row(
			giu.InputInt("", &counter).Size(50),
			giu.Button("Count").OnClick(func() {
				counter++
			}),
		),
	)
}

func main() {
	counter = 0
	wnd := giu.NewMasterWindow("Counter", winWidth, winHeight, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}

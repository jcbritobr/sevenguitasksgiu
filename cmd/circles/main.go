package main

import (
	"image"
	"image/color"

	"github.com/AllenDang/giu"
)

var (
	pos        = image.Point{}
	circleList = []image.Point{}
)

func loop() {
	giu.SingleWindow("canvas").Layout(
		giu.Custom(func() {
			canvas := giu.GetCanvas()
			pos = giu.GetMousePos()
			c := color.RGBA{255, 0, 0, 255}

			for _, position := range circleList {
				canvas.AddCircleFilled(position, 30, c)
			}

			if giu.IsMouseClicked(giu.MouseButtonLeft) {
				circleList = append(circleList, pos)
			}
		}),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Circles", 400, 400, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}

package main

import (
	"image"
	"image/color"

	"github.com/AllenDang/giu"
)

var (
	pos        = image.Point{}
	circleList = []image.Point{}
	counter    = 0
	index      = 0
)

func undo() {
	if index > 0 {
		index--
	}
}

func redo() {
	if index < counter {
		index++
	}
}

func loop() {
	giu.SingleWindow("canvas").Layout(
		giu.Custom(func() {
			canvas := giu.GetCanvas()
			pos = giu.GetMousePos()
			c := color.RGBA{255, 0, 0, 255}

			for _, position := range circleList[:index] {
				canvas.AddCircleFilled(position, 30, c)
			}

			if giu.IsMouseClicked(giu.MouseButtonLeft) && pos.Y > 70 {
				circleList = append(circleList, pos)
				counter++
				index++
			}
		}),
		giu.Row(
			giu.Button("Undo").Size(100, 25).OnClick(undo),
			giu.Button("Redo").Size(100, 25).OnClick(redo),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Circles", 400, 400, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}

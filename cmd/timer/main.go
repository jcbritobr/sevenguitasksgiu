package main

import (
	"time"

	"github.com/AllenDang/giu"
)

var (
	progress float32 = 0.0
	dragInt  int32   = 10
	maxCount int32   = 10
	count    int32   = 0
)

const (
	winHeight = 100
	winWidth  = 400
)

func onSliderChange() {
	maxCount = dragInt
}

func refresh() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		if count < maxCount {
			count++
			progress = (float32(count*100) / float32(maxCount)) / 100
			giu.Update()
		}
		<-ticker.C
	}
}

func reset() {
	count = 0
}

func loop() {
	giu.SingleWindow("##Timer").Layout(
		giu.ProgressBar(progress).Size(-1, 0).Overlay("Progress"),
		giu.SliderInt("Duration", &dragInt, 0, 30).OnChange(onSliderChange),
		giu.Button("Reset").Size(100, 25).OnClick(reset),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Timer", winWidth, winHeight, giu.MasterWindowFlagsNotResizable, nil)
	go refresh()
	wnd.Run(loop)
}

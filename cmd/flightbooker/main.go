package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	start         time.Time = time.Now()
	_return       time.Time = time.Now()
	selectionData int32     = 0
	comboData               = []string{"One-way flight", "Return flight"}
	isReturn      bool      = false
	errReturnDate bool      = false
)

const (
	winHeight = 300
	winWidth  = 500
)

func onComboChanged() {
	if selectionData != 0 {
		isReturn = true
	} else {
		isReturn = false
	}
}

func onDatePickChanged() {
	if start.After(_return) && selectionData != 0 {
		errReturnDate = true
	} else {
		errReturnDate = false
	}
}

func onBookClick() {
	if selectionData == 0 {
		giu.OpenPopup("Book##oneway")
	} else {
		giu.OpenPopup("Book##return")
	}
}

func loop() {
	giu.SingleWindow("Booker Flight").Layout(
		giu.Combo("flight", comboData[selectionData], comboData, &selectionData).OnChange(onComboChanged),
		giu.DatePicker("start", &start).Size(325).OnChange(onDatePickChanged),
		giu.Condition(isReturn, giu.Layout{giu.DatePicker("return", &_return).Size(325).OnChange(onDatePickChanged)}, nil),
		giu.Condition(
			errReturnDate,
			giu.Layout{
				giu.Style().
					SetColor(imgui.StyleColorText, color.RGBA{R: 255, G: 0, B: 0, A: 255}).
					To(giu.Label("Return date is lesser than\nstart date!")),
			},
			nil,
		),
		giu.Condition(!errReturnDate, giu.Layout{giu.Button("Book").Size(70, 25).OnClick(onBookClick)}, nil),
		giu.PopupModal("Book##oneway").Layout(
			giu.Label(fmt.Sprintf("You have booked a one way\nflight for %s", start.String())),
			giu.Button("Ok").OnClick(func() { giu.CloseCurrentPopup() }),
		),
		giu.PopupModal("Book##return").Layout(
			giu.Label(fmt.Sprintf("You have booked a return\nflight for %s\nto %s", start.String(), _return.String())),
			giu.Button("Ok").OnClick(func() { giu.CloseCurrentPopup() }),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Booker Flight", winWidth, winHeight, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}

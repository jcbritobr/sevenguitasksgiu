package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

const (
	winWidth  = 400
	winHeight = 160
)

var (
	filterPrefix        string
	inputName           string
	inputSurname        string
	dataListItem        []string
	dataListItemCurrent int
)

func removeIndex(list []string, index int) []string {
	return append(list[:index], list[index+1:]...)
}

func onCreate() {
	dataListItem = append(dataListItem, fmt.Sprintf("%s, %s", inputSurname, inputName))
	clearFields()
}

func onListChange(index int) {
	dataListItemCurrent = index
}

func onUpdate() {
	onDelete()
	onCreate()
}

func onDelete() {
	dataListItem = removeIndex(dataListItem, dataListItemCurrent)
}

func loop() {
	giu.SingleWindow("##Crud").Layout(
		giu.Row(
			giu.Label("Filter Prefix"),
			giu.InputText("##filter-prefix-input", &filterPrefix).Size(-1),
		),
		giu.Row(
			giu.ListBox("##list-item", dataListItem).Size(200, 100).OnChange(onListChange),
			giu.Column(
				giu.Row(
					giu.Label("Name   "),
					giu.InputText("##input-name", &inputName).Size(-1),
				),
				giu.Row(
					giu.Label("Surname"),
					giu.InputText("##input-surname", &inputSurname).Size(-1),
				),
			),
		),
		giu.Row(
			giu.Button("Create").OnClick(onCreate),
			giu.Button("Update").OnClick(onUpdate),
			giu.Button("Delete").OnClick(onDelete),
		),
	)
}

func clearFields() {
	inputName = ""
	inputSurname = ""
}

func main() {
	wnd := giu.NewMasterWindow("CRUD", winWidth, winHeight, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}

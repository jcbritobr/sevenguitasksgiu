package main

import "github.com/AllenDang/giu"

var (
	inputCelsius    int32
	inputFahrenheit int32
)

const (
	winHeight = 70
	winWidth  = 320
)

func onFahrenheitChange() {
	inputCelsius = int32((float32(inputFahrenheit) - 32.0) * float32(5.0/9.0))
}

func onCelsiusChange() {
	inputFahrenheit = inputCelsius*(9/5) + 32.0
}

func loop() {
	giu.SingleWindow("Temperature").Layout(
		giu.Row(
			giu.InputInt("Celsius =", &inputCelsius).
				Size(75.0).
				OnChange(onCelsiusChange),
			giu.InputInt("Fahrenheit", &inputFahrenheit).
				Size(75.0).
				OnChange(onFahrenheitChange),
		),
	)
}

func main() {
	inputCelsius = 0
	inputFahrenheit = 0
	wnd := giu.NewMasterWindow("Temperature Converter", winWidth, winHeight, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}

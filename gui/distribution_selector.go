package gui

import (
	"fyne.io/fyne/v2/widget"
	"log"
)

func distributionFormSelector(selected string) {
	switch selected {
	case "Normal(μ, σ)":
		widgetMeanEntry := widget.NewEntry()
		widgetStdDevEntry := widget.NewEntry()
		widgetMeanEntry.SetText("0")
		widgetStdDevEntry.SetText("1")
		log.Println("Selected Normal distribution: mean/std entries created")
	case "Uniform(a, b)":
		widgetLowerLimitEntry := widget.NewEntry()
		widgetUpperLimitEntry := widget.NewEntry()
		widgetLowerLimitEntry.SetText("0")
		widgetUpperLimitEntry.SetText("1")
		log.Println("Selected Uniform distribution: a/b entries created")
	default:
		log.Println("Selected distribution:", selected)
	}
}

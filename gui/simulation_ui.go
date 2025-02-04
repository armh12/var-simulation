package gui

import (
	"fyne.io/fyne/v2/widget"
	"log"
)

func simTypeSelection() *widget.Select {
	simTypeSelect := widget.NewSelect([]string{
		"Metropolis-Hastings Advanced",
	}, func(value string) {
		log.Println("Simulation selected:", value)
	})
	simTypeSelect.SetSelected("Metropolis-Hastings Advanced")
	return simTypeSelect
}

func mhInputValuesForm(deltaEntry, lowerLimitEntry, upperLimitEntry, numOfSamplesEntry *widget.Entry) *widget.Form {
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Delta:", Widget: deltaEntry},
			{Text: "Lower limit:", Widget: lowerLimitEntry},
			{Text: "Upper limit:", Widget: upperLimitEntry},
			{Text: "Iteration Count:", Widget: numOfSamplesEntry},
		},
	}
	return form
}

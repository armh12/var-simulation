package gui

import (
	"fyne.io/fyne/v2/app"
)

func StartGUI() {
	simApp := app.New()
	mainWindow := simApp.NewWindow("Variable Simulation")
	mainWindow.ShowAndRun()
}

package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func AppGUI() {
	simApp := app.New()
	mainWindow := simApp.NewWindow("Variable Simulation")

	simTypeSelect := simTypeSelection()

	targetDistParams := newDistParamsWidgets()
	proposalDistParams := newDistParamsWidgets()

	targetDistSelect := distributionSelection("Target", targetDistParams)
	proposalDistSelect := distributionSelection("Proposal", proposalDistParams)

	topBar := container.NewGridWithColumns(3,
		simTypeSelect,
		targetDistSelect,
		proposalDistSelect,
	)

	deltaEntry := widget.NewEntry()
	deltaEntry.SetText("0.1")

	lowerLimitEntry := widget.NewEntry()
	lowerLimitEntry.SetText("0")

	upperLimitEntry := widget.NewEntry()
	upperLimitEntry.SetText("10")

	numOfSamplesEntry := widget.NewEntry()
	numOfSamplesEntry.SetText("1000")

	form := mhInputValuesForm(deltaEntry, lowerLimitEntry, upperLimitEntry, numOfSamplesEntry)

	placeholderHist := canvas.NewText("Histogram", nil)
	placeholderTrace := canvas.NewText("Trace Plot", nil)
	placeholderRunningMean := canvas.NewText("Running Mean", nil)

	graphContainer := container.NewAppTabs(
		container.NewTabItem("Histogram", container.NewCenter(placeholderHist)),
		container.NewTabItem("Trace Plot", container.NewCenter(placeholderTrace)),
		container.NewTabItem("Running Mean", container.NewCenter(placeholderRunningMean)),
	)
	graphContainer.SetTabLocation(container.TabLocationBottom)

	simulateButton := widget.NewButton("Simulate", func() {
		deltaVal, err := strconv.ParseFloat(deltaEntry.Text, 64)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Not correct delta value: %v", err), mainWindow)
			return
		}
		lowerVal, err := strconv.ParseFloat(lowerLimitEntry.Text, 64)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Not correct lower limit value: %v", err), mainWindow)
			return
		}
		upperVal, err := strconv.ParseFloat(upperLimitEntry.Text, 64)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Not correct upper limit value: %v", err), mainWindow)
			return
		}
		numSamples, err := strconv.Atoi(numOfSamplesEntry.Text)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Not correct iterations count value: %v", err), mainWindow)
			return
		}

		targetFunc := getDistributionFunction(targetDistSelect, targetDistParams, mainWindow)
		proposalFunc := getDistributionFunction(proposalDistSelect, proposalDistParams, mainWindow)

		go func() {
			results, err := SimulateMH(
				targetFunc,
				proposalFunc,
				deltaVal,
				lowerVal,
				upperVal,
				numSamples,
			)
			if err != nil {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Simulation Error",
					Content: err.Error(),
				})
				return
			}

			func() {
				placeholderHist.Text = fmt.Sprintf(
					"Simulation Ready!\nResults: %d\nExample: %.4f ...",
					len(results), results[0],
				)
				placeholderHist.Refresh()
			}()
		}()
	})

	plotTargetButton := widget.NewButton("Plot Target Dist", func() {
		plotDistribution(mainWindow, "Target", targetDistSelect, targetDistParams)
	})
	plotProposalButton := widget.NewButton("Plot Proposal Dist", func() {
		plotDistribution(mainWindow, "Proposal", proposalDistSelect, proposalDistParams)
	})

	// Собираем форму с параметрами дистрибуции для Target и Proposal + кнопки
	targetDistBox := container.NewVBox(
		widget.NewLabel("Target Distribution Parameters"),
		targetDistParams.form,
		plotTargetButton,
	)

	proposalDistBox := container.NewVBox(
		widget.NewLabel("Proposal Distribution Parameters"),
		proposalDistParams.form,
		plotProposalButton,
	)
	centerSplit := container.NewHSplit(
		form,
		graphContainer,
	)
	centerSplit.SetOffset(0.3)
	bottomBox := container.NewGridWithColumns(2,
		targetDistBox,
		proposalDistBox,
	)

	mainContent := container.NewBorder(
		topBar,    // top
		bottomBox, // bottom
		nil,       // left
		nil,       // right
		container.NewBorder(
			nil,
			simulateButton, // под формой MH будет кнопка Simulate
			nil,
			nil,
			centerSplit,
		),
	)

	mainWindow.SetContent(mainContent)
	mainWindow.Resize(fyne.NewSize(1280, 720))
	mainWindow.ShowAndRun()
}

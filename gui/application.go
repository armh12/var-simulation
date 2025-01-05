package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"var-simulation/types"
)

func simTypeSelection() *widget.Select {
	simTypeSelect := widget.NewSelect([]string{
		"Metropolis-Hastings Advanced",
		"Monte-Carlo",
	}, func(value string) {
		log.Println("Simulation selected:", value)
	})
	simTypeSelect.SetSelected("Metropolis-Hastings")
	return simTypeSelect
}

func distributionSelection(distributionType string) *widget.Select {
	distributionNames := []string{
		"Normal(μ, σ)",
		"Uniform(a, b)",
		"Cauchy(location, scale)",
		"LogNormal(mean, stdDev)",
		"Exponential(mean)",
		"Weibull(scale, shape)",
		"Pareto(scale, shape)",
		"Gamma(shape, scale)",
	}
	DistSelect := widget.NewSelect(distributionNames, distributionFormSelector)
	DistSelect.SetSelected("Normal(μ, σ)")
	return DistSelect
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

func AppGUI() {
	simApp := app.New()
	mainWindow := simApp.NewWindow("Variable Simulation")

	simTypeSelect := simTypeSelection()
	targetDistSelect := distributionSelection("Target")
	proposalDistSelect := distributionSelection("Proposal")

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

		targetFunc := types.Function(func(x float64) float64 {
			return 1.0
		})
		proposalFunc := types.Function(func(x float64) float64 {
			return 1.0
		})

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
				placeholderHist.Text = fmt.Sprintf("Simulation Ready!\nResults: %d\nExample: %.4f ...",
					len(results), results[0])
				placeholderHist.Refresh()
			}()
		}()
	})

	centerSplit := container.NewHSplit(
		form,
		graphContainer,
	)
	centerSplit.SetOffset(0.3)

	mainContent := container.NewBorder(
		topBar,         // top
		simulateButton, // bottom
		nil,            // left
		nil,            // right
		centerSplit,    // center
	)

	mainWindow.SetContent(mainContent)
	mainWindow.Resize(fyne.NewSize(1920, 1080))
	mainWindow.ShowAndRun()
}

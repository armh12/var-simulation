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

func AppGUI() {
	simApp := app.New()
	mainWindow := simApp.NewWindow("Variable Simulation")

	simTypeSelect := widget.NewSelect([]string{
		"Metropolis-Hastings Advanced",
		"Monte-Carlo",
	}, func(value string) {
		log.Println("Simulation selected:", value)
	})
	simTypeSelect.SetSelected("Metropolis-Hastings") // по умолчанию

	targetDistSelect := widget.NewSelect([]string{
		"Normal(μ, σ)",
		"Uniform(a, b)",
	}, func(value string) {
		log.Println("Target distribution selected:", value)
	})
	targetDistSelect.SetSelected("Normal(μ, σ)")

	proposalDistSelect := widget.NewSelect([]string{
		"Normal(μ, σ)",
		"Uniform(a, b)",
	}, func(value string) {
		log.Println("Proposal distribution selected:", value)
	})
	proposalDistSelect.SetSelected("Normal(μ, σ)")

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

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Delta:", Widget: deltaEntry},
			{Text: "Lower limit:", Widget: lowerLimitEntry},
			{Text: "Upper limit:", Widget: upperLimitEntry},
			{Text: "Количество итераций:", Widget: numOfSamplesEntry},
		},
	}

	placeholderGraph1 := canvas.NewText("First Graph", nil)
	placeholderGraph2 := canvas.NewText("Second Graph", nil)

	tabContainer := container.NewAppTabs(
		container.NewTabItem("График 1", container.NewCenter(placeholderGraph1)),
		container.NewTabItem("График 2", container.NewCenter(placeholderGraph2)),
	)
	tabContainer.SetTabLocation(container.TabLocationBottom)

	simulateButton := widget.NewButton("Simulate", func() {
		// 1) Считываем параметры (в реальном коде делайте валидацию, обработку ошибок)
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
				placeholderGraph1.Text = fmt.Sprintf("Simulation Ready!\nResults: %d\nExample: %.4f ...",
					len(results), results[0])
				placeholderGraph1.Refresh()
			}()
		}()
	})

	centerSplit := container.NewHSplit(
		form,
		tabContainer,
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
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}

package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/expr-lang/expr"
	"log"
	"strconv"
	"var-simulation/math_utils"
	"var-simulation/parser"
)

type distParamsWidgets struct {
	meanEntry        *widget.Entry
	stdDevEntry      *widget.Entry
	lowerLimitEntry  *widget.Entry
	upperLimitEntry  *widget.Entry
	scaleEntry       *widget.Entry
	shapeEntry       *widget.Entry
	location         *widget.Entry
	userDefinedEntry *widget.Entry

	form *widget.Form
}

func newDistParamsWidgets() *distParamsWidgets {
	dp := &distParamsWidgets{
		meanEntry:        widget.NewEntry(),
		stdDevEntry:      widget.NewEntry(),
		lowerLimitEntry:  widget.NewEntry(),
		upperLimitEntry:  widget.NewEntry(),
		scaleEntry:       widget.NewEntry(),
		shapeEntry:       widget.NewEntry(),
		location:         widget.NewEntry(),
		userDefinedEntry: widget.NewEntry(),
	}

	dp.meanEntry.SetText("0")
	dp.stdDevEntry.SetText("1")
	dp.lowerLimitEntry.SetText("0")
	dp.upperLimitEntry.SetText("1")
	dp.scaleEntry.SetText("1")
	dp.shapeEntry.SetText("1")
	dp.location.SetText("0")

	// Создаём форму (Form) со всеми возможными элементами,
	// но их отображение будем управлять позже (через .Hide() / .Show()).
	dp.form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Mean (μ):", Widget: dp.meanEntry},
			{Text: "Std Dev (σ):", Widget: dp.stdDevEntry},
			{Text: "Lower (a):", Widget: dp.lowerLimitEntry},
			{Text: "Upper (b):", Widget: dp.upperLimitEntry},
			{Text: "Scale (σ):", Widget: dp.scaleEntry},
			{Text: "Shape (α):", Widget: dp.shapeEntry},
			{Text: "Location (μ):", Widget: dp.location},
			{Text: "User defined distribution", Widget: dp.userDefinedEntry},
		},
	}
	return dp
}

func (dp *distParamsWidgets) updateVisibility(selectedDist string) {
	// По умолчанию скрываем все поля.
	dp.form.Items[0].Widget.Hide() // mean
	dp.form.Items[1].Widget.Hide() // stdDev
	dp.form.Items[2].Widget.Hide() // lowerLimit
	dp.form.Items[3].Widget.Hide() // upperLimit
	dp.form.Items[4].Widget.Hide() // scale
	dp.form.Items[5].Widget.Hide() // shape
	dp.form.Items[6].Widget.Hide() // location
	dp.form.Items[7].Widget.Hide() // userDefined

	switch selectedDist {
	case "Normal(μ, σ)":
		dp.form.Items[0].Widget.Show()
		dp.form.Items[1].Widget.Show()
	case "Uniform(a, b)":
		dp.form.Items[2].Widget.Show()
		dp.form.Items[3].Widget.Show()
	case "Cauchy(location, scale)":
		dp.form.Items[6].Text = "Location"
		dp.form.Items[4].Text = "Scale"
		dp.form.Items[6].Widget.Show()
		dp.form.Items[4].Widget.Show()
	case "LogNormal(mean, stdDev)":
		dp.form.Items[0].Widget.Show()
		dp.form.Items[1].Widget.Show()
	case "Exponential(mean)":
		dp.form.Items[0].Text = "Mean"
		dp.form.Items[0].Widget.Show()
	case "Weibull(scale, shape)":
		dp.form.Items[4].Text = "Scale"
		dp.form.Items[5].Text = "Shape"
		dp.form.Items[4].Widget.Show()
		dp.form.Items[5].Widget.Show()
	case "Pareto(scale, shape)":
		dp.form.Items[4].Text = "Scale"
		dp.form.Items[5].Text = "Shape"
		dp.form.Items[4].Widget.Show()
		dp.form.Items[5].Widget.Show()
	case "Gamma(shape, scale)":
		dp.form.Items[4].Text = "Shape"
		dp.form.Items[5].Text = "Scale"
		dp.form.Items[4].Widget.Show()
		dp.form.Items[5].Widget.Show()
	case "Custom distribution":
		dp.form.Items[7].Text = "User defined distribution"
		dp.form.Items[7].Widget.Show()
	default:
		log.Println("Selected distribution:", selectedDist)
	}
}

var distributionNames = []string{
	"Normal(μ, σ)",
	"Uniform(a, b)",
	"Cauchy(location, scale)",
	"LogNormal(mean, stdDev)",
	"Exponential(mean)",
	"Weibull(scale, shape)",
	"Pareto(scale, shape)",
	"Gamma(shape, scale)",
	"Custom distribution",
}

func distributionSelection(distType string, dp *distParamsWidgets) *widget.Select {
	distSelect := widget.NewSelect(distributionNames, func(selected string) {
		log.Printf("[%s] distribution selected: %s\n", distType, selected)
		dp.updateVisibility(selected)
	})
	distSelect.SetSelected("Normal(μ, σ)")
	dp.updateVisibility("Normal(μ, σ)")
	return distSelect
}

func plotDistribution(parent fyne.Window, title string, distSelect *widget.Select, dp *distParamsWidgets) {
	selected := distSelect.Selected
	win := fyne.CurrentApp().NewWindow(fmt.Sprintf("Plot: %s", title))
	infoText := canvas.NewText(fmt.Sprintf("Plotting %s distribution...", selected), nil)

	switch selected {
	case "Normal(μ, σ)":
		mean, _ := strconv.ParseFloat(dp.meanEntry.Text, 64)
		stdDev, _ := strconv.ParseFloat(dp.stdDevEntry.Text, 64)
		infoText.Text = fmt.Sprintf("Plot Normal(μ=%.2f, σ=%.2f)", mean, stdDev)
	case "Uniform(a, b)":
		a, _ := strconv.ParseFloat(dp.lowerLimitEntry.Text, 64)
		b, _ := strconv.ParseFloat(dp.upperLimitEntry.Text, 64)
		infoText.Text = fmt.Sprintf("Plot Uniform(a=%.2f, b=%.2f)", a, b)
	default:
		infoText.Text = fmt.Sprintf("Plotting: %s (params read from entries)", selected)
	}

	infoText.Alignment = fyne.TextAlignCenter
	content := container.NewCenter(infoText)
	win.SetContent(content)
	win.Resize(fyne.NewSize(400, 300))
	win.Show()
}

func getDistributionFunction(
	distSelect *widget.Select,
	dp *distParamsWidgets,
	parentWindow fyne.Window,
) math_utils.DistributionFunc {
	distributionFuncs := math_utils.Distributions{}
	selected := distSelect.Selected

	switch selected {
	case "Normal(μ, σ)":
		meanStr := dp.meanEntry.Text
		sigmaStr := dp.stdDevEntry.Text

		mean, err := strconv.ParseFloat(meanStr, 64)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Error parse mean: %v", err), parentWindow)
			return func(x float64) float64 { return 0.0 }
		}
		sigma, err := strconv.ParseFloat(sigmaStr, 64)
		if err != nil || sigma <= 0 {
			dialog.ShowError(fmt.Errorf("Error parse sigma: %v", err), parentWindow)
			return func(x float64) float64 { return 0.0 }
		}
		return func(x float64) float64 {
			// x is ignored since our distribution function doesn't need it.
			return distributionFuncs.GaussianDistribution(mean, sigma)
		}

	// And similarly for the other cases, wrapping each in a closure that ignores x:
	case "Uniform(a, b)":
		aStr := dp.lowerLimitEntry.Text
		bStr := dp.upperLimitEntry.Text
		aVal, err := strconv.ParseFloat(aStr, 64)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Error parse lower limit: %v", err), parentWindow)
			return func(x float64) float64 { return 0.0 }
		}
		bVal, err := strconv.ParseFloat(bStr, 64)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Error parse upper limit: %v", err), parentWindow)
			return func(x float64) float64 { return 0.0 }
		}
		if bVal <= aVal {
			dialog.ShowError(fmt.Errorf("b <= a in Uniform(a, b)"), parentWindow)
			return func(x float64) float64 { return 0.0 }
		}
		return func(x float64) float64 {
			return distributionFuncs.UniformDistribution(aVal, bVal)
		}

	case "User-defined":
		exprText := dp.userDefinedEntry.Text
		fn, err := parser.ParseUserInput(exprText)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Parsing user-defined distribution: %v", err), parentWindow)
			return func(x float64) float64 { return 0.0 }
		}
		// If your user-defined function already accepts a float64,
		// you can just return it. Otherwise, wrap it:
		return func(x float64) float64 {
			result, err := expr.Run(fn, map[string]interface{}{"x": x})
			if err != nil {
				dialog.ShowError(fmt.Errorf("Running user-defined distribution: %v", err), parentWindow)
				return 0.0
			}
			return result.(float64)
		}
	}
	return func(x float64) float64 { return 0.0 }
}

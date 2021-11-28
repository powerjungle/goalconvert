package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fmt"
	"strconv"

	"github.com/powerjungle/goalconvert/alconvert"
)

func alcCalcAll(alcval *alconvert.Alcovalues) {
	alconvert.CalcGotUnits(alcval)
	alconvert.CalcTargetUnits(alcval)
	alconvert.CalcTargetPercent(alcval)
	alconvert.CalcTargetMl(alcval)
}

type allLabels struct {
	unitsLabel      *widget.Label
	finalMlLabel    *widget.Label
	finalRemAmLabel *widget.Label
	finalTarPerc    *widget.Label
	finalTarPercA   *widget.Label
	finalTarMlP     *widget.Label
	finalTarMlD     *widget.Label
}

func initAllLabels() *allLabels {
	return &allLabels{
		unitsLabel:      widget.NewLabel("0"),
		finalMlLabel:    widget.NewLabel("0"),
		finalRemAmLabel: widget.NewLabel("0"),
		finalTarPerc:    widget.NewLabel("0"),
		finalTarPercA:   widget.NewLabel("0"),
		finalTarMlP:     widget.NewLabel("0"),
		finalTarMlD:     widget.NewLabel("0"),
	}
}

func resetAllLabels(alcoval *alconvert.Alcovalues, aLa *allLabels) {
	aLa.unitsLabel.SetText(strconv.FormatFloat(float64(alconvert.GotUnits(alcoval)), 'f', -1, 32))
	aLa.finalMlLabel.SetText(strconv.FormatFloat(float64(alconvert.FinalMl(alcoval)), 'f', -1, 32))
	aLa.finalRemAmLabel.SetText(strconv.FormatFloat(float64(alconvert.FinalRemoveAmount(alcoval)), 'f', -1, 32))
	aLa.finalTarPerc.SetText(strconv.FormatFloat(float64(alconvert.FinalTargetPercent(alcoval)), 'f', -1, 32))
	aLa.finalTarPercA.SetText(strconv.FormatFloat(float64(alconvert.FinalTargetPercentAll(alcoval)), 'f', -1, 32))
	aLa.finalTarMlP.SetText(strconv.FormatFloat(float64(alconvert.FinalTargetMlPercent(alcoval)), 'f', -1, 32))
	aLa.finalTarMlD.SetText(strconv.FormatFloat(float64(alconvert.FinalTargetMlDiff(alcoval)), 'f', -1, 32))
}

type inputWidgets struct {
	amountLabel  *widget.Label
	amountSlider *widget.Slider
	amountEntry  *widget.Entry
}

func fromInputToAlcv(label string, alcoval *alconvert.Alcovalues, changeTo float64) {
	switch label {
	case "Milliliters":
		alcoval.UserSet.Milliliters = float32(changeTo)
		break
	case "Percentage":
		alcoval.UserSet.Percent = float32(changeTo)
		break
	case "Unit Target":
		alcoval.UserSet.UnitTarget = float32(changeTo)
		break
	case "Percentage Target":
		alcoval.UserSet.PercenTarget = float32(changeTo)
		break
	case "Milliliter Target":
		alcoval.UserSet.TargetMl = float32(changeTo)
		break
	}
}

func newInputWidgets(label string,
	rangeMin float64,
	rangeMax float64,
	alcoval *alconvert.Alcovalues,
	aLa *allLabels) *inputWidgets {
	iw := inputWidgets{
		amountLabel:  widget.NewLabel(label),
		amountSlider: widget.NewSlider(rangeMin, rangeMax),
		amountEntry:  widget.NewEntry(),
	}

	iw.amountEntry.SetText("0")

	iw.amountSlider.Step = 1
	iw.amountSlider.OnChanged = func(slideVal float64) {
		iw.amountEntry.SetText(strconv.FormatFloat(slideVal, 'f', -1, 64))
		fromInputToAlcv(label, alcoval, slideVal)
		alcCalcAll(alcoval)
		resetAllLabels(alcoval, aLa)
	}

	iw.amountEntry.OnChanged = func(inputStr string) {
		if f, err := strconv.ParseFloat(inputStr, 32); err == nil {
			fromInputToAlcv(label, alcoval, f)
		} else {
			fmt.Println(err)
		}

		if f, err := strconv.ParseFloat(inputStr, 64); err == nil {
			iw.amountSlider.SetValue(f)
		} else {
			fmt.Println(err)
		}
	}

	return &iw
}

func main() {
	av := alconvert.NewAV()

	alcApp := app.New()
	alcWindow := alcApp.NewWindow("goalconvert")
	alcWindow.Resize(fyne.NewSize(300, 0))

	aLa := initAllLabels()

	mlInputWidgets := newInputWidgets("Milliliters", 0, 2000, av, aLa)

	percInputWidgets := newInputWidgets("Percentage", 0, 100, av, aLa)

	unitTarInputWidgets := newInputWidgets("Unit Target", 0, 20, av, aLa)
	unitTarInputWidgets.amountSlider.Step = 0.1

	percTarInputWidgets := newInputWidgets("Percentage Target", 0, 20, av, aLa)

	mlTarInputWidgets := newInputWidgets("Milliliter Target", 0, 2000, av, aLa)

	alcWindow.SetContent(
		container.NewVBox(
			mlInputWidgets.amountLabel,
			mlInputWidgets.amountEntry,
			mlInputWidgets.amountSlider,
			percInputWidgets.amountLabel,
			percInputWidgets.amountEntry,
			percInputWidgets.amountSlider,
			widget.NewLabel("Calculated Units using ml and %"),
			aLa.unitsLabel,
			unitTarInputWidgets.amountLabel,
			unitTarInputWidgets.amountEntry,
			unitTarInputWidgets.amountSlider,
			widget.NewLabel("Calculated Unit Target"),
			container.NewHBox(
				widget.NewLabel("ml for target units"),
				layout.NewSpacer(),
				widget.NewLabel("removed ml for target"),
			),
			container.NewHBox(
				aLa.finalMlLabel,
				layout.NewSpacer(),
				aLa.finalRemAmLabel,
			),
			percTarInputWidgets.amountLabel,
			percTarInputWidgets.amountEntry,
			percTarInputWidgets.amountSlider,
			widget.NewLabel("Calculated Percent Target"),
			container.NewHBox(
				widget.NewLabel("Amount of Water To Add"),
				layout.NewSpacer(),
				widget.NewLabel("Total Amount Left"),
			),
			container.NewHBox(
				aLa.finalTarPerc,
				layout.NewSpacer(),
				aLa.finalTarPercA,
			),
			mlTarInputWidgets.amountLabel,
			mlTarInputWidgets.amountEntry,
			mlTarInputWidgets.amountSlider,
			widget.NewLabel("Calculated Percent Target"),
			container.NewHBox(
				widget.NewLabel("After Adding Water (%)"),
				layout.NewSpacer(),
				widget.NewLabel("Difference between ml"),
			),
			container.NewHBox(
				aLa.finalTarMlP,
				layout.NewSpacer(),
				aLa.finalTarMlD,
			),
		),
	)

	alcWindow.ShowAndRun()
}

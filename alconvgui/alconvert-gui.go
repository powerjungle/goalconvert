package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

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

func newInputWidgets(label string, rangeMin float64, rangeMax float64,
	alcoval *alconvert.Alcovalues, aLa *allLabels) *inputWidgets {

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
		} else if inputStr != "" {
			fmt.Println(err)
		}

		if f, err := strconv.ParseFloat(inputStr, 64); err == nil {
			iw.amountSlider.SetValue(f)
		} else if inputStr != "" {
			fmt.Println(err)
		}
	}

	return &iw
}

func makeIOCanvasObjects(iWs1 *inputWidgets, iWs2 *inputWidgets,
	calcValueLabel1 *widget.Label, calcValueLabel2 *widget.Label,
	calculatedLabel string, firstResultLabel string, secondResultLabel string) []fyne.CanvasObject {

	label1 := widget.NewLabel(firstResultLabel)
	label2 := widget.NewLabel(secondResultLabel)
	label2spacer := layout.NewSpacer()

	if firstResultLabel == "nope" {
		label1.Hide()
	}

	if secondResultLabel == "nope" {
		label2.Hide()
		label2spacer.Hide()
	}

	co := []fyne.CanvasObject{}

	line := canvas.NewLine(color.White)

	cL := widget.NewLabel(calculatedLabel)
	cL.TextStyle.Bold = true
	co = append(co, cL)

	if iWs1 != nil {
		co = append(co, iWs1.amountLabel, iWs1.amountEntry, iWs1.amountSlider)
	}

	if iWs2 != nil {
		co = append(co, iWs2.amountLabel, iWs2.amountEntry, iWs2.amountSlider)
	}

	cont1 := container.NewHBox(label1, label2spacer, label2)

	cont2obj := []fyne.CanvasObject{}

	if calcValueLabel1 != nil {
		calcValueLabel1.TextStyle.Bold = true
		cont2obj = append(cont2obj, calcValueLabel1)
	}

	cont2obj = append(cont2obj, label2spacer)

	if calcValueLabel2 != nil {
		calcValueLabel2.TextStyle.Bold = true
		cont2obj = append(cont2obj, calcValueLabel2)
	}

	cont2 := container.NewHBox(cont2obj...)

	co = append(co, cont1, cont2, line)

	return co
}

func compileCanvasObjects(objs ...[]fyne.CanvasObject) []fyne.CanvasObject {
	allObjs := []fyne.CanvasObject{}
	for _, v := range objs {
		for _, v2 := range v {
			allObjs = append(allObjs, v2)
		}
	}
	return allObjs
}

func main() {
	av := alconvert.NewAV()

	alcApp := app.New()
	alcWindow := alcApp.NewWindow("goalconvert")
	alcWindow.Resize(fyne.NewSize(600, 600))

	aLa := initAllLabels()

	// Init input widgets
	mlInputWidgets := newInputWidgets("Milliliters", 0, 2000, av, aLa)

	percInputWidgets := newInputWidgets("Percentage", 0, 100, av, aLa)

	unitTarInputWidgets := newInputWidgets("Unit Target", 0, 20, av, aLa)
	unitTarInputWidgets.amountSlider.Step = 0.1

	percTarInputWidgets := newInputWidgets("Percentage Target", 0, 100, av, aLa)

	mlTarInputWidgets := newInputWidgets("Milliliter Target", 0, 2000, av, aLa)
	///////////////////////

	// Init input CanvasObjects
	gotUnitsObjects := makeIOCanvasObjects(
		mlInputWidgets,
		percInputWidgets,
		aLa.unitsLabel,
		nil,
		"Calculated Units",
		"Calculated Units using ml and %",
		"nope",
	)

	gotUnitTargetObjects := makeIOCanvasObjects(
		unitTarInputWidgets,
		nil,
		aLa.finalMlLabel,
		aLa.finalRemAmLabel,
		"Calculated Unit Target",
		"ml for target units",
		"removed ml for target",
	)

	gotPercTargetObjects := makeIOCanvasObjects(
		percTarInputWidgets,
		nil,
		aLa.finalTarPerc,
		aLa.finalTarPercA,
		"Calculated Percent Target",
		"Amount of Water To Add",
		"Total Amount Left",
	)

	gotMlTargetObjects := makeIOCanvasObjects(
		mlTarInputWidgets,
		nil,
		aLa.finalTarMlP,
		aLa.finalTarMlD,
		"Calculated Milliliter Target",
		"After Adding Water (%)",
		"Difference between ml",
	)
	///////////////////////

	allObjs := compileCanvasObjects(
		gotUnitsObjects,
		gotUnitTargetObjects,
		gotPercTargetObjects,
		gotMlTargetObjects,
	)

	makeContainer := container.NewVBox(allObjs...)
	scrollCont := container.NewVScroll(makeContainer)

	alcWindow.SetContent(scrollCont)

	alcWindow.ShowAndRun()
}

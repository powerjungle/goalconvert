package main

import (
	"flag"
	"fmt"
	"github.com/powerjungle/goalconvert/alconvert"
)

var (
	ml      = flag.Float64("ml", 0, "input milliliters\na.k.a. 'Milliliters'")
	perc    = flag.Float64("perc", 0, "input percentage (concentration)\na.k.a. 'Percent'")
	taruni  = flag.Float64("taruni", 0, "needed units (target units)\na.k.a. 'UnitTarget'")
	tarperc = flag.Float64("tarperc", 0,
		"needed percentage (target percentage/concentration)\na.k.a. 'PercenTarget'")
	tarml      = flag.Float64("tarml", 0, "needed milliliters (target ml/amount)\na.k.a. 'TargetMl'")
	calcuni    = flag.Bool("calcuni", false, "calculate the units by using ml and perc\na.k.a. 'CalcGotUnits()'")
	calctaruni = flag.Bool("calctaruni", false,
		"calculate the target units by using ml, perc and taruni\na.k.a. 'CalcTargetUnits()'")
	calctarperc = flag.Bool("calctarperc", false,
		"calculate the target percentage by using ml, perc and tarperc\na.k.a. 'CalcTargetPercent()'")
	calctarml = flag.Bool("calctarml", false,
		"calculate the target ml by using ml, perc, tarml\na.k.a. 'CalcTargetMl()'")
)

func main() {
	flag.Parse()

	av := alconvert.NewAV()

	av.UserSet.Milliliters = float32(*ml)
	av.UserSet.Percent = float32(*perc)
	av.UserSet.UnitTarget = float32(*taruni)
	av.UserSet.PercenTarget = float32(*tarperc)
	av.UserSet.TargetMl = float32(*tarml)

	if *calcuni || *ml != 0 && *perc != 0 {
		alconvert.CalcGotUnits(av)

		if *calctaruni || *taruni != 0 {
			alconvert.CalcTargetUnits(av)
		}

		if *calctarperc || *tarperc != 0 {
			alconvert.CalcTargetPercent(av)
		}

		if *calctarml || *tarml != 0 {
			alconvert.CalcTargetMl(av)
		}

		alconvert.PrintForHumans(av)
	} else {
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Example 1: ./goalconvert -ml 2000 -perc 5")
		fmt.Println("Example 2: ./goalconvert -ml 250 -perc 40 -tarperc 5")
	}
}

package main

import (
	"github.com/powerjungle/goalconvert/alconvert"
	"flag"
)

var (
	ml = flag.Int("ml", 0, "input milliliters\na.k.a. 'Milliliters'")
	perc = flag.Float64("perc", 0, "input percentage (concentration)\na.k.a. 'Percent'")
	taruni = flag.Float64("taruni", 0, "needed units (target units)\na.k.a. 'UnitTarget'")
	tarperc = flag.Float64("tarperc", 0, "needed percentage (target percentage/concentration)\na.k.a. 'PercenTarget'")
	tarml = flag.Int("tarml", 0, "needed milliliters (target ml/amount)\na.k.a. 'TargetMl'")
	calcuni = flag.Bool("calcuni", false, "calculate the units by using ml and perc\na.k.a. 'CalcGotUnits()'")
	calctaruni = flag.Bool("calctaruni", false, "calculate the target units by using ml, perc and taruni\na.k.a. 'CalcTargetUnits()'")
	calctarperc = flag.Bool("calctarperc", false, "calculate the target percentage by using ml, perc and tarperc\na.k.a. 'CalcTargetPercent()'")
	calctarml = flag.Bool("calctarml", false, "calculate the target ml by using ml, perc, tarml\na.k.a. 'CalcTargetMl()'")
	printjson = flag.Bool("json", false, "print all the alcohol values as json")
)

func main() {
	flag.Parse()

	av := alconvert.NewAV()

	av.Milliliters = int16(*ml)
	av.Percent = float32(*perc)
	av.UnitTarget = float32(*taruni)
	av.PercenTarget = float32(*tarperc)
	av.TargetMl = int16(*tarml)

	if *calcuni || *ml != 0 && *perc != 0 {
        	alconvert.CalcGotUnits(av)
	}

	if *calctaruni || *taruni != 0 {
        	alconvert.CalcTargetUnits(av)
	}

	if *calctarperc || *tarperc != 0 {
		alconvert.CalcTargetPercent(av)
	}

	if *calctarml || *tarml != 0 {
		alconvert.CalcTargetMl(av)
	}

	if *printjson {
		alconvert.PrintJson(av)
	} else {
		alconvert.PrintForHumans(av)
	}
}

package main

import (
	"github.com/powerjungle/goalconvert/alconvert"
	"flag"
)

var (
	ml = flag.Int("ml", 0, "input milliliters - aka 'Milliliters'")
	perc = flag.Float64("perc", 0, "input percentage (concentration) - aka 'Percent'")
	taruni = flag.Float64("taruni", 0, "needed units (target units) - aka 'UnitTarget'")
	tarperc = flag.Float64("tarperc", 0, "needed percentage (target percentage/concentration) - aka 'PercenTarget'")
	tarml = flag.Int("tarml", 0, "needed milliliters (target ml/amount) - aka 'TargetMl'")
	calcuni = flag.Bool("calcuni", false, "calculate the units by using ml and perc - aka 'CalcGotUnits()'")
	calctaruni = flag.Bool("calctaruni", false, "calculate the target units by using ml, perc and taruni - aka 'CalcTargetUnits()'")
	calctarperc = flag.Bool("calctarperc", false, "calculate the target percentage by using ml, perc and tarperc - aka 'CalcTargetPercent()'")
	calctarml = flag.Bool("calctarml", false, "calculate the target ml by using ml, perc, tarml - aka 'CalcTargetMl()'")
	human = flag.Bool("human", false, "print human readable results (sorta)")
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

	if *calcuni {
        	alconvert.CalcGotUnits(av)
	} else if *calctaruni {
        	alconvert.CalcTargetUnits(av)
	} else if *calctarperc {
		alconvert.CalcTargetPercent(av)
	} else if *calctarml {
		alconvert.CalcTargetMl(av)
	}

	if *human {
		alconvert.PrintForHumans(av)
	} else if *printjson {
		alconvert.PrintJson(av)
	}
}

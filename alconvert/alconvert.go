package alconvert

import (
	"fmt"
	"time"
)

// The structures are in the structs.go file in the current directory.

// Creates a new instance of the Alcovalues struct.
func NewAV() *Alcovalues {
	return &Alcovalues{}
}

// Resets values from an existing instance by going through all fields.
func (alcval *Alcovalues) ResetAV() {
	alcval.UserSet.Milliliters = 0
	alcval.UserSet.Percent = 0
	alcval.UserSet.UnitTarget = 0
	alcval.UserSet.PercenTarget = 0
	alcval.UserSet.TargetMl = 0

	alcval.calcGotUnits.gotUnits = 0

	alcval.calcTargetUnits.gotTargUnitsFinalAmount = 0
	alcval.calcTargetUnits.gotTargUnitsRemAmount = 0

	alcval.calcTargetPercent.gotTargPercAddWater = 0
	alcval.calcTargetPercent.gotTargPercAlcLeft = 0

	alcval.calcTargetMl.gotTargMlNewAlcPerc = 0
	alcval.calcTargetMl.gotTargMlNeededWater = 0
}

// Print Alcohol Values Human Readable (sorta)
func (alcval *Alcovalues) PrintForHumans() {
	fmt.Printf("----- User Set -----\n")
	// Values set by user

	if alcval.UserSet.Milliliters != 0 {
		fmt.Printf("milliliters:\n\t%g\n", alcval.UserSet.Milliliters)
	}
	if alcval.UserSet.Percent != 0 {
		fmt.Printf("alcohol percentage (concentration):\n\t%g\n", alcval.UserSet.Percent)
	}
	if alcval.UserSet.UnitTarget != 0 {
		fmt.Printf("needed units (target units):\n\t%g\n", alcval.UserSet.UnitTarget)
	}
	if alcval.UserSet.PercenTarget != 0 {
		fmt.Printf("needed percentage (target percentage/concentration):\n\t%g\n", alcval.UserSet.PercenTarget)
	}
	if alcval.UserSet.TargetMl != 0 && alcval.UserSet.TargetMl != alcval.UserSet.Milliliters {
		fmt.Printf("needed milliliters (target ml/amount):\n\t%g\n", alcval.UserSet.TargetMl)
	}

	fmt.Printf("\n----- Calculations -----\n")
	// Values set by functions after doing the calculations

	if alcval.calcGotUnits.gotUnits != 0 {
		fmt.Printf("calculated units using milliliters and percentage:\n\t%g\n", alcval.calcGotUnits.gotUnits)
	}

	if alcval.lastOperation == "CalcTargetUnits" {
		fmt.Printf("calculated amount of alcohol (in ml) to remove,\nin order to reach the target "+
			"units (at the same percentage):\n\t%g\n", alcval.calcTargetUnits.gotTargUnitsRemAmount)

		fmt.Printf("total amount alcohol left,\nafter removing calculated alcohol for target units:"+
			"\n\t%g\n", alcval.calcTargetUnits.gotTargUnitsFinalAmount)
	}

	if alcval.lastOperation == "CalcTargetPercent" {
		fmt.Printf("calculated amount of water (in ml) to add,\nto reach target percentage:"+
			"\n\t%g\n", alcval.calcTargetPercent.gotTargPercAddWater)

		fmt.Printf("total amount diluted alcohol left,\nafter adding calculated water:"+
			"\n\t%g\n", alcval.calcTargetPercent.gotTargPercAlcLeft)
	}

	if alcval.lastOperation == "CalcTargetMl" {
		fmt.Printf("alcohol becomes this percentage(concentration),\n"+
			"after adding water for target ml:\n\t%g\n", alcval.calcTargetMl.gotTargMlNewAlcPerc)

		fmt.Printf("total amount of water added\nin alcohol for target ml:"+
			"\n\t%g\n", alcval.calcTargetMl.gotTargMlNeededWater)
	}

	fmt.Println("\nLast operation:", alcval.lastOperation)

	fmt.Printf("\n----- Timestamp -----\n")

	fmt.Print("Last calculation done: ")
	fmt.Println(alcval.timestamp)
}

// Calculates the resulting units using the
// UserSet Milliliters and Percent.
//
// results: GotUnits
func (alcval *Alcovalues) CalcGotUnits() {
	if alcval.UserSet.Percent != 0 {
		alcval.calcGotUnits.gotUnits = (alcval.UserSet.Milliliters * (alcval.UserSet.Percent / 100)) / 10
	}

	alcval.lastOperation = "CalcGotUnits"
	alcval.timestamp = time.Now()
}

// Calculates the amount of alcohol that needs to be
// removed, so that the UserSet UnitTarget can be reached.
//
// results: GotTargUnitsFinalAmount, GotTargUnitsRemAmount
func (alcval *Alcovalues) CalcTargetUnits() {
	if alcval.UserSet.UnitTarget != 0 && alcval.UserSet.Percent != 0 {
		alcval.calcTargetUnits.gotTargUnitsFinalAmount = (alcval.UserSet.UnitTarget * 10) /
			(alcval.UserSet.Percent / 100)

		alcval.calcTargetUnits.gotTargUnitsRemAmount = alcval.UserSet.Milliliters -
			alcval.calcTargetUnits.gotTargUnitsFinalAmount
	}

	alcval.lastOperation = "CalcTargetUnits"
	alcval.timestamp = time.Now()
}

// Calculates the amount of diluted alcohol that needs
// to be reached, so that the UserSet PercenTarget is reached.
//
// results: GotTargPercAddWater, GotTargPercAlcLeft
func (alcval *Alcovalues) CalcTargetPercent() {
	if alcval.UserSet.Percent != 0 && alcval.UserSet.PercenTarget != 0 {
		alcval.calcTargetPercent.gotTargPercAddWater = ((alcval.UserSet.Percent / alcval.UserSet.PercenTarget) *
			alcval.UserSet.Milliliters) - alcval.UserSet.Milliliters

		alcval.calcTargetPercent.gotTargPercAlcLeft = alcval.calcTargetPercent.gotTargPercAddWater +
			alcval.UserSet.Milliliters
	}

	alcval.lastOperation = "CalcTargetPercent"
	alcval.timestamp = time.Now()
}

// Calculates the new alcohol percentage after dilution
// and the water needed to reach the UserSet TargetMl.
//
// results: GotTargMlNewAlcPerc, GotTargMlNeededWater
func (alcval *Alcovalues) CalcTargetMl() {
	if alcval.UserSet.Milliliters != 0 && alcval.UserSet.TargetMl != 0 {
		alcval.calcTargetMl.gotTargMlNewAlcPerc = (alcval.UserSet.Milliliters / alcval.UserSet.TargetMl) *
			alcval.UserSet.Percent

		alcval.calcTargetMl.gotTargMlNeededWater = alcval.UserSet.TargetMl - alcval.UserSet.Milliliters
	}

	alcval.lastOperation = "CalcTargetMl"
	alcval.timestamp = time.Now()
}

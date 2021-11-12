package alconvert

import (
	"fmt"
	"time"
)

type UserSet struct {
	// starting amount
	Milliliters float32
	// alcohol percentage (concentration)
	Percent float32
	// needed units (target units)
	UnitTarget float32
	// needed percentage (target percentage/concentration)
	PercenTarget float32
	// needed milliliters (target ml/amount)
	TargetMl float32
}

type calcGotUnits struct {
	// the calculated units using the present Ml and UserSet.Percent
	gotUnits float32
}

type calcTargetUnits struct {
	// calculated milliliters needed for target units at same concentration
	finalMl float32

	// calculated amount to remove, to get to finalMl
	// this could be a negative number indicating amount to add
	finalRemoveAmount float32
}

type calcTargetPercent struct {
	// amount of water (in ml) to add in order to reach final_target_percent_all
	finalTargetPercent float32

	// total amount after adding water for target percent
	finalTargetPercentAll float32
}

type calcTargetMl struct {
	// if water is added this is the percent it becomes
	finalTargetMlPercent float32

	// the difference between starting ml and target ml
	finalTargetMlDiff float32
}

// Alcovalues contains all of the variables used when calculating and initial setup
type Alcovalues struct {
	// values that should be used as initializing inputs////////////////////
	UserSet UserSet

	// values that are set by the functions in this file////////////////////
	calcGotUnits      calcGotUnits
	calcTargetUnits   calcTargetUnits
	calcTargetPercent calcTargetPercent
	calcTargetMl      calcTargetMl

	// time of last calculation/////////////////////////////////////////////
	timestamp time.Time

	// the last function used for calculating///////////////////////////////
	lastOperation string
}

// NewAV Create a new instance of Alcovalues
func NewAV() *Alcovalues {
	return &Alcovalues{}
}

// ResetAV Reset values from an existing instance by going through all fields
// No need for a new instance every time you need a fresh calculation
func ResetAV(alcval *Alcovalues) {
	alcval.UserSet.Milliliters = 0
	alcval.UserSet.Percent = 0
	alcval.UserSet.UnitTarget = 0
	alcval.UserSet.PercenTarget = 0
	alcval.UserSet.TargetMl = 0

	alcval.calcGotUnits.gotUnits = 0

	alcval.calcTargetUnits.finalMl = 0
	alcval.calcTargetUnits.finalRemoveAmount = 0

	alcval.calcTargetPercent.finalTargetPercent = 0
	alcval.calcTargetPercent.finalTargetPercentAll = 0

	alcval.calcTargetMl.finalTargetMlPercent = 0
	alcval.calcTargetMl.finalTargetMlDiff = 0
}

// PrintForHumans - Print Alcohol Values Human Readable (sorta)
func PrintForHumans(alcval *Alcovalues) {
	fmt.Println()
	fmt.Println("-------------------- User set --------------------")
	fmt.Println()
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

	fmt.Println()
	fmt.Println("-------------------- Calculations --------------------")
	fmt.Println()
	// Values set by functions after doing the calculations

	if alcval.calcGotUnits.gotUnits != 0 {
		fmt.Printf("calculated units using milliliters and percentage:\n\t%g\n", alcval.calcGotUnits.gotUnits)
	}

	if alcval.lastOperation == "CalcTargetUnits" {
		fmt.Printf("calculated amount of alcohol (in ml) to remove\nin order to reach the target "+
			"units (at the same percentage):\n\t%g\n", alcval.calcTargetUnits.finalRemoveAmount)

		fmt.Printf("total amount alcohol left after removing calculated alcohol for target units:"+
			"\n\t%g\n", alcval.calcTargetUnits.finalMl)
	}

	if alcval.lastOperation == "CalcTargetPercent" {
		fmt.Printf("calculated amount of water (in ml) to add, to reach target percentage:"+
			"\n\t%g\n", alcval.calcTargetPercent.finalTargetPercent)

		fmt.Printf("total amount alcohol left after adding calculated water:"+
			"\n\t%g\n", alcval.calcTargetPercent.finalTargetPercentAll)
	}

	if alcval.lastOperation == "CalcTargetMl" {
		fmt.Printf("alcohol becomes this percentage(concentration)\n"+
			"after adding water for target ml:\n\t%g\n", alcval.calcTargetMl.finalTargetMlPercent)

		fmt.Printf("total amount of water added in alcohol for target ml:"+
			"\n\t%g\n", alcval.calcTargetMl.finalTargetMlDiff)
	}

	fmt.Println()
	fmt.Println("-------------------- Timestamp --------------------")
	fmt.Println()

	fmt.Print("Last calculation done: ")
	fmt.Println(alcval.timestamp)

	fmt.Println()
	fmt.Println("-------------------- End --------------------")
	fmt.Println()
}

// CalcGotUnits calculate units from the basic milliliters and
// percentage in the Alcovalues struct
func CalcGotUnits(alcval *Alcovalues) {
	if alcval.UserSet.Percent != 0 {
		alcval.calcGotUnits.gotUnits = (alcval.UserSet.Milliliters * (alcval.UserSet.Percent / 100)) / 10
	}

	alcval.lastOperation = "CalcGotUnits"
	alcval.timestamp = time.Now()
}

// CalcTargetUnits calculate amount of alcohol that needs to be
// removed so that the target units can be reached
func CalcTargetUnits(alcval *Alcovalues) {
	if alcval.UserSet.UnitTarget != 0 && alcval.UserSet.Percent != 0 {
		alcval.calcTargetUnits.finalMl = (alcval.UserSet.UnitTarget * 10) / (alcval.UserSet.Percent / 100)

		alcval.calcTargetUnits.finalRemoveAmount = alcval.UserSet.Milliliters - alcval.calcTargetUnits.finalMl
	}

	alcval.lastOperation = "CalcTargetUnits"
	alcval.timestamp = time.Now()
}

// CalcTargetPercent calculate amount of alcohol (diluted) that needs
// to be reached so that the target percentage is reached
func CalcTargetPercent(alcval *Alcovalues) {
	if alcval.UserSet.Percent != 0 && alcval.UserSet.PercenTarget != 0 {
		alcval.calcTargetPercent.finalTargetPercent = ((alcval.UserSet.Percent / alcval.UserSet.PercenTarget) *
			alcval.UserSet.Milliliters) - alcval.UserSet.Milliliters

		alcval.calcTargetPercent.finalTargetPercentAll = alcval.calcTargetPercent.finalTargetPercent +
			alcval.UserSet.Milliliters
	}

	alcval.lastOperation = "CalcTargetPercent"
	alcval.timestamp = time.Now()
}

// CalcTargetMl calculate the amount of dilution and final percentage
// if we want to reach the target milliliters
func CalcTargetMl(alcval *Alcovalues) {
	if alcval.UserSet.Milliliters != 0 && alcval.UserSet.TargetMl != 0 {
		alcval.calcTargetMl.finalTargetMlPercent = (alcval.UserSet.Milliliters / alcval.UserSet.TargetMl) *
			alcval.UserSet.Percent

		alcval.calcTargetMl.finalTargetMlDiff = alcval.UserSet.TargetMl - alcval.UserSet.Milliliters
	}

	alcval.lastOperation = "CalcTargetMl"
	alcval.timestamp = time.Now()
}

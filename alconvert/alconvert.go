package alconvert

import (
	"fmt"
	"time"
)

// Alcovalues contains all of the variables used when calculating and initial setup
type Alcovalues struct {
	// values that should be used as initializing inputs////////////////////

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
	////////////////////////////////////////////////////////////////////////

	// values that are set by the functions in this file///////////////////

	// the calculated units using the present Ml and Percent
	gotUnits float32
	// calculated milliliters needed for target units at same concentration
	finalTargetUnitsMl float32
	// calculated amount to remove, to get to finalTargetUnitsMl
	// this could be a negative number indicating amount to add
	finalRemoveAmount float32
	// amount of water (in ml) to add in order to reach final_target_percent_all
	finalTargetPercent float32
	// total amount after adding water for target percent
	finalTargetPercentAll float32
	// if water is added this is the percent it becomes
	finalTargetMlPercent float32
	// the difference between starting ml and target ml
	finalTargetMlDiff float32
	// time of last calculation
	timestamp time.Time
	// the last function used for calculating
	lastOperation string
	////////////////////////////////////////////////////////////////////////
}

// NewAV Create a new instance of Alcovalues
func NewAV() *Alcovalues {
	return &Alcovalues{}
}

// ResetAV Reset values from an existing instance by going through all fields
// No need for a new instance every time you need a fresh calculation
func ResetAV(alcval *Alcovalues) {
	alcval.Milliliters = 0
	alcval.Percent = 0
	alcval.UnitTarget = 0
	alcval.PercenTarget = 0
	alcval.TargetMl = 0

	alcval.gotUnits = 0
	alcval.finalTargetUnitsMl = 0
	alcval.finalRemoveAmount = 0
	alcval.finalTargetPercent = 0
	alcval.finalTargetPercentAll = 0
	alcval.finalTargetMlPercent = 0
	alcval.finalTargetMlDiff = 0
}

// PrintForHumans - Print Alcohol Values Human Readable (sorta)
func PrintForHumans(alcval *Alcovalues) {
	fmt.Println()
	fmt.Println("-------------------- User set --------------------")
	fmt.Println()
	// Values set by user

	if alcval.Milliliters != 0 {
		fmt.Printf("milliliters:\n\t%g\n", alcval.Milliliters)
	}
	if alcval.Percent != 0 {
		fmt.Printf("alcohol percentage (concentration):\n\t%g\n", alcval.Percent)
	}
	if alcval.UnitTarget != 0 {
		fmt.Printf("needed units (target units):\n\t%g\n", alcval.UnitTarget)
	}
	if alcval.PercenTarget != 0 {
		fmt.Printf("needed percentage (target percentage/concentration):\n\t%g\n", alcval.PercenTarget)
	}
	if alcval.TargetMl != 0 && alcval.TargetMl != alcval.Milliliters {
		fmt.Printf("needed milliliters (target ml/amount):\n\t%g\n", alcval.TargetMl)
	}

	fmt.Println()
	fmt.Println("-------------------- Calculations --------------------")
	fmt.Println()
	// Values set by functions after doing the calculations

	if alcval.gotUnits != 0 {
		fmt.Printf("calculated units using milliliters and percentage:\n\t%g\n", alcval.gotUnits)
	}

	if (alcval.lastOperation == "CalcTargetUnits") {
		fmt.Printf("calculated amount of alcohol (in ml) to remove in order to reach the target\nunits (at the same percentage):\n\t%g\n", alcval.finalRemoveAmount)

		fmt.Printf("total amount alcohol left after removing calculated alcohol for target units:\n\t%g\n", alcval.finalTargetUnitsMl)
	}

	if (alcval.lastOperation == "CalcTargetPercent") {
		fmt.Printf("calculated amount of water (in ml) to add, to reach target percentage:\n\t%g\n", alcval.finalTargetPercent)

		fmt.Printf("total amount alcohol left after adding calculated water:\n\t%g\n", alcval.finalTargetPercentAll)
	}

	if (alcval.lastOperation == "CalcTargetMl") {
		fmt.Printf("alcohol becomes this percentage(concentration)\nafter adding water for target ml:\n\t%g\n", alcval.finalTargetMlPercent)

		fmt.Printf("total amount of water added in alcohol for target ml:\n\t%g\n", alcval.finalTargetMlDiff)
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
	if alcval.Percent != 0 {
		alcval.gotUnits = (alcval.Milliliters * (alcval.Percent / 100)) / 10
	}
	alcval.lastOperation = "CalcGotUnits"
	alcval.timestamp = time.Now()
}

// CalcTargetUnits calculate amount of alcohol that needs to be
// removed so that the target units can be reached
func CalcTargetUnits(alcval *Alcovalues) {
	if alcval.UnitTarget != 0 && alcval.Percent != 0 {
		alcval.finalTargetUnitsMl = (alcval.UnitTarget * 10) / (alcval.Percent / 100)
		alcval.finalRemoveAmount = alcval.Milliliters - alcval.finalTargetUnitsMl
	}
	alcval.lastOperation = "CalcTargetUnits"
	alcval.timestamp = time.Now()
}

// CalcTargetPercent calculate amount of alcohol (diluted) that needs
// to be reached so that the target percentage is reached
func CalcTargetPercent(alcval *Alcovalues) {
	if alcval.Percent != 0 && alcval.PercenTarget != 0 {
		alcval.finalTargetPercent = (alcval.Percent/alcval.PercenTarget)*alcval.Milliliters - alcval.Milliliters
		alcval.finalTargetPercentAll = alcval.finalTargetPercent + alcval.Milliliters
	}
	alcval.lastOperation = "CalcTargetPercent"
	alcval.timestamp = time.Now()
}

// CalcTargetMl calculate the amount of dilution and final percentage
// if we want to reach the target milliliters
func CalcTargetMl(alcval *Alcovalues) {
	if alcval.Milliliters != 0 && alcval.TargetMl != 0 {
		alcval.finalTargetMlPercent = (alcval.Milliliters / alcval.TargetMl) * alcval.Percent
		alcval.finalTargetMlDiff = alcval.TargetMl - alcval.Milliliters
	}
	alcval.lastOperation = "CalcTargetMl"
	alcval.timestamp = time.Now()
}

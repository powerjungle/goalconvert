package alconvert

import (
	"encoding/json"
	"fmt"
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
	GotUnits float32
	// calculated milliliters needed for target units at same concentration
	FinalTargetUnitsMl float32
	// calculated amount to remove, to get to FinalTargetUnitsMl
	// this could be a negative number indicating amount to add
	FinalRemoveAmount float32
	// amount of water (in ml) to add in order to reach final_target_percent_all
	FinalTargetPercent float32
	// total amount after adding water for target percent
	FinalTargetPercentAll float32
	// if water is added this is the percent it becomes
	FinalTargetMlPercent float32
	// the difference between starting ml and target ml
	FinalTargetMlDiff float32
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

	alcval.GotUnits = 0
	alcval.FinalTargetUnitsMl = 0
	alcval.FinalRemoveAmount = 0
	alcval.FinalTargetPercent = 0
	alcval.FinalTargetPercentAll = 0
	alcval.FinalTargetMlPercent = 0
	alcval.FinalTargetMlDiff = 0
}

// PrintForHumans Print Alcohol Values Human Readable (sorta)
// Use this carefully as it might give you a headache if you constantly spam it
func PrintForHumans(alcval *Alcovalues) {
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
	fmt.Println("------")
	if alcval.GotUnits != 0 {
		fmt.Printf("calculated units using milliliters and percentage:\n\t%g\n", alcval.GotUnits)
	}
	if alcval.FinalRemoveAmount != 0 {
		fmt.Printf("calculated amount of alcohol (in ml) to remove in\norder to reach the target units\n(at the same percentage):\n\t%g\n", alcval.FinalRemoveAmount)
	}
	if alcval.FinalTargetUnitsMl != 0 && alcval.FinalTargetUnitsMl != alcval.Milliliters {
		fmt.Printf("total amount alcohol left after removing\ncalculated alcohol for target units:\n\t%g\n", alcval.FinalTargetUnitsMl)
	}
	if alcval.FinalTargetPercent != 0 {
		fmt.Printf("calculated amount of water (in ml) to add,\nto reach target percentage:\n\t%g\n", alcval.FinalTargetPercent)
	}
	if alcval.FinalTargetPercentAll != 0 && alcval.FinalTargetPercentAll != alcval.Milliliters {
		fmt.Printf("total amount alcohol left after\nadding calculated water:\n\t%g\n", alcval.FinalTargetPercentAll)
	}
	if alcval.FinalTargetMlPercent != 0 {
		fmt.Printf("alcohol becomes this percentage(concentration)\nafter adding water for target ml:\n\t%g\n", alcval.FinalTargetMlPercent)
	}
	if alcval.FinalTargetMlDiff != 0 {
		fmt.Printf("total amount of water added\nin alcohol for target ml:\n\t%g\n", alcval.FinalTargetMlDiff)
	}
}

// PrintJSON prints the Alcovalues in json format
func PrintJSON(alcval *Alcovalues) {
	ret, err := json.MarshalIndent(alcval, "", "\t")
	if err == nil {
		fmt.Println(string(ret))
	} else {
		fmt.Println(err)
	}
}

// CalcGotUnits calculate units from the basic milliliters and
// percentage in the Alcovalues struct
func CalcGotUnits(alcval *Alcovalues) {
	if alcval.Percent != 0 {
		alcval.GotUnits = (alcval.Milliliters * (alcval.Percent / 100)) / 10
	}
}

// CalcTargetUnits calculate amount of alcohol that needs to be
// removed so that the target units can be reached
func CalcTargetUnits(alcval *Alcovalues) {
	if alcval.UnitTarget != 0 && alcval.Percent != 0 {
		alcval.FinalTargetUnitsMl = (alcval.UnitTarget * 10) / (alcval.Percent / 100)
	}

	alcval.FinalRemoveAmount = alcval.Milliliters - alcval.FinalTargetUnitsMl
}

// CalcTargetPercent calculate amount of alcohol (diluted) that needs
// to be reached so that the target percentage is reached
func CalcTargetPercent(alcval *Alcovalues) {
	if alcval.Percent != 0 && alcval.PercenTarget != 0 {
		alcval.FinalTargetPercent = (alcval.Percent/alcval.PercenTarget)*alcval.Milliliters - alcval.Milliliters
	}
	alcval.FinalTargetPercentAll = alcval.FinalTargetPercent + alcval.Milliliters
}

// CalcTargetMl calculate the amount of dilution and final percentage
// if we want to reach the target milliliters
func CalcTargetMl(alcval *Alcovalues) {
	if alcval.Milliliters != 0 && alcval.TargetMl != 0 {
		alcval.FinalTargetMlPercent = (alcval.Milliliters / alcval.TargetMl) * alcval.Percent
	}
	alcval.FinalTargetMlDiff = alcval.TargetMl - alcval.Milliliters
}
